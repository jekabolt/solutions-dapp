package app

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/rs/zerolog/log"
)

func (s *Server) upsertNFT(w http.ResponseWriter, r *http.Request) {
	data := &NFTMintRequestRequest{}
	if err := render.Bind(r, data); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:render.Bind [%v]", err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err := data.Validate(); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:render.ValidateOffchain [%v]", err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	if _, err := s.DB.UpsertNFT(data.NFTMintRequest); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:s.DB.UpsertNFTMintRequest [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
		return
	}
	render.Render(w, r, NewNFTMintResponse(data.NFTMintRequest))
}

func (s *Server) deleteNFT(w http.ResponseWriter, r *http.Request) {
	sp := strings.Split(r.URL.Path, "/")
	id := sp[len(sp)-1]

	if _, err := s.DB.DeleteNFT(id); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:s.DB.UpsertNFTMintRequest [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
		return
	}
	render.Render(w, r, StatusOK)
}
