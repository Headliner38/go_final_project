package auth

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/*func generateRSAKey() (*rsa.PrivateKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}
	return privateKey, nil
}*/

func GenerateToken(password string) (string, error) {
	secretKey := GetPrivateKey()
	passwordHash := sha256.Sum256([]byte(password))
	claims := jwt.MapClaims{
		"pwd_hash": hex.EncodeToString(passwordHash[:]),
		"exp":      time.Now().Add(8 * time.Hour).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	fmt.Println(ValidateToken(signedToken, password))

	return signedToken, nil
}

func ValidateToken(tokenString, currentPassword string) (bool, error) {
	passwordHash := sha256.Sum256([]byte(currentPassword))
	expectedHash := hex.EncodeToString(passwordHash[:])

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return GetPublicKey(), nil
	})
	if err != nil {
		return false, fmt.Errorf("ошибка парсинга токена: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if hash, ok := claims["pwd_hash"].(string); ok {
			return hash == expectedHash, nil
		}
	}
	return false, nil
}
