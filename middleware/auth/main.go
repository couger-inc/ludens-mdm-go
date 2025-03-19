package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	userStore "github.com/couger-inc/ludens-mdm/crud"
	"github.com/couger-inc/ludens-mdm/crud/db"
	"github.com/golang-jwt/jwt/v5"
)
type AuthenticatorFunc func(context.Context, events.APIGatewayProxyRequest) error

func GetPublicKeys() map[string]string {
	return map[string]string {
		"bZ-_5g": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIENJ2HCDANBgkqhkiG9w0BAQsFADAzMQ8wDQYDVQQDEwZH\naXRraXQxEzARBgNVBAoTCkdvb2dsZSBJbmMxCzAJBgNVBAYTAlVTMB4XDTI0MTIw\nNjAyMjUzM1oXDTI1MTIwMTAyMjUzM1owMzEPMA0GA1UEAxMGR2l0a2l0MRMwEQYD\nVQQKEwpHb29nbGUgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBANyVAukjUNaA9z55PCK3I803APf7g5o+x+h89pOhFQdeHZd+\nAMamdbtLsbmgfZ0lxTaIAbaEdWW9ZLFTLbsO9F8Vc38n0goxdnngS85d1stih0Wm\nY2p04qAkuyFpjVMGLoTtcep9rguc+0UuDTBya0PsEsltE0Dgt8HVGl8ZFnF4QNY4\ncIFl/JTF0JPpGxCj801L/Za+KMneni1bMPxdn7NThoVbw39MVqdIYTjnWxFnDnZ5\nUSLIxOhFHtCaf6kQw3uNkykiZiM90XzADEb3RU1uShEweuSh/W890tnG8uXOe34/\nMVSRvwcbD4sJvMC4EKsYzzJ4mN/i4GC5gHfSjyMCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjAOBgNVHQ8BAf8EBAMCB4AwDQYJ\nKoZIhvcNAQELBQADggEBAEpzGwGkgLPrumwbatUrw+sT2TuRRS3co+1cVIO6+8qz\nkv+3LgcnmUNUlLlXaSK3YKK35B8wFzrjPEt4vH5m61apkTePFcf081iee719iIH+\nfSO0G0joaKQiUwAXAiy2Nyyokpm2/hMk/5QGKCI7j7gBa1gz1Qd9+3TdSxXFL1OP\n2vs0mUqrsFFNr0ZAB+Vq4QXqBEdnZzxF+b8dV4QlfTGq/JZ6Nds6/ruR/qzivebA\nIy/1786h90Tzj6eC5s2Xlyb+IoAirszQCEh5lvx89xN07cCuQfuk0dl+ryXSGA2U\nS3hiVQLrAdXoqRKs3hG+rv1N36g3Vk4CD6BMvJlgk0I=\n-----END CERTIFICATE-----\n",
  		"QDgRaA": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIEDbefzTANBgkqhkiG9w0BAQsFADAzMQ8wDQYDVQQDEwZH\naXRraXQxEzARBgNVBAoTCkdvb2dsZSBJbmMxCzAJBgNVBAYTAlVTMB4XDTI1MDEx\nNjE4MjQ0MloXDTI2MDExMTE4MjQ0MlowMzEPMA0GA1UEAxMGR2l0a2l0MRMwEQYD\nVQQKEwpHb29nbGUgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAJ2zLkg+BflOy8dT0XOr/1PwrHWoQ2MCqFgJNXzBBSSYltbn\nGLVvdnSQud4kXTSuC5oSewYAaJMcKowxImtiaN+t8LcgshMGS7DaHEi/HOlWrYc+\nY0v9eEv5lBJ/xaApjl0c1iOWxXYXKR9rrhl8Dd4OQlgrcsM7E9fdOFtpvk7y7khe\nb2UXEyuPnbNnH/xW62eAi68bwf8rJtlqVRcYdzM/2drNZ9qVmYVdoykpUi5STr9v\nV7xGLk/ksmKZ5KFEpdKkcRXqYs/1dX8r9/B49imetiV8rUNWDFTEE4UUpAKGZE6f\nSZu0TzrQJ9iPSB83+Y82rND0mHaAxuCy4o2rOVECAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjAOBgNVHQ8BAf8EBAMCB4AwDQYJ\nKoZIhvcNAQELBQADggEBAAo2O2Sa1dfxtCJ+N7WZRrnac9StST/E9ANU7Ggv+X0S\n5FENJzP8QPjdwGmO9B73ayhMAqYsChYLR9SR6lna3e0nmNoSq1B8qQVPrXLdnlsG\nhy3QuEfp1MW+W9MfFhUICeUcTaEggs+OO5ffG4kA1xkKRNcrPKdISBV1eM5ue3KB\n35K4guBR+5CaQa4NbgeQ2D8sYFsFfPK7uZBwswvB1k5Ema0Sq7CsEUcGUGiw1CYm\nc+rC73O11V/viWIchNEIDeAhM4MCvAmNz0/1YqfxIkmc8e1HqgN3WjFEfLSVTCWA\nh+6t/LEzj7wDtAoCSekceWcSZ2ci4RE16McguFGd30M=\n-----END CERTIFICATE-----\n",
  		"rPWZTg": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIEWibRRjANBgkqhkiG9w0BAQsFADAzMQ8wDQYDVQQDEwZH\naXRraXQxEzARBgNVBAoTCkdvb2dsZSBJbmMxCzAJBgNVBAYTAlVTMB4XDTI1MDIy\nMjA0MDAxOVoXDTI2MDIxNzA0MDAxOVowMzEPMA0GA1UEAxMGR2l0a2l0MRMwEQYD\nVQQKEwpHb29nbGUgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAOwsBl9G9zcXHdgnAwrxHFMR9Smv0cDeUqlhzj3G4YT5Srx3\nboJaoWNDRvDmTBgje25vqmRXQdmuYeaOyv2+R3cH8LwzQbutdC9RKUogkd+XJ2Qm\n7YNX+qY2/UdR03MIOBstHoqKhGiDUcYMzEi5n9yGMC0yF+9BXfk5vHduu94B7W72\nckDwTGRqf4QNYvjk7J02svOTBWgm1JJ4n7abPe7DJrK1f69MUzL72Z1Fh5idGztK\nlDyXup7YZxIJsGkslqbsE+8OG7JT97Lmv/7fyuwAKr0mdtTyJ4zRXtncu9LkebxK\npKEK1j3Bhi+MIDfydtqM+PZy9jNdn3aOcy0CMYsCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjAOBgNVHQ8BAf8EBAMCB4AwDQYJ\nKoZIhvcNAQELBQADggEBAG/kDk/d2ftzpygOaq5XWnOIHSH0bDhvo6WlPgrBJDEC\nOWWKg11W+p4NE2oB96lDjSux50weJabpDRE/0VTGo7yrAli8BMHrP7afrNzGs8bE\nxSDTO2/dKeywzLkfiGi4YG61kL48/h9U3Qx1vUWojFUXCtJ7T38tLwKbMBgybhWi\ndx4WEf3RNbN54rNT4zUJRgaNYoQ19OvorRBl/VoiwRs630wilkQ4uEsMwrBSNgkC\nCTbU2vqP/r6/lqlz6jaXiwJsWH6HmaMGcZkPMvEtDhmmNckQaOxINlnzICM6MI/Y\nkWQWiSBmlpbOsdo19KcR+H5ZaBj5nAIXiTEGeV3AdrY=\n-----END CERTIFICATE-----\n",
  		"KRSnmg": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIEKh2gFTANBgkqhkiG9w0BAQsFADAzMQ8wDQYDVQQDEwZH\naXRraXQxEzARBgNVBAoTCkdvb2dsZSBJbmMxCzAJBgNVBAYTAlVTMB4XDTI0MDky\nMTE3MjcwN1oXDTI1MDkxNjE3MjcwN1owMzEPMA0GA1UEAxMGR2l0a2l0MRMwEQYD\nVQQKEwpHb29nbGUgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBAMN5WgJ7EQbhDF8UquawppXBt9gvjbL6gnfVXaQi5KvrSp+P\nPffa3UBipxezgwjGfSfp7z02HZike5bKBSIa6sGWxoDfejLyz2lkRDGpdv0vtJdt\nC9b2xqIZ2jq0UD1Vn6aWGEE+y0mvp1QEWTRK5vF6bI/QGNwRuFIGSi1Sb8KVraFW\nIgw4RsS+B5aJlZqE8leHhjO1l5NJkWEh/uwwUKFs+dpWV/9SoBKrDTyPDBt0ZvF5\nYo8Xs5PxVIoEr38JysLZpJ6AWXXLIQN3mdGBd4Wm73o5MW39vObzgsJhgZ4+0jjV\ntWVUL+KpV3mSLaUxpjGa1Fz5qKXyqRfWwSSx3DcCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjAOBgNVHQ8BAf8EBAMCB4AwDQYJ\nKoZIhvcNAQELBQADggEBAKHtwCXqXZ4vD6jCqPQtOx5C+E7kIP93siygxzOlq70I\nVcZNqbnsKZLXsbZ5VeConk1ZQicF7M6fpvyK4YrG+E3V4iBbqITZUJbmSbGE+uZo\nAq29Ax3K6b78mM7haySe2Z7Sr61WSlPMZD2T4qoRF6ogmw/N0peV0Z5mu4CX5ebS\nMXI51a4fz4E++Zy3z+uU1gXRfm4cDAFi/L+u3huNY2adyUJdqfbWuGv1JDxsdm+U\nHNckD8pdUBhycl361kX2HCt0CBgfJJQvs4rpVMd95MPzYE4N5kDoDcQZck5BGWAZ\nPtFHFEYmF0uingVJnd520EgwDNHywYYh64uxylwFtW8=\n-----END CERTIFICATE-----\n",
  		"-WZpKQ": "-----BEGIN CERTIFICATE-----\nMIIDHDCCAgSgAwIBAgIEK0Au9TANBgkqhkiG9w0BAQsFADAzMQ8wDQYDVQQDEwZH\naXRraXQxEzARBgNVBAoTCkdvb2dsZSBJbmMxCzAJBgNVBAYTAlVTMB4XDTI0MTAz\nMTAxMTcwOFoXDTI1MTAyNjAxMTcwOFowMzEPMA0GA1UEAxMGR2l0a2l0MRMwEQYD\nVQQKEwpHb29nbGUgSW5jMQswCQYDVQQGEwJVUzCCASIwDQYJKoZIhvcNAQEBBQAD\nggEPADCCAQoCggEBANfahbn31sd4bntOODPKabKHU/eeqMiXOsaiai5JvIDQS5oL\nbWSkjs0L5A04kzPaETpRQCYXzEF2Ntad96fVpESAUhXD6DUuJarlAOOQyvF0FFtC\noWwfaqnDFbkB9n8v6sK9K9XcmTInp+FocJ5T+JOGGeZNQp6Bvfz6Yejwrg2kCamo\nXj0W7WeMThJURvd0k3ntxyJtpyoH43Ljci7+ZBhgtN3HyewNruqqFTQFfxzDjPok\n1pGDW8YxeQAZVlegg9hl/UBo1yja+rSJYP9T5XrAXgMBAEyicIRORMAPi0nRahO0\nQCLOtjLMEfc/JM6s5MR4G6LPOyp3SMk1Xbmn0vcCAwEAAaM4MDYwDAYDVR0TAQH/\nBAIwADAWBgNVHSUBAf8EDDAKBggrBgEFBQcDAjAOBgNVHQ8BAf8EBAMCB4AwDQYJ\nKoZIhvcNAQELBQADggEBAMvctzjYTxBsqJLmgZXqDPs0fdxsCMzPqdFNc52p9MCM\nQksYD/AGfqeN6i22J8qTZnbMdYOJqiiED7LHdCc8RtEUlme158vfs82CZvrRLu4u\nmdnzfb9YGEJmX8MuBtVVGh27xpfhW7cb4fm+HlpO44y2ULE7srM07hC5XjVclvdH\nPD4ruJo6h2wnxJWA2R0XwFNT5sfs6FC+/tQfM0BtvAbCGyPdFJHhG/cP33/loMyQ\nWUNcnbK6AqeJxWOxksMpweD/BzTL+EQuM/6YY6xPBmm2UGY+loiu+C89FSxiq+1q\neRmT5VHwY69lEqpmBujGUk88jCA3Qmz/CAJ3X+xag8I=\n-----END CERTIFICATE-----\n",
	}
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
   publicKey, found := GetPublicKeys()[kid.(string)]
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