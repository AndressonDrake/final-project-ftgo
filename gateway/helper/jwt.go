package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	SECRETKEY string
)

func InitSecret(secret string) {
	SECRETKEY = secret
}

func GenerateToken(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(SECRETKEY))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRETKEY), nil
	})
}
