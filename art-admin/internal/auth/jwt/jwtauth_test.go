package jwt

import (
	"testing"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/matryer/is"
)

func TestToken(t *testing.T) {
	is := is.New(t)

	jwtAuth := jwtauth.New("HS256", []byte("secret"), nil)
	tok, err := NewToken(jwtAuth, time.Hour)
	is.NoErr(err)

	subToken, err := VerifyToken(jwtAuth, tok)
	is.NoErr(err)

	t.Log(subToken)

}
