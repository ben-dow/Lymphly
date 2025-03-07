package auth

import (
	"context"
	"errors"
	"lymphly/internal/cfg"
	"lymphly/internal/log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var cognito *cognitoidentityprovider.Client

// Initialize the Cognito Client at Package Import for usage in routes
func init() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	cognito = cognitoidentityprovider.NewFromConfig(cfg)
}

var ErrPasswordReset = errors.New("password reset required")

func GetAccessToken(ctx context.Context, username, password string) (string, error) {
	output, err := cognito.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(cfg.Cfg().ClientId),
		AuthParameters: map[string]string{"USERNAME": username, "PASSWORD": password},
	})

	if err != nil {
		var resetRequired *types.PasswordResetRequiredException
		if errors.As(err, &resetRequired) {
			return "", errors.Join(ErrPasswordReset, err)
		}

		log.Error("could not initiate login", err)
		return "", err
	}

	return *output.AuthenticationResult.AccessToken, nil
}
