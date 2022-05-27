package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/jekabolt/solutions-dapp/art-admin/auth"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/rs/zerolog/log"
)

func (s *Server) auth(w http.ResponseWriter, r *http.Request) {
	ar := &auth.AuthRequest{}

	if err := render.Bind(r, ar); err != nil {
		log.Error().Err(err).Msgf("auth:render.Bind [%v]", err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if s.Auth.IsAuthorized(ar) {
		token, err := s.Auth.GetJWT()
		if err != nil {
			log.Error().Err(err).Msgf("auth:GetJWT [%v]", err.Error())
			render.Render(w, r, ErrInternalServerError(err))
			return
		}
		render.Render(w, r, NewAuthResponse(&AuthResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
		}))
		return
	}

	render.Render(w, r, ErrUnauthorizedError(fmt.Errorf("password or refresh token is invalid")))
}

func (s *Server) Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			log.Error().Err(err).Msgf("Authenticator:jwtauth.FromContext [%v]", err.Error())
			render.Render(w, r, ErrUnauthorizedError(err))
			return
		}

		if token == nil || jwt.Validate(token) != nil {
			log.Error().Err(err).Msgf("Authenticator:jwt.Validate [%v]", err.Error())
			render.Render(w, r, ErrUnauthorizedError(err))
			return
		}

		if token.Subject() == auth.RefreshTokenSub {
			render.Render(w, r, ErrUnauthorizedError(fmt.Errorf("use access token instead")))
			return
		}

		next.ServeHTTP(w, r)
	})
}
