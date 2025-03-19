module github.com/couger-inc/ludens-mdm/user-store/get

go 1.24.0

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/couger-inc/ludens-mdm/crud v0.0.0-00010101000000-000000000000
	github.com/couger-inc/ludens-mdm/middlewares v0.0.0-00010101000000-000000000000
	github.com/couger-inc/ludens-mdm/middlewares/auth v0.0.0-00010101000000-000000000000
	github.com/couger-inc/ludens-mdm/openapi v0.0.0-00010101000000-000000000000
	github.com/mitchellh/mapstructure v1.5.0
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/steebchen/prisma-client-go v0.47.0 // indirect
	go.mongodb.org/mongo-driver/v2 v2.0.1 // indirect
)

replace github.com/couger-inc/ludens-mdm/crud => ../../crud


replace github.com/couger-inc/ludens-mdm/openapi => ../../openapi

replace github.com/couger-inc/ludens-mdm/middlewares => ../../middleware

replace github.com/couger-inc/ludens-mdm/middlewares/auth => ../../middleware/auth
