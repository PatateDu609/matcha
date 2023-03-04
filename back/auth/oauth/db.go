package oauth

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"github.com/PatateDu609/matcha/auth/jwt/key"
	"github.com/PatateDu609/matcha/auth/session"
	"github.com/PatateDu609/matcha/utils/database"
	"github.com/PatateDu609/matcha/utils/log"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type Provider string

func (pder *Provider) Value() (driver.Value, error) {
	return string(*pder), nil
}

func (pder *Provider) Scan(src any) error {
	str, ok := src.(string)
	if !ok {
		return fmt.Errorf("couldn't interpret source value")
	}
	*pder = Provider(str)
	return nil
}

const (
	GOOGLE   Provider = "google"
	DISCORD           = "discord"
	GITHUB            = "github"
	SCHOOL42          = "42"
)

type Table struct {
	ID           uuid.UUID
	UserID       uuid.NullUUID
	State        uuid.NullUUID
	Provider     Provider
	AccessToken  *string
	RefreshToken *string
	Expiration   *time.Time
}

func (auth Table) GetName() string {
	return "user_oauth"
}

func (auth Table) GetAlias() string {
	return "oauth"
}

func (auth Table) GetColumns() []string {
	return database.GetColumns[Table](true)
}

func (auth Table) GetMandatoryColumns() []string {
	return database.GetColumns[Table](false)
}

func (auth Table) PrepareInsertion() ([]string, []any) {
	columns := auth.GetColumns()
	return columns, database.GetInterfaceArray(&auth)
}

func UpdateToken(ctx context.Context, auth *Table, tok *oauth2.Token) error {
	idTokenString, ok := tok.Extra("id_token").(string)
	if !ok {
		return fmt.Errorf("couldn't read id_token")
	}

	idToken, err := jwt.ParseString(idTokenString, jwt.WithKeySet(key.GoogleKeySet, jws.WithInferAlgorithmFromKey(true)))
	if err != nil {
		return fmt.Errorf("couldn't parse id_token: %s", err)
	}

	claims := idToken.PrivateClaims()
	log.Logger.Infof("id token claims: %+v", claims)

	patch := database.Patch{
		"access_token":  tok.AccessToken,
		"refresh_token": tok.RefreshToken,
		"expiration":    tok.Expiry,
	}

	cond := database.NewCondition("id", database.EqualTo, auth.ID)

	_, err = database.Update[Table](ctx, patch, cond)
	if err != nil {
		return err
	}
	return nil
}

func getUserIDFromSession(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	sess, err := session.GlobalSessions.SessionStart(w, r)
	if err != nil {
		if err := session.GlobalSessions.SessionDestroy(w, r); err != nil {
			return uuid.UUID{}, fmt.Errorf("couldn't destroy invalid session: %s", err)
		}
		return uuid.UUID{}, fmt.Errorf("session cookie is set but can't start session: %s", err)
	}

	id, ok := sess.Get("uuid").(string)
	if !ok {
		if err := session.GlobalSessions.SessionDestroy(w, r); err != nil {
			return uuid.UUID{}, fmt.Errorf("couldn't destroy invalid session: %s", err)
		}
		return uuid.UUID{}, fmt.Errorf("session is invalid, no uuid is set")
	}

	var parsedID uuid.UUID
	if parsedID, err = uuid.Parse(id); err != nil {
		if err := session.GlobalSessions.SessionDestroy(w, r); err != nil {
			return uuid.UUID{}, fmt.Errorf("couldn't destroy invalid session: %s", err)
		}
		return uuid.UUID{}, fmt.Errorf("couldn't parse the uuid stored in session: %s", err)
	}

	return parsedID, nil
}

func Create(w http.ResponseWriter, r *http.Request, provider Provider) (*Table, error) {
	var err error
	state, err := GetRandomState()
	if err != nil {
		return nil, err
	}

	auth := &Table{
		Provider: provider,
		UserID: uuid.NullUUID{
			Valid: false,
		},
		State: uuid.NullUUID{
			UUID:  state,
			Valid: true,
		},
		AccessToken:  nil,
		RefreshToken: nil,
		Expiration:   nil,
	}

	auth.ID, err = uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if session.GlobalSessions.CheckSessionCookie(r) {
		userID, err := getUserIDFromSession(w, r) // if user is logged in, we can bind it
		if err != nil {
			return nil, err
		}
		auth.UserID.UUID = userID
		auth.UserID.Valid = true
	}

	if err = database.Insert(r.Context(), auth); err != nil {
		return nil, fmt.Errorf("couldn't insert oauth for user: %s", err)
	}

	return auth, nil
}

func GetAuthByState(r *http.Request, state string) (*Table, error) {
	if _, err := uuid.Parse(state); err != nil {
		return nil, fmt.Errorf("couldn't parse state as uuid: %s", err)
	}

	cond := database.NewCondition("state", database.EqualTo, state)
	rows, err := database.Select[Table](r.Context(), cond)
	if err != nil {
		return nil, fmt.Errorf("couldn't get auth table: %s", err)
	}
	if len(rows) != 1 {
		return nil, fmt.Errorf("got wrong number of auth instances")
	}
	return &rows[0], nil
}
