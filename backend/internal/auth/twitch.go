package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/twitch"
)

func NewTwitchAuth(cfg Config) *TwitchAuth {
	return &TwitchAuth{
		OAuthConfig: &oauth2.Config{
			ClientID:     cfg.ClientID,
			ClientSecret: cfg.ClientSecret,
			RedirectURL:  cfg.RedirectURI,
			Scopes:       []string{"user:read:email"},
			Endpoint:     twitch.Endpoint,
		},
	}
}
func (t *TwitchAuth) GetAuthURL(state string) string {
	return t.OAuthConfig.AuthCodeURL(state)
}
func (t *TwitchAuth) ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	return t.OAuthConfig.Exchange(ctx, code)
}
