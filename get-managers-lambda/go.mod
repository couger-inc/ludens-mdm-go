module github.com/couger-inc/ludens-mdm/manager/get

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
	cel.dev/expr v0.16.1 // indirect
	cloud.google.com/go v0.117.0 // indirect
	cloud.google.com/go/auth v0.13.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.6 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/firestore v1.18.0 // indirect
	cloud.google.com/go/iam v1.2.2 // indirect
	cloud.google.com/go/longrunning v0.6.2 // indirect
	cloud.google.com/go/monitoring v1.21.2 // indirect
	cloud.google.com/go/storage v1.49.0 // indirect
	firebase.google.com/go/v4 v4.15.2 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp v1.25.0 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/metric v0.48.1 // indirect
	github.com/GoogleCloudPlatform/opentelemetry-operations-go/internal/resourcemapping v0.48.1 // indirect
	github.com/MicahParks/keyfunc v1.9.0 // indirect
	github.com/aws/aws-sdk-go v1.55.6 // indirect
	github.com/census-instrumentation/opencensus-proto v0.4.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cncf/xds/go v0.0.0-20240905190251-b4127c9b8d78 // indirect
	github.com/couger-inc/ludens-mdm/aws/parameters-and-secrets v0.0.0-00010101000000-000000000000 // indirect
	github.com/couger-inc/ludens-mdm/firebase v0.0.0-00010101000000-000000000000 // indirect
	github.com/envoyproxy/go-control-plane v0.13.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v1.1.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
	github.com/planetscale/vtprotobuf v0.6.1-0.20240319094008-0393e58bdf10 // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/steebchen/prisma-client-go v0.47.0 // indirect
	go.mongodb.org/mongo-driver/v2 v2.0.1 // indirect
	go.opentelemetry.io/contrib/detectors/gcp v1.29.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.54.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.54.0 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk v1.29.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/net v0.33.0 // indirect
	golang.org/x/oauth2 v0.25.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.8.0 // indirect
	google.golang.org/api v0.215.0 // indirect
	google.golang.org/appengine/v2 v2.0.6 // indirect
	google.golang.org/genproto v0.0.0-20241118233622-e639e219e697 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241223144023-3abc09e42ca8 // indirect
	google.golang.org/grpc v1.67.3 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
)

replace github.com/couger-inc/ludens-mdm/crud => ../crud

replace github.com/couger-inc/ludens-mdm/aws/parameters-and-secrets => ../aws/parameter-and-secrets

replace github.com/couger-inc/ludens-mdm/middlewares => ../middleware

replace github.com/couger-inc/ludens-mdm/openapi => ../openapi

replace github.com/couger-inc/ludens-mdm/middlewares/auth => ../middleware/auth

replace github.com/couger-inc/ludens-mdm/firebase => ../firebase
