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
	userconsole "github.com/couger-inc/ludens-mdm/user-console"
)

func convertRequest(event events.APIGatewayProxyRequest, request *openapi.CreateManagersJSONRequestBody) error {
	err := json.Unmarshal([]byte(event.Body), &request)
	return err
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (string, int) {
	storeId := event.PathParameters["storeId"]
	var requestBody openapi.CreateManagersJSONRequestBody
	err := convertRequest(event, &requestBody)
	if (err != nil) {
		return fmt.Sprintf("Unable to decode request parameters: %v", err.Error()), 500
	}
	basics, err := userStore.CreateClient()
	if err != nil {
		return fmt.Sprintf("Unable to connect to the database: %v", err.Error()), 500
	}
	results := []openapi.ManagerObject{}
	for _, manager := range *requestBody.Managers {
		usersResponse, err := userconsole.GetUsers(manager.Email)
		if err != nil {
			return fmt.Sprintf("Unable to query ludens user's console. Err: %v", err.Error()), 500
		} else if usersResponse.TotalCount == 0 {
			return fmt.Sprintf("User: %v not found in ludens user's console", manager.Email), 500
		} else {
			createdManager, err := basics.AddUserStore(ctx, storeId, userStore.Manager{Name: manager.Name, Email: manager.Email})
			if (err == nil) {
				results = append(results, openapi.ManagerObject{
					Name: createdManager.Name,
					Email: createdManager.Email,
				})
			} else {
				return fmt.Sprintf("Unable to create a user: %v", err.Error()), 500
			}
		}
	}
	body := openapi.CreateManagersResponse{
		Managers: results,
	}
	apiResponse, err := json.Marshal(body)
	if err != nil {
		return fmt.Sprintf("Unable to marshal response %v", err.Error()), 500
	}
	defer basics.Disconnect()
	return string(apiResponse), 200
}

func main() {
	lambda.Start(middleware.RequestResponseLogger(middleware.APIGatewayProxyResponseMiddleware(middleware.AuthenticateAny(handler, auth.AuthenticateWithCookie, auth.AuthenticateWithToken, auth.AuthenticateWithAccessKey))))
}
