package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/jwtauth/v5"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/auth/jwt"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/auth/pwhash"
	pb_auth "github.com/jekabolt/solutions-dapp/art-admin/proto/auth"
)

// Server implements the heartbeat service.
type Server struct {
	pb_auth.UnimplementedAuthServer
	pwhash            *pwhash.PasswordHasher
	jwtAuth           *jwtauth.JWTAuth
	jwtTTL            time.Duration
	c                 *Config
	AdminPasswordHash string
}

type Config struct {
	JWTSecret                string `env:"AUTH_JWT_SECRET" envDefault:"hehe"`
	AdminPassword            string `env:"AUTH_ADMIN_PASSWORD" envDefault:"hehe"`
	PasswordHasherSaltSize   int    `env:"AUTH_PASSWORD_HASHER_SALT" envDefault:"16"`
	PasswordHasherIterations int    `env:"AUTH_PASSWORD_HASHER_ITERATIONS" envDefault:"100000"`
	JWTTTL                   string `env:"AUTH_JWT_TTL" envDefault:"60m"` // in minutes
}

// New creates a new auth server.
func (c *Config) New() (*Server, error) {
	ttl, err := time.ParseDuration(c.JWTTTL)
	if err != nil && c.JWTTTL != "" {
		return nil, fmt.Errorf("bad duration for ttl [%s] - %v", c.JWTTTL, err.Error())
	}
	phasher, err := pwhash.New(c.PasswordHasherSaltSize, c.PasswordHasherIterations)
	if err != nil {
		return nil, fmt.Errorf("cannot create password hasher: %v", err)
	}
	adminPwHash, err := phasher.HashPassword(c.AdminPassword)
	if err != nil {
		return nil, fmt.Errorf("cannot hash admin password: %v", err)
	}
	s := &Server{
		pwhash:            phasher,
		jwtAuth:           jwtauth.New("HS256", []byte(c.JWTSecret), nil),
		jwtTTL:            ttl,
		c:                 c,
		AdminPasswordHash: adminPwHash,
	}

	return s, nil
}

func (s *Server) Login(ctx context.Context, req *pb_auth.LoginRequest) (*pb_auth.LoginResponse, error) {

	if err := s.pwhash.Validate(req.Password, s.AdminPasswordHash); err != nil {
		return nil, fmt.Errorf("cannot validate password: %s", err.Error())
	}

	token, err := jwt.NewToken(s.jwtAuth, s.jwtTTL)
	if err != nil {
		return nil, fmt.Errorf("cannot create token: %s", err.Error())
	}

	return &pb_auth.LoginResponse{
		AuthToken: token,
	}, nil
}

// CheckAuth middleware checks if the user is authenticated.
func (s *Server) CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := strings.TrimPrefix(r.Header.Get("Grpc-Metadata-Authorization"), "Bearer ")
		_, err := jwt.VerifyToken(s.jwtAuth, token)
		if err != nil {
			http.Error(w, fmt.Sprintf("invalid token %v", err.Error()), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
