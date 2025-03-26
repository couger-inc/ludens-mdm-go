package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	userStore "github.com/couger-inc/ludens-mdm/crud"
	"github.com/couger-inc/ludens-mdm/crud/db"
	"github.com/golang-jwt/jwt/v5"
)

const (
	publicKeysUrl = "https://www.googleapis.com/identitytoolkit/v3/relyingparty/publicKeys"
)


type AuthenticatorFunc func(context.Context, events.APIGatewayProxyRequest) error

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
//     const decoded = jwt.decode(token, { complete: true });
//     if (decoded == null) throw new Error("invalid JWT");

//     return decoded.payload as jwt.JwtPayload;
//   }
	//TODO:
	//serviceKey, serviceKeyExists := os.LookupEnv("SERVICE_KEY")
	//if !serviceKeyExists {
	//		return fmt.Errorf("No key")/
	//	}
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
   decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(publicKey))
		return key, err
	}, jwt.WithAudience("ludens-couger-dev"), jwt.WithIssuer("https://session.firebase.google.com/ludens-couger-dev"))
	if err != nil {
		return nil, err
	}
	// if !t.Valid {
	// 	return nil, fmt.Errorf("invalid JWT (decoded as : %v)", decoded)
	// }
	return decoded.Claims, nil
}


func AuthenticateWithCookie(ctx context.Context, request events.APIGatewayProxyRequest) error{
	cookies, err := http.ParseCookie(request.Headers["Cookie"])
	for _, cookie := range cookies {
		if (cookie.Name == "session") {
			_, err := verifyCookie(cookie.Value)
			if err != nil {
				return err
			} else {
				return nil
			}
		}
	}
	return err
}

func AuthenticateWithToken(ctx context.Context, request events.APIGatewayProxyRequest) error{
	// TODO: 
	if token := getToken(request); token != nil {
		return nil
	}
	return fmt.Errorf("Test")
}

func AuthenticateWithAccessKey(ctx context.Context, request events.APIGatewayProxyRequest) error{
	if token := getToken(request); token != nil {
		ACCESS_KEY_NAME := "admin"
		basics, err := userStore.CreateClient()
		if (err != nil) {
			return err
		}
		accessKey, err := basics.PrismaClient.AccessKey.FindUnique(db.AccessKey.Name.Equals(ACCESS_KEY_NAME)).Exec(ctx)
		if (err != nil) {
			return err
		}
		if (*token != accessKey.Value) {
			return fmt.Errorf("Invalid access key")
		}
	}
	return fmt.Errorf("No token")
}