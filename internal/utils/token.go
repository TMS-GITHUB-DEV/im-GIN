package utils

import (
	"github.com/golang-jwt/jwt/v5"
	"im-GIN/config"
	"time"
)

type Claims struct {
	Subject uint64 `json:"sub"`
	jwt.RegisteredClaims
}

// GenerateAccessToken
// 生成access_token
func GenerateAccessToken(sub uint64, expires time.Duration) (string, error) {
	claims := &Claims{
		Subject: sub,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expires)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Cfg.App.Secret))
}

func GenerateRefreshToken(sub uint64) (string, error) {
	return "", nil
}

func ValidateAccessToken(accessToken string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Cfg.App.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
