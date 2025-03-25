package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	userStore "github.com/couger-inc/ludens-mdm/crud"
	middleware "github.com/couger-inc/ludens-mdm/middlewares"
	auth "github.com/couger-inc/ludens-mdm/middlewares/auth"
	"github.com/couger-inc/ludens-mdm/openapi"
	"github.com/mitchellh/mapstructure"
)


func convertRequest(event events.APIGatewayProxyRequest, request *openapi.GetManagersParams) error {
	offset := "0"
	limit := "100"
	request.Offset = &offset
	request.Limit = &limit
	err := mapstructure.Decode(event.QueryStringParameters, &request)
	return err
}

func handler(ctx context.Context, event events.APIGatewayProxyRequest) (string, int) {
	storeId := event.PathParameters["storeId"]
	var request openapi.GetManagersParams
	if err := convertRequest(event, &request); err != nil {
		return fmt.Sprintf("Unable to decode request parameters: %v", err.Error()), 500
	}
	basics, err := userStore.CreateClient()
	if err != nil {
		return fmt.Sprintf("Unable to connect to the database: %v", err.Error()), 500
	}
	skip, err := strconv.Atoi(*request.Offset)
	if err != nil {
		return fmt.Sprintf("Unable to convert request parameter, Offset, to an integer: %v", err.Error()), 500
	}
	take, err := strconv.Atoi(*request.Limit)
	if err != nil {
		return fmt.Sprintf("Unable to convert request parameter, Limit, to an integer: %v", err.Error()), 500
	}
	stores, totalCount, err := basics.GetUserStores(ctx, skip, take, storeId)
	if err != nil {
		return fmt.Sprintf("Unable to retrieve stores: %v", err.Error()), 500
	}
	convertedStoreObjects := []openapi.StoreObject{}
	for _, store := range stores {
		managers := []openapi.ManagerObject{}
		for _, manager := range store.UserStore() {
			managers = append(managers, openapi.ManagerObject{
				Email: manager.Email,
				Name: manager.Name,
			})
		}
		convertedStoreObjects = append(convertedStoreObjects, openapi.StoreObject{
			Id: &storeId,
			Name: store.Name,
			Managers: &managers,
		})
	}
	body := openapi.GetManagersResponse{Stores: convertedStoreObjects, TotalCount: totalCount}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return fmt.Sprintf("Unable to convert response object to json: %v", err.Error()), 500
	}
	defer basics.Disconnect()
	return string(jsonBody), 200
}

func main() {
	lambda.Start(middleware.RequestResponseLogger(middleware.APIGatewayProxyResponseMiddleware(middleware.AuthenticateAny(handler, auth.AuthenticateWithCookie, auth.AuthenticateWithToken, auth.AuthenticateWithAccessKey))))
}
