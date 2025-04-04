package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	userStore "github.com/couger-inc/ludens-mdm/crud"
	middleware "github.com/couger-inc/ludens-mdm/middlewares"
	auth "github.com/couger-inc/ludens-mdm/middlewares/auth"
	"github.com/couger-inc/ludens-mdm/openapi"
)

func convertRequest(event events.APIGatewayProxyRequest, request *openapi.DeleteManagersJSONRequestBody) error {
	err := json.Unmarshal([]byte(event.Body), &request)
	return err
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (string, int) {
	storeId := event.PathParameters["storeId"]
	var requestBody openapi.DeleteManagersJSONRequestBody
	err := convertRequest(event, &requestBody)
	if (err != nil) {
		return fmt.Sprintf("Unable to unmarshal request body: %v", err.Error()), 500
	}
	basics, err := userStore.CreateClient()
	if err != nil {
		return fmt.Sprintf("Unable to connect to database: %v", err.Error()), 500
	}
	results := []openapi.ManagerObject{}
	var emails []string
	for _, manager  := range *requestBody.Managers{
		emails = append(emails, manager.Email)
	}
	_, err = basics.DeleteUserStore(ctx, storeId, emails)
	if (err != nil) {
		return fmt.Sprintf("Unable to delete users: %v", err.Error()), 500
	}
	for _, manager  := range *requestBody.Managers{
		results = append(results, openapi.ManagerObject{Email: manager.Email, Name: manager.Name})
	}
	body := openapi.DeleteManagersResponse{
		Managers: results,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	defer basics.Disconnect()

	return string(jsonBody),200
}

func main() {
	lambda.Start(middleware.RequestResponseLogger(middleware.ParamStoreMiddleware(middleware.APIGatewayProxyResponseMiddleware(middleware.AuthenticateAny(handler, auth.AuthenticateWithCookie, auth.AuthenticateWithToken, auth.AuthenticateWithAccessKey)))))
}
