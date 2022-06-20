package jwt

import (
	"testing"

	"github.com/go-chi/jwtauth/v5"
	"github.com/matryer/is"
)

func TestToken(t *testing.T) {
	is := is.New(t)

	jwtAuth := jwtauth.New("HS256", []byte("secret"), nil)
	tok, err := NewToken(jwtAuth, 10)
	is.NoErr(err)

	subToken, err := VerifyToken(jwtAuth, tok)
	is.NoErr(err)

	t.Log(subToken)

}
