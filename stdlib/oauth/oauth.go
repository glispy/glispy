package oauth

import (
	"context"
	"net/http"

	"github.com/hatchify/oauth2"
)

// NewClient will return a new OAuth client using the provided client ID, secret, and authorization code
func NewClient(clientID, clientSecret, authCode string, scopes ...string) (client *http.Client, err error) {
	ctx := context.Background()
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://provider.com/o/oauth2/auth",
			TokenURL: "https://provider.com/o/oauth2/token",
		},
	}

	var token *oauth2.Token
	if token, err = conf.Exchange(ctx, authCode); err != nil {
		return
	}

	client = conf.Client(ctx, token)
	return
}
