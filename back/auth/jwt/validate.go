package jwt

import (
	"github.com/PatateDu609/matcha/auth/jwt/key"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

func ValidateToken(bytes []byte) (jwt.Token, error) {
	return jwt.Parse(bytes, jwt.WithKey(jwa.RS256, key.Private()))
}
