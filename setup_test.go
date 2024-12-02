package apaleo_test

import (
	"context"
	"os"

	"github.com/omniboost/go-apaleo"
)

func client() *apaleo.Client {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")

	oauthConfig := apaleo.NewOauth2ClientCredentialsConfig()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client := apaleo.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)
	return client
}
