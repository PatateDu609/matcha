package jwt

import (
	"errors"
	"time"

	"github.com/PatateDu609/matcha/auth/jwt/key"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
	errors2 "github.com/pkg/errors"
)

const issuer = "github.com/PatateDu609/matcha"

var (
	ErrGeneratingToken = errors.New("couldn't generate token")
)

// SignToken returns a token that claims the given data and validating
// incoming requests
func SignToken(data map[string]any) ([]byte, error) {
	now := time.Now()
	exp := now.Add(time.Hour)

	builder := jwt.NewBuilder()

	for k, val := range data {
		builder.Claim(k, val)
	}

	token, err := builder.
		Issuer(issuer).
		IssuedAt(now).
		Expiration(exp).
		Audience([]string{"users"}).
		Build()

	if err != nil {
		return nil, errors2.Wrap(err, ErrGeneratingToken.Error())
	}

	signed, err := jwt.Sign(token, jwt.WithKey(jwa.RS256, key.Private()))
	if err != nil {
		return nil, errors2.Wrap(err, errors2.Wrap(errors2.Errorf("couldn't sign token"), ErrGeneratingToken.Error()).Error())
	}

	return signed, nil
}
