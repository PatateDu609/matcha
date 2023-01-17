package jwt

import (
	"fmt"

	"github.com/PatateDu609/matcha/auth/jwt/key"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func ValidateToken(bytes []byte) (jwt.Token, error) {
	tok, err := jwt.Parse(bytes, jwt.WithKey(jwa.RS256, key.Private()), jwt.WithValidate(true), jwt.WithVerify(true))
	if err != nil {
		return nil, err
	}

	if tok.Issuer() != issuer {
		return nil, fmt.Errorf("wrong issuer")
	}

	return tok, nil
}
