package apaleo

import (
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scope                = ""
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://identity.apaleo.com/connect/authorize",
				TokenURL: "https://identity.apaleo.com/connect/token",
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	// These are not registered in the oauth library by default
	oauth2.RegisterBrokenAuthHeaderProvider(baseURL.String())

	c.Config.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/oauth",
		TokenURL: baseURL.String() + "/access_token",
	}
}

type Oauth2ClientCredentialsConfig struct {
	clientcredentials.Config
}

func NewOauth2ClientCredentialsConfig() *Oauth2ClientCredentialsConfig {
	config := &Oauth2ClientCredentialsConfig{
		Config: clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			TokenURL:     "https://identity.apaleo.com/connect/token",
			AuthStyle:    oauth2.AuthStyleInHeader,
		},
	}

	return config
}
