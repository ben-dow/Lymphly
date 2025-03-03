package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"lymphly/internal/cfg"
	"lymphly/internal/log"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/go-chi/chi/v5"
)

var cognito *cognitoidentityprovider.Client

func init() {
	cfg, _ := config.LoadDefaultConfig(context.Background())
	cognito = cognitoidentityprovider.NewFromConfig(cfg)
}

func AuthRoutes(r chi.Router) {
	r.Get("/login", Login)
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("x-username")
	password := r.Header.Get("x-password")

	var authResult *types.AuthenticationResultType
	output, err := cognito.InitiateAuth(r.Context(), &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       "USER_PASSWORD_AUTH",
		ClientId:       aws.String(cfg.Cfg().ClientId),
		AuthParameters: map[string]string{"USERNAME": username, "PASSWORD": password},
	})
	if err != nil {
		var resetRequired *types.PasswordResetRequiredException
		if errors.As(err, &resetRequired) {
			w.Write([]byte(*resetRequired.Message))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Error("could not initiate login", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	authResult = output.AuthenticationResult

	out := map[string]string{"token": *authResult.AccessToken}
	outBytes, _ := json.Marshal(out)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-cache")
	w.Write(outBytes)

}
