package config

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"

	"github.com/PatateDu609/matcha/utils/log"
)

type OAuthKey struct {
	ClientID     string `mapstructure:"client-id"`
	ClientSecret string `mapstructure:"client-secret"`
}

type OAuth struct {
	Google   OAuthKey `mapstructure:"google"`
	Github   OAuthKey `mapstructure:"github"`
	Discord  OAuthKey `mapstructure:"discord"`
	School42 OAuthKey `mapstructure:"42school"`
}

//goland:noinspection HttpUrlsUsage
func (config OAuth) GetGoogleConfig() *oauth2.Config {
	redirectURL := fmt.Sprintf("http://%s/auth/redirect/google", Conf.API.FrontURL)
	if config.Google.ClientID == "" || config.Google.ClientSecret == "" {
		log.Logger.Errorln("credentials not set correctly, aborting config creation")
		return nil
	}

	return &oauth2.Config{
		ClientID:     config.Google.ClientID,
		ClientSecret: config.Google.ClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
	}
}

func (config OAuth) GetGithubConfig() *oauth2.Config {
	redirectURL := fmt.Sprintf("http://%s/auth/redirect/github", Conf.API.FrontURL)
	if config.Github.ClientID == "" || config.Github.ClientSecret == "" {
		log.Logger.Errorln("credentials not set correctly, aborting config creation")
		return nil
	}

	return &oauth2.Config{
		ClientID:     config.Github.ClientID,
		ClientSecret: config.Github.ClientSecret,
		Endpoint:     github.Endpoint,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"read:user",
			"email:user",
		},
	}
}

func (config OAuth) GetDiscordConfig() *oauth2.Config {
	redirectURL := fmt.Sprintf("http://%s/auth/redirect/github", Conf.API.FrontURL)
	if config.Discord.ClientID == "" || config.Discord.ClientSecret == "" {
		log.Logger.Errorln("credentials not set correctly, aborting config creation")
		return nil
	}

	return &oauth2.Config{
		ClientID:     config.Discord.ClientID,
		ClientSecret: config.Discord.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://discord.com/oauth2/authorize",
			TokenURL:  "https://discord.com/api/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: redirectURL,
		Scopes: []string{
			"read:user",
			"email:user",
		},
	}
}

func (config OAuth) Get42SchoolConfig() *oauth2.Config {
	redirectURL := fmt.Sprintf("http://%s/auth/redirect/42school", Conf.API.FrontURL)
	if config.School42.ClientID == "" || config.School42.ClientSecret == "" {
		log.Logger.Errorln("credentials not set correctly, aborting config creation")
		return nil
	}

	return &oauth2.Config{
		ClientID:     config.School42.ClientID,
		ClientSecret: config.School42.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://api.intra.42.fr/oauth/authorize",
			TokenURL:  "https://api.intra.42.fr/oauth/token",
			AuthStyle: oauth2.AuthStyleInParams,
		},
		RedirectURL: redirectURL,
		Scopes: []string{
			"read:user",
			"email:user",
		},
	}
}
