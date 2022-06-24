package jwt

import (
	"time"

	"github.com/go-chi/jwtauth/v5"
)

func VerifyToken(jwtAuth *jwtauth.JWTAuth, token string) (string, error) {
	t, err := jwtauth.VerifyToken(jwtAuth, token)
	if err != nil {
		return "", err
	}
	return t.Subject(), nil
}

func NewToken(jwtAuth *jwtauth.JWTAuth, ttl time.Duration) (string, error) {
	_, ts, err := jwtAuth.Encode(map[string]interface{}{
		"exp": time.Now().Add(ttl).Unix(),
	})
	if err != nil {
		return ts, err
	}
	return ts, nil
}
