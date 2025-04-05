package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

var jwtkey = []byte("секретный ключ")

func GenerateToken(UserID uint, tokenType string, duration time.Duration) (string, error) {
	claims := TokenClaims{
		UserID: UserID,
		Type:   tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtkey)

}
