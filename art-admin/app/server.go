package app

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"strings"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jekabolt/solutions-dapp/art-admin/app/auth"
	"github.com/jekabolt/solutions-dapp/art-admin/app/nft"
	pb_auth "github.com/jekabolt/solutions-dapp/art-admin/proto/auth"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/rs/zerolog/log"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Server struct {
	gs   *grpc.Server
	c    *Config
	done chan struct{}
}

var (
	//go:embed static
	fs embed.FS

	pages = map[string]string{
		"/nft/api": "static/swagger/index.html",
	}
)

type Config struct {
	Port int `env:"SERVER_PORT" envDefault:"3999"`
}

func (c *Config) Init() *Server {
	return &Server{
		gs:   grpc.NewServer(),
		c:    c,
		done: make(chan struct{}),
	}
}

// Stop stops the application and waits for all services to exit
func (s *Server) Stop(ctx context.Context) {
	s.gs.GracefulStop()
	close(s.done)
}

// Done returns a channel that is closed after the application has exited
func (s *Server) Done() chan struct{} {
	return s.done
}

// Start starts the server
func (s *Server) Start(ctx context.Context,
	authServer *auth.Server,
	nftServer *nft.Server,
) (err error) {

	s.gs = grpc.NewServer()
	pb_auth.RegisterAuthServer(s.gs, authServer)
	pb_nft.RegisterNftServer(s.gs, nftServer)

	var clientHTTPHandler http.Handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			s.gs.ServeHTTP(w, r)
		} else {
			if clientHTTPHandler == nil {
				w.WriteHeader(http.StatusNotImplemented)
				return
			}
			clientHTTPHandler.ServeHTTP(w, r)
		}
	})

	go func() {
		log.Ctx(ctx).Err(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", s.c.Port),
			h2c.NewHandler(handler, &http2.Server{}))).Msgf("error while ListenAndServe on port :%v", s.c.Port)
	}()

	clientHTTPHandler, err = s.setupHTTPAPI(authServer)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) setupHTTPAPI(auth *auth.Server) (http.Handler, error) {
	ctx := context.Background()
	r := chi.NewRouter()

	authHandler, err := s.authHandler(context.Background())
	if err != nil {
		return nil, err
	}

	airdropHandler, err := s.nftHandler(context.Background())
	if err != nil {
		return nil, err
	}

	r.HandleFunc("/nft/api", func(w http.ResponseWriter, r *http.Request) {
		page, ok := pages[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		tpl, err := template.ParseFS(fs, page)
		if err != nil {
			log.Ctx(ctx).Err(err).Msg("get swagger template error")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		if err := tpl.Execute(w, nil); err != nil {
			return
		}
	})

	r.Mount("/api/nft", auth.CheckAuth(airdropHandler))
	r.Mount("/api/auth", authHandler)

	log.Ctx(ctx).Info().Msg("api ok")

	r.Mount("/", http.FileServer(http.FS(fs)))

	return r, nil
}

func (s *Server) authHandler(ctx context.Context) (http.Handler, error) {
	apiEndpoint := fmt.Sprintf("0.0.0.0:%d", s.c.Port)

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			EnumsAsInts:  true,
			EmitDefaults: true,
		},
	))

	err := pb_auth.RegisterAuthHandlerFromEndpoint(ctx, mux, apiEndpoint,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return nil, err
	}

	return mux, nil
}

func (s *Server) nftHandler(ctx context.Context) (http.Handler, error) {
	apiEndpoint := fmt.Sprintf("0.0.0.0:%d", s.c.Port)

	mux := runtime.NewServeMux(runtime.WithMarshalerOption(
		runtime.MIMEWildcard,
		&runtime.JSONPb{
			EnumsAsInts:  true,
			EmitDefaults: true,
		},
	))

	err := pb_nft.RegisterNftHandlerFromEndpoint(ctx, mux, apiEndpoint,
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	if err != nil {
		return nil, err
	}

	return mux, nil
}
