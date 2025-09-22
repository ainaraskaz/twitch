package auth

import "golang.org/x/oauth2"

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}
type TwitchAuth struct {
	OAuthConfig *oauth2.Config
}
