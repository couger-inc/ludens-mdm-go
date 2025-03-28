module github.com/couger-inc/ludens-mdm/user-store/put

go 1.24.0

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/couger-inc/ludens-mdm/crud v0.0.0-00010101000000-000000000000
	github.com/couger-inc/ludens-mdm/middlewares v0.0.0-00010101000000-000000000000
	github.com/couger-inc/ludens-mdm/middlewares/auth v0.0.0-00010101000000-000000000000
	github.com/couger-inc/ludens-mdm/openapi v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.55.6 // indirect
	github.com/couger-inc/ludens-mdm/aws/parameters-and-secrets v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
)

require (
	github.com/couger-inc/ludens-mdm/user-console v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/steebchen/prisma-client-go v0.47.0 // indirect
	go.mongodb.org/mongo-driver/v2 v2.0.1 // indirect
)

replace github.com/couger-inc/ludens-mdm/crud => ../crud

replace github.com/couger-inc/ludens-mdm/user-console => ../user-console

replace github.com/couger-inc/ludens-mdm/openapi => ../openapi

replace github.com/couger-inc/ludens-mdm/middlewares => ../middleware

replace github.com/couger-inc/ludens-mdm/middlewares/auth => ../middleware/auth

replace github.com/couger-inc/ludens-mdm/aws/parameters-and-secrets => ../aws/parameter-and-secrets
