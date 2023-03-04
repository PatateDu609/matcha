package key

import (
	"context"
	"log"

	"github.com/lestrrat-go/jwx/v2/jwk"
)

const googleKeyURL = "https://www.googleapis.com/oauth2/v3/certs"

var GoogleKeySet jwk.Set

func init() {
	set, err := jwk.Fetch(context.Background(), googleKeyURL)
	if err != nil {
		log.Printf("failed to parse JWK: %s", err)
		return
	}
	GoogleKeySet = set
}
