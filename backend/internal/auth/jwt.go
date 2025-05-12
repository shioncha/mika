package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(id string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"nbf":     time.Now().Unix(),
		"iss":     "mika.shioncha.com",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	privateKey, err := loadPrivateKey()
	if err != nil {
		return "", err
	}

	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	publicKey, err := loadPublicKey()
	if err != nil {
		return nil, err
	}
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func loadPrivateKey() (*rsa.PrivateKey, error) {
	b64 := os.Getenv("JWT_PRIVATE_KEY_BASE64")
	if b64 == "" {
		return nil, fmt.Errorf("environment variable JWT_PRIVATE_KEY_BASE64 is not set")
	}

	// Base64デコード
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 private key: %v", err)
	}

	// PEM形式の秘密鍵をパース
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(pemBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	return privateKey, nil
}

func loadPublicKey() (*rsa.PublicKey, error) {
	b64 := os.Getenv("JWT_PUBLIC_KEY_BASE64")
	pemBytes, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pemBytes)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
