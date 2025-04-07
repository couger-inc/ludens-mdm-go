package middleware

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	parameterandsecrets "github.com/couger-inc/ludens-mdm-go/aws/parameters-and-secrets"
	"github.com/couger-inc/ludens-mdm-go/middlewares/auth"
)
type apiGatewayHandlerFunc func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)
type handlerFunc func(context.Context, events.APIGatewayProxyRequest) (string, int)
// the extra set of instructions
// things to be done before running the business logic
func RequestResponseLogger(f apiGatewayHandlerFunc) apiGatewayHandlerFunc {
	return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		log.Printf("remote_addr: %s", r.RequestContext.Identity.SourceIP)
		log.Printf("query params: %s", r.QueryStringParameters)
		log.Printf("path params: %s", r.PathParameters)
		log.Printf("body: %s", r.Body)
		result, err := f(ctx, r)
		if (err != nil) {
			log.Printf("Error: %v", err)
		} else {
			log.Printf("Response{StatusCode: %v, Body: %s}", result.StatusCode, result.Body)
		}
		return result, err
	}
}

func APIGatewayProxyResponseMiddleware(f handlerFunc) apiGatewayHandlerFunc {
	return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		response, statusCode := f(ctx, r)
		return events.APIGatewayProxyResponse{
			StatusCode: statusCode,
			Body: response,
		}, nil
	}
}

func ParamStoreMiddleware(f apiGatewayHandlerFunc) apiGatewayHandlerFunc {
	return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		ssmsvc := parameterandsecrets.NewSSMClient()
		result,err := ssmsvc.Param("/ludens-mdm/database_url", true).GetValue()
		if (err != nil) {
			log.Printf("Unable to fetch parameters from ssm store. Error: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body: err.Error(),
			}, nil
		}
		os.Setenv("DATABASE_URL", result)
		result,err = ssmsvc.Param("/ludens-user-console/google_application_credentials", true).GetValue()
		if (err != nil) {
			log.Printf("Unable to fetch parameters from ssm store. Error: %v", err)
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body: err.Error(),
			}, nil
		}
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS_DATA", result)
		return f(ctx, r)
	}
}

func AuthenticateAny(f handlerFunc, authenticators ...auth.AuthenticatorFunc) handlerFunc {
	return func(ctx context.Context, r events.APIGatewayProxyRequest) (string, int) {
		authenticated := false
		newCtx := ctx
		for _, authenticator := range authenticators {
			authCtx, err := authenticator(ctx, r)
			if err == nil {
				authenticated = true
				newCtx = authCtx
				log.Printf("%v, authenticated", authenticator);
				break
			}
			log.Printf("%v, %v, not authenticated", authenticator, err);
		}
		if (authenticated) {
			return f(newCtx, r)
		} else {
			return "Not authenticated", 401
		}
	}
}