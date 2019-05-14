package apaleo_test

import (
	"os"

	"github.com/omniboost/go-apaleo"
	"golang.org/x/oauth2"
)

func client() *apaleo.Client {
	clientID := os.Getenv("OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
	refreshToken := os.Getenv("OAUTH_REFRESH_TOKEN")
	tokenURL := os.Getenv("OAUTH_TOKEN_URL")

	oauthConfig := apaleo.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(oauth2.NoContext, token)

	client := apaleo.NewClient(httpClient)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)
	return client
}
