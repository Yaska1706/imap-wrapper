package authconnector

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthClient interface {
	AuthURL() string
	Exchange(code string) (*oauth2.Token, error)
}

type gmailOAuthClient struct {
	config *oauth2.Config
}

func NewGmailOAuthClient(clientID, clientSecret, redirectURL string, scopes []string) *gmailOAuthClient {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       scopes,
		Endpoint:     google.Endpoint,
	}

	return &gmailOAuthClient{config}
}

func (c *gmailOAuthClient) AuthURL() string {
	return c.config.AuthCodeURL("state")
}

func (c *gmailOAuthClient) Exchange(code string) (*oauth2.Token, error) {
	return c.config.Exchange(context.Background(), code)
}
