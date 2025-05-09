package jwtutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenCfg struct {
	Secret string
	TTL    time.Duration
}

// TODO: abstract claims to separate util for more generic token generation
func Sign(userId int32, cfg TokenCfg) (string, error) {
	claims := jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(cfg.TTL).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	signedToken, err := token.SignedString([]byte(cfg.Secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func Parse(tokenStr, secret string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return nil, nil, err
	}

	claims, _ := token.Claims.(jwt.MapClaims)

	return token, claims, nil
}
