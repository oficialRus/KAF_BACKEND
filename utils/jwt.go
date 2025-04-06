package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID uint   `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

var jwtkey = []byte("секретный ключ")

// /Генерация токена
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

var ErrInvalidToken = errors.New("некорректный токен")

func ParseToken(tokenStr string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return jwtkey, nil

	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	return claims, nil
}
