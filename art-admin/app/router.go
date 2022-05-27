package app

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/go-chi/jwtauth/v5"
	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (s *Server) Router() *chi.Mux {
	r := chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins: s.Config.Hosts,
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodDelete,
		},
		Debug: s.Config.Debug,
	})

	// Init middlewares
	r.Use(cors.Handler)
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Use(httprate.Limit(
		10,             // requests
		15*time.Second, // per duration
		httprate.WithLimitHandler(func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
		}),
	))

	fs := http.FileServer(http.Dir("./assets"))
	r.Handle("/assets/*", http.StripPrefix("/assets/", fs))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Post("/auth", s.auth)

	r.Route("/api", func(r chi.Router) {

		r.Put("/mint/requests", s.upsertNFTMintRequest)
		r.Get("/mint/requests", s.getAllNFTMintRequestsList) // public
		r.Delete("/mint/requests/{id}", s.deleteNFTMintRequestById)

		r.Put("/nft", s.upsertNFT)
		r.Delete("/nft/{id}", s.deleteNFT)

		r.Put("/upload/offchain", s.uploadOffchain)
		r.Put("/upload/ipfs", s.uploadIPFS)

		r.Group(func(r chi.Router) {

			r.Use(jwtauth.Verifier(s.Auth.JWTAuth))
			r.Use(s.Authenticator)

		})

	})

	return r
}

func (s *Server) Serve() error {
	log.Info().Msg("Listening on :" + s.Config.Port)
	return http.ListenAndServe(":"+s.Config.Port, s.Router())
}
