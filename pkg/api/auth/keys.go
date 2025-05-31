package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

var (
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
)

func InitKeys() error {
	priv, err := generateRSAKey()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	privateKey = priv
	PublicKey = &priv.PublicKey
	return nil
}

func generateRSAKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return privateKey, nil
}

func GetPrivateKey() *rsa.PrivateKey {
	return privateKey
}

func GetPublicKey() *rsa.PublicKey {
	return PublicKey
}
