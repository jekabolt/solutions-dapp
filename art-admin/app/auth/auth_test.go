package auth

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/proto/auth"
	"github.com/matryer/is"
)

const (
	JWTSecret     = "hehe"
	AdminPassword = "hehe"
)

func TestAuth(t *testing.T) {

	is := is.New(t)

	ctx := context.Background()
	c := Config{
		JWTSecret:                JWTSecret,
		AdminPassword:            AdminPassword,
		PasswordHasherSaltSize:   16,
		PasswordHasherIterations: 100000,
		JWTTTL:                   "60m",
	}
	authSrv, err := c.New()
	is.NoErr(err)

	resp, err := authSrv.Login(ctx, &auth.LoginRequest{
		Password: AdminPassword,
	})
	is.NoErr(err)

	token := fmt.Sprintf("Bearer %s", resp.AuthToken)

	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	handlerAuth := authSrv.CheckAuth(nextHandler)

	// create a mock request to use
	req := httptest.NewRequest("GET", "http://testing", nil)
	req.Header.Set("Grpc-Metadata-Authorization", token)

	rec := httptest.NewRecorder()
	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerAuth.ServeHTTP(rec, req)
	is.Equal(rec.Body.String(), "OK")
	is.Equal(rec.Code, http.StatusOK)

	// bad token case
	req.Header.Set("Grpc-Metadata-Authorization", "bad token")
	rec = httptest.NewRecorder()
	// call the handler using a mock response recorder (we'll not use that anyway)
	handlerAuth.ServeHTTP(rec, req)
	is.Equal(rec.Code, http.StatusUnauthorized)

}
