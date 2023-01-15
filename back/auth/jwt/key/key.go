package key

import (
	"crypto/rsa"
	"crypto/x509"
	_ "embed"
	"encoding/pem"

	"github.com/PatateDu609/matcha/utils/log"
)

//go:embed private.pem
var rawKey []byte

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func init() {
	block, _ := pem.Decode(rawKey)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		log.Logger.Fatalf("couldn't load private key")
	}

	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Logger.Fatalf("couldn't parse private key: %s", err)
	}

	privateKey = private
	publicKey = &privateKey.PublicKey
}

func Private() *rsa.PrivateKey {
	return privateKey
}

func Public() *rsa.PublicKey {
	return publicKey
}
