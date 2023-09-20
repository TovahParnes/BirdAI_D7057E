package authentication

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

type IAuthentication interface {
	VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error)
}
type Authentication struct {
	*oidc.Provider
	oauth2.Config
}
