package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func LoadKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := loadPrivateKey()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load private key: %v", err)
	}

	publicKey, err := loadPublicKey()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to load public key: %v", err)
	}

	return privateKey, publicKey, nil
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	b64 := os.Getenv("JWT_PRIVATE_KEY_BASE64")
	if b64 == "" {
		return nil, fmt.Errorf("environment variable JWT_PRIVATE_KEY_BASE64 is not set")
	}

	// Decode the base64 encoded private key
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 private key: %v", err)
	}

	// Parse the PEM bytes to get the private key
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pemBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	return privateKey, nil
}

func loadPublicKey() (*rsa.PublicKey, error) {
	b64 := os.Getenv("JWT_PUBLIC_KEY_BASE64")
	if b64 == "" {
		return nil, fmt.Errorf("environment variable JWT_PUBLIC_KEY_BASE64 is not set")
	}

	// Decode the base64 encoded public key
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 public key: %v", err)
	}

	// Parse the PEM bytes to get the public key
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pemBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	return publicKey, nil
}
