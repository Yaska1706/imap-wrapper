package authconnector

import (
	"github.com/emersion/go-imap/v2/imapclient"
	"github.com/emersion/go-sasl"
)

type authConnector struct {
	imapAddress string
}

type AuthConnector interface {
	Connect(username string, accessToken string) (*imapclient.Client, error)
}

func NewAuthConnector(imapAddress string) *authConnector {
	return &authConnector{
		imapAddress: imapAddress,
	}
}

// Connect establishes a connection to the IMAP server using the provided username and access token for authentication.
func (ac *authConnector) Connect(username string, accessToken string) (*imapclient.Client, error) {

	imapClient, err := imapclient.DialTLS(ac.imapAddress, nil)
	if err != nil {
		return nil, err
	}
	saslOptions := &sasl.OAuthBearerOptions{
		Username: username,
		Token:    accessToken,
	}

	if err := imapClient.Authenticate(sasl.NewOAuthBearerClient(saslOptions)); err != nil {
		return nil, err
	}
	return imapClient, nil
}
