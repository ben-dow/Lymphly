module lymphly

go 1.24

require (
	github.com/aws/aws-lambda-go v1.47.0
	github.com/aws/aws-sdk-go-v2 v1.36.3
	github.com/aws/aws-sdk-go-v2/config v1.29.8
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.18.6
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.7.72
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v1.51.0
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.41.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.16.2
	github.com/caarlos0/env/v11 v11.3.1
	github.com/go-chi/chi/v5 v5.2.1
	github.com/mmcloughlin/geohash v0.10.0
	github.com/stretchr/testify v1.7.2
	github.com/umahmood/haversine v0.0.0-20151105152445-808ab04add26
	golang.org/x/crypto v0.35.0
	golang.org/x/sync v0.11.0
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.17.61 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.16.30 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.3.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.6.34 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.25.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.12.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.10.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.12.15 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.25.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.29.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.33.16 // indirect
	github.com/aws/smithy-go v1.22.3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/tools v0.5.1-0.20230111220935-a7f7db3f17fc // indirect
	golang.org/x/tools/cmd/cover v0.1.0-deprecated // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

tool golang.org/x/tools/cmd/cover
