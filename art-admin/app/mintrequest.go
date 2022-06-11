package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
	"github.com/rs/zerolog/log"
)

func (s *Server) deleteNFTMintRequestById(w http.ResponseWriter, r *http.Request) {
	sp := strings.Split(r.URL.Path, "/")
	id := sp[len(sp)-1]
	if err := s.nftStore.DeleteNFTMintRequestById(id); err != nil {
		log.Error().Err(err).Msgf("deleteNFTMintRequestById:s.DB.DeleteNFTMintRequestById [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
		return
	}
	render.Render(w, r, StatusOK)
}

func (s *Server) getAllNFTMintRequestsList(w http.ResponseWriter, r *http.Request) {
	nftMRs, err := s.nftStore.GetAllNFTMintRequests()
	if err != nil {
		log.Error().Err(err).Msgf("getAllNFTMintRequestsList:s.DB.GetAllNFTMintRequests [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
		return
	}
	if err := render.RenderList(w, r, NewNFTMintRequestListResponse(nftMRs)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (s *Server) upsertNFTMintRequest(w http.ResponseWriter, r *http.Request) {
	req := &NFTMintRequestRequest{}
	if err := render.Bind(r, req); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:render.Bind [%v]", err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}
	if err := req.Validate(); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:render.Validate [%v]", err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	sampleImages := []bucket.Image{}
	for i, si := range req.SampleImages {
		if si.RawB64Image != nil &&
			*si.RawB64Image != "" &&
			!strings.Contains(si.FullSize, "https://") {

			img, err := s.Bucket.UploadContentImage(*si.RawB64Image, &bucket.PathExtra{
				TxHash:       req.TxHash,
				EthAddr:      req.ETHAddress,
				MintSequence: fmt.Sprintf("%d-%d", req.MintSequenceNumber, i),
			})
			if err != nil {
				log.Error().Err(err).Msgf("upsertNFTMintRequest:s.Bucket.UploadContentImage [%v]", err.Error())
				render.Render(w, r, ErrInternalServerError(err))
				return
			}
			sampleImages = append(sampleImages, *img)
		} else {
			sampleImages = append(sampleImages, si)
		}
	}
	req.SampleImages = sampleImages

	if _, err := s.nftStore.UpsertNFTMintRequest(req.NFTMintRequest); err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:s.DB.UpsertNFTMintRequest [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
	}
	render.Render(w, r, NewNFTMintResponse(*req.NFTMintRequest))
}

type NFTMintRequestRequest struct {
	*nft.NFTMintRequest
}

func (p *NFTMintRequestRequest) Bind(r *http.Request) error {
	return p.Validate()
}
