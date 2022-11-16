package nft

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft    *pb_nft.UnimplementedNftServer
	db     redis.Store
	bucket bucket.FileStore
	c      *Config
}
type Config struct {
	NFTTotalSupply int `env:"NFT_TOTAL_SUPPLY" envDefault:"1000"`
}

func (c *Config) New(
	db redis.Store,
	bucket bucket.FileStore,
) (*Server, error) {
	s := &Server{
		c:      c,
		db:     db,
		bucket: bucket,
	}

	return s, nil
}

func (s *Server) UploadRawImage(ito *pb_nft.ImageToUpload, pe *bucket.PathExtra) (*pb_nft.ImageList, error) {
	switch {
	case ito.Raw == "":
		log.Error().Msgf("UploadRawImage empty raw image")
		return nil, fmt.Errorf("cannot upload raw image: raw is empty")
	case strings.Contains(ito.Raw, "https://"), strings.Contains(ito.Raw, "https://"):
		return &pb_nft.ImageList{
			FullSize:   ito.Raw,
			Compressed: ito.Raw,
		}, nil
	default:
		img, err := s.bucket.UploadContentImage(ito.Raw, pe)
		if err != nil {
			log.Error().Err(err).Msgf("UploadRawImage:s.Bucket.UploadContentImage [%v]", err.Error())
			return nil, fmt.Errorf("cannot upload raw image: %s", err.Error())
		}
		return img, err
	}
}

func (s *Server) NewNFTMintRequest(ctx context.Context, req *pb_nft.NFTMintRequestToUpload) (*pb_nft.NFTMintRequestWithStatus, error) {
	sampleImages := []*pb_nft.ImageList{}
	pe := &bucket.PathExtra{
		TxHash:  req.NftMintRequest.TxHash,
		EthAddr: req.NftMintRequest.EthAddress,
	}
	for i, si := range req.SampleImages {
		pe.MintSequence = fmt.Sprintf("%d-%d", req.NftMintRequest.MintSequenceNumber, i)
		img, err := s.UploadRawImage(si, pe)
		if err != nil {
			return nil, err
		}
		sampleImages = append(sampleImages, img)
	}

	mr, err := s.db.NewNFTMintRequest(ctx, req, sampleImages)
	if err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:s.DB.UpsertNFTMintRequest [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert mint request: %s", err.Error())
	}
	return mr, nil
}

func (s *Server) ListNFTMintRequestsPaged(ctx context.Context, req *pb_nft.ListPagedRequest) (*pb_nft.NFTMintRequestListArray, error) {
	nftMRs, err := s.db.GetNFTMintRequestsPaged(ctx, req)
	if err != nil {
		log.Error().Err(err).Msgf("ListNFTMintRequests:s.DB.GetAllNFTMintRequests [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert list mint request: %s", err.Error())
	}
	return &pb_nft.NFTMintRequestListArray{
		NftMintRequests: nftMRs,
	}, err
}

func (s *Server) DeleteNFTMintRequestById(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.DeleteStatus, error) {
	if err := s.db.DeleteNFTMintRequestById(ctx, req.Id); err != nil {
		log.Error().Err(err).Msgf("DeleteNFTMintRequestById:s.db.DeleteNFTMintRequestById [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert delete mint request: %s", err.Error())
	}
	return &pb_nft.DeleteStatus{
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (s *Server) UpdateNFTOffchainUrl(ctx context.Context, req *pb_nft.UpdateNFTOffchainUrlRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	img, err := s.UploadRawImage(req.NftOffchainUrl, nil)
	if err != nil {
		return nil, err
	}
	nftMR, err := s.db.UpdateOffchainUrl(ctx, req.Id, img.FullSize)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.UpdateNFTOffchainUrl [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	return nftMR, err
}

func (s *Server) DeleteNFTOffchainUrl(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.NFTMintRequestWithStatus, error) {
	nftMR, err := s.db.DeleteNFTOffchainUrl(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msgf("DeleteNFOffchainUrl:s.db.DeleteNFTMintRequestById [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	return nftMR, err
}

func (s *Server) Burn(ctx context.Context, req *pb_nft.BurnRequest) (*emptypb.Empty, error) {
	_, err := s.db.UpdateShippingInfo(ctx, req)
	if err != nil {
		log.Error().Err(err).Msgf("Burn:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot submit burn data: %s", err.Error())
	}
	return nil, err
}

func (s *Server) SetTrackingNumber(ctx context.Context, req *pb_nft.SetTrackingNumberRequest) (*emptypb.Empty, error) {
	_, err := s.db.UpdateTrackingNumber(ctx, req)
	if err != nil {
		log.Error().Err(err).Msgf("GetAllBurned:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot get burn data: %s", err.Error())
	}
	return nil, err
}
