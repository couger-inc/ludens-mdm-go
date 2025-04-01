package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/aws/aws-lambda-go/events"
	userStore "github.com/couger-inc/ludens-mdm/crud"
	"github.com/couger-inc/ludens-mdm/crud/db"
	mdm_firebase_credentials "github.com/couger-inc/ludens-mdm/firebase"
	"github.com/golang-jwt/jwt/v5"
)

const (
	publicKeysUrl = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/publicKeys"
	cookieName = "session"
)


type AuthenticatorFunc func(context.Context, events.APIGatewayProxyRequest) (context.Context, error)

func GetPublicKeys() (map[string]string, error) {
	req, err := http.NewRequest("GET", publicKeysUrl, nil)
	if (err != nil) {
		return nil, err
	}
	response, err := http.DefaultClient.Do(req)
	if (err != nil) {
		return nil, err
	}
	responseData, err := io.ReadAll(response.Body)
	if (err != nil) {
		return nil, err
	}
	defer response.Body.Close()
	log.Printf("StatusCode: %v, Body: %v", response.StatusCode, string(responseData))
	if (response.StatusCode != 200) {
		return nil, fmt.Errorf("unable to query luden's user console. StatusCode: %v, Message: %v", response.StatusCode, string(responseData))
	}
	var publicKeys map[string]string
	err = json.Unmarshal(responseData, &publicKeys)
	return publicKeys, nil
}


func getQueryToken(query map[string]string) *string {
	token, ok := query["idToken"]
	if (ok) {
		return &token
	}
	return nil
}

func getBearerToken(headers map[string]string) *string {
	token, ok := headers["Authorization"]
	if (ok) {
		token = strings.TrimSpace(strings.Replace(token, "Bearer","", 1))
		return &token
	}
	return nil
}

func getToken(request events.APIGatewayProxyRequest) *string {
	if token := getBearerToken(request.Headers); token != nil {
		return token
	}
	if token := getQueryToken(request.QueryStringParameters); token != nil {
		return token
	}
	return nil
}

func verifyCookie(token string) (jwt.Claims, error) {
	//TODO:
	  // Firebase Local Emulator Suite を使っている場合は、セッション Cookie が署名されない.
//   if (process.env.FIREBASE_AUTH_EMULATOR === "true") {
	if authEmulator, exists := os.LookupEnv("FIREBASE_AUTH_EMULATOR"); exists && authEmulator == "true" {
    	decoded, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(token, jwt.MapClaims{})
    	return decoded.Claims, err
	}
	projectId, projectExists := os.LookupEnv("FIREBASE_PROJECT_ID")
	if !projectExists {
		return nil, fmt.Errorf("Firebase Project not set")
	}
	t, _, err := jwt.NewParser(jwt.WithoutClaimsValidation()).ParseUnverified(token, jwt.MapClaims{})
   if (err != nil) {
		return nil, err
   }
   
   kid, found := t.Header["kid"]
   if !found {
	return nil, fmt.Errorf("kid is not found")
   }
   publicKeys, err := GetPublicKeys()
   if (err != nil) {
	return nil, fmt.Errorf("error retrieving public keys %v", err.Error())
   }
   publicKey, found := publicKeys[kid.(string)]
   if !found {
	return nil, fmt.Errorf("no public key for %v", kid)
   }
   decoded, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		return key, err
	}, jwt.WithAudience(projectId), jwt.WithIssuer(fmt.Sprintf("https://session.firebase.google.com/%s", projectId)))
	if err != nil {
		return nil, err
	}
	return decoded.Claims, nil
}


func AuthenticateWithCookie(ctx context.Context, request events.APIGatewayProxyRequest) (context.Context, error) {
	cookies, err := http.ParseCookie(request.Headers["Cookie"])
	for _, cookie := range cookies {
		if (cookie.Name == cookieName) {
			_, err := verifyCookie(cookie.Value)
			if err != nil {
				return ctx, err
			} else {
				return ctx, nil
			}
		}
	}
	return ctx, err
}

func verifyIDToken(ctx context.Context, app *firebase.App, idToken string) *auth.Token {
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	log.Printf("Verified ID token: %v\n", token)
	return token
}

func AuthenticateWithToken(ctx context.Context, request events.APIGatewayProxyRequest) (context.Context, error) {
	app, _ := mdm_firebase_credentials.GetApp(ctx)
	if token := getToken(request); token != nil {
		decodedIdToken := verifyIDToken(ctx, app, *token)
		return context.WithValue(context.WithValue(ctx, "decodedIdToken", decodedIdToken), "uid", decodedIdToken.UID), nil
	}
	return ctx, fmt.Errorf("Unable to retrieve token")
}

func AuthenticateWithAccessKey(ctx context.Context, request events.APIGatewayProxyRequest) (context.Context, error) {
	if token := getToken(request); token != nil {
		ACCESS_KEY_NAME := "admin"
		basics, err := userStore.CreateClient()
		if (err != nil) {
			return ctx, err
		}
		accessKey, err := basics.PrismaClient.AccessKey.FindUnique(db.AccessKey.Name.Equals(ACCESS_KEY_NAME)).Exec(ctx)
		if (err != nil) {
			return ctx, err
		}
		if (*token != accessKey.Value) {
			return ctx, fmt.Errorf("Invalid access key")
		}
	}
	return ctx, fmt.Errorf("No token")
}