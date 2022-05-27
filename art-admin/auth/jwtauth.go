package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

const (
	RefreshTokenSub = "refresh"
)

type Config struct {
	JWTSecret   string `env:"JWT_SECRET" envDefault:"kek"`
	AdminSecret string `env:"ADMIN_SECRET" envDefault:"kek"`
}

type Auth struct {
	JWTAuth *jwtauth.JWTAuth
	*Config
}

func (c *Config) New() *Auth {
	return &Auth{
		Config:  c,
		JWTAuth: jwtauth.New("HS256", []byte(c.JWTSecret), nil),
	}
}

type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthRequest struct {
	Password     string `json:"password"`
	RefreshToken string `json:"refreshToken,omitempty"`
}

func (a *AuthRequest) Bind(r *http.Request) error {
	return a.Validate()
}

func (a *AuthRequest) Validate() error {
	if a.Password == "" && a.RefreshToken == "" {
		return fmt.Errorf("nor password and refresh token was send")
	}
	return nil
}

func (a *Auth) CheckAdminSecret(ar *AuthRequest) error {
	if a.AdminSecret != ar.Password {
		return fmt.Errorf("password not match")
	}
	return nil
}

func (a *Auth) IsAuthorized(ar *AuthRequest) bool {
	if ar.Password != "" {
		err := a.CheckAdminSecret(ar)
		return err == nil
	}

	if ar.RefreshToken != "" {
		rt, err := jwtauth.VerifyToken(a.JWTAuth, ar.RefreshToken)
		if err != nil {
			return false
		}
		if rt.Subject() == RefreshTokenSub {
			return true
		}
		return false
	}
	return false
}

func (a *Auth) GetJWT() (*AuthToken, error) {
	_, ts, err := a.JWTAuth.Encode(map[string]interface{}{
		"iss": "backend.grbpwr.com",
		"exp": time.Now().Add(time.Hour * 15).Unix(),
	})
	if err != nil {
		return nil, err
	}

	_, rts, err := a.JWTAuth.Encode(map[string]interface{}{
		"sub": RefreshTokenSub,
		"exp": time.Now().Add((time.Hour * 24) * 5).Unix(),
	})
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken:  ts,
		RefreshToken: rts,
	}, nil
}
