package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/rs/zerolog/log"
)

func (s *Server) uploadOffchain(w http.ResponseWriter, r *http.Request) {

	mrs, err := s.nftStore.GetAllToUpload()
	if err != nil {
		log.Error().Err(err).Msgf("uploadIPFS:DB.GetAllToUpload [%v]", err.Error())
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// TODO: QUEUE THIS SHIT
	// TODO: QUEUE THIS SHIT
	// TODO: QUEUE THIS SHIT
	metadata, err := s.ipfs.BulkUploadIPFS(mrs)
	if err != nil {
		log.Error().Err(err).Msgf("uploadIPFS:ipfs.BulkUploadIPFS [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	for i := 0; i < s.Config.NFTTotalSupply; i++ {
		_, ok := metadata[i]
		if !ok {
			metadata[i] = bucket.Metadata{
				Name:        s.descs.GetCollectionName(i),
				Description: s.descs.GetDescription(i),
				Image:       s.descs.GetImage(i),
				Edition:     i,
				Date:        time.Now().Unix(),
			}
		}
	}

	url, err := s.Bucket.UploadMetadata(metadata)
	if err != nil {
		log.Error().Err(err).Msgf("uploadIPFS:ipfs.BulkUploadIPFS [%v]", err.Error())
		render.Render(w, r, ErrInternalServerError(err))
		return
	}

	fmt.Println("url - ", url)
}

func (s *Server) uploadIPFS(w http.ResponseWriter, r *http.Request) {

}
