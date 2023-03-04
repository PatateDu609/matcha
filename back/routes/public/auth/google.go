package auth

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/PatateDu609/matcha/auth/oauth"
	"github.com/PatateDu609/matcha/config"
	"github.com/PatateDu609/matcha/routes/payloads"
	oauthPayloads "github.com/PatateDu609/matcha/routes/payloads/oauth"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/sirupsen/logrus"
)

func Google(w http.ResponseWriter, r *http.Request) {
	conf := config.Conf.OAuth.GetGoogleConfig()

	if conf == nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	auth, err := oauth.Create(w, r, oauth.GOOGLE)
	if err != nil {
		log.Logger.Errorf("couldn't create new auth in database: %s", err)
		w.WriteHeader(500)
		return
	}

	url := conf.AuthCodeURL(auth.State.UUID.String(), oauth2.AccessTypeOffline)
	log.Logger.Infof("google url: %s", url)

	if nb, err := w.Write([]byte(url)); err != nil || nb != len(url) {
		errStr := ""
		if err != nil {
			errStr = fmt.Sprintf(": %s", err)
		}
		log.Logger.Errorf("couldn't send google url to front%s", errStr)
	}
}

func GoogleRedirect(w http.ResponseWriter, r *http.Request) {
	conf := config.Conf.OAuth.GetGoogleConfig()

	if conf == nil {
		w.WriteHeader(http.StatusNotImplemented)
		return
	}

	payload, err := payloads.Unmarshal[oauthPayloads.GooglePayload](r.Body)
	if err != nil {
		log.Logger.Errorf("couldn't get oauth query params: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if payload.Code == "" {
		log.Logger.Errorf("no code found in response")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	auth, err := oauth.GetAuthByState(r, payload.State)
	if err != nil {
		log.Logger.Errorf("couldn't get auth entry by state: %s", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	if auth == nil || auth.Provider != oauth.GOOGLE {
		log.Logger.Errorf("auth entry not found for given state")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tok, err := conf.Exchange(r.Context(), payload.Code)
	if err != nil {
		log.Logger.Errorf("couldn't exchange code for token: %s", err)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	fields := logrus.Fields{
		"access":  tok.AccessToken,
		"refresh": tok.RefreshToken,
		"type":    tok.Type(),
		"expiry":  tok.Expiry.String(),
		"valid":   tok.Valid(),
	}
	log.Logger.WithFields(fields).Info()

	if err := oauth.UpdateToken(r.Context(), auth, tok); err != nil {
		log.Logger.Errorf("couldn't update token: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
