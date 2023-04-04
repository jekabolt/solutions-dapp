package nft

import (
	"context"
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/mongo"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft    *pb_nft.UnimplementedNftServer
	db     mongo.Store
	bucket bucket.FileStore
	c      *Config
}
type Config struct {
	NFTTotalSupply int `env:"NFT_TOTAL_SUPPLY" envDefault:"1000"`
}

func (c *Config) New(
	db mongo.Store,
	bucket bucket.FileStore,
) (*Server, error) {
	s := &Server{
		c:      c,
		db:     db,
		bucket: bucket,
	}

	return s, nil
}

func (s *Server) NewNFTMintRequest(ctx context.Context, req *pb_nft.NFTMintRequestToUpload) (*pb_nft.NFTMintRequestWithStatus, error) {
	sampleImages := []*pb_nft.ImageList{}
	folder := path.Join("nft", "mint", req.EthAddress)
	for i, si := range req.SampleImages {
		img, err := s.uploadRawImage(si, folder, fmt.Sprintf("reference_%d", i))
		if err != nil {
			return nil, err
		}
		sampleImages = append(sampleImages, img)
	}
	mr, err := s.db.New(ctx, req, sampleImages)
	if err != nil {
		log.Error().Err(err).Msgf("NewNFTMintRequest:s.db.New [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert mint request: %s", err.Error())
	}
	return mr, nil
}

// helper function to upload raw image
func (s *Server) uploadRawImage(ito *pb_nft.ImageToUpload, folder string, imageName string) (*pb_nft.ImageList, error) {
	switch {
	case ito.Raw == "":
		log.Error().Msgf("uploadRawImage empty raw image")
		return nil, fmt.Errorf("cannot upload raw image: raw is empty")
	case strings.Contains(ito.Raw, "https://"), strings.Contains(ito.Raw, "https://"):
		return &pb_nft.ImageList{
			FullSize:   ito.Raw,
			Compressed: ito.Raw,
		}, nil
	default:
		img, err := s.bucket.UploadContentImage(ito.Raw, folder, imageName)
		if err != nil {
			log.Error().Err(err).Msgf("uploadRawImage:s.Bucket.UploadContentImage [%v]", err.Error())
			return nil, fmt.Errorf("cannot upload raw image: %s", err.Error())
		}
		return img, err
	}
}

func (s *Server) ListNFTMintRequestsPaged(ctx context.Context, req *pb_nft.ListPagedRequest) (*pb_nft.NFTMintRequestListArray, error) {
	nftMRs, err := s.db.GetPaged(ctx, req)
	if err != nil {
		log.Error().Err(err).Msgf("ListNFTMintRequestsPaged:s.db.GetPaged [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert list mint request: %s", err.Error())
	}
	return &pb_nft.NFTMintRequestListArray{
		NftMintRequests: nftMRs,
	}, err
}

func (s *Server) DeleteNFTMintRequestById(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.DeleteStatus, error) {
	if err := s.db.DeleteById(ctx, req.Id); err != nil {
		log.Error().Err(err).Msgf("DeleteNFTMintRequestById:s.db.DeleteById [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert delete mint request: %s", err.Error())
	}
	return &pb_nft.DeleteStatus{
		Message: http.StatusText(http.StatusOK),
	}, nil
}

func (s *Server) UpdateNFTOffchainUrl(ctx context.Context, req *pb_nft.UpdateNFTOffchainUrlRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	full, err := s.db.IsFull(ctx, req.CollectionId)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.IsFull [%v]", err.Error())
		return nil, fmt.Errorf("cannot update offchain url: %s", err.Error())
	}

	if full {
		return nil, fmt.Errorf("cannot update offchain url: collection is full")
	}
	nftMR, err := s.db.GetById(ctx, req.Id)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.GetById [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}

	folder := path.Join("nft", "mint", fmt.Sprintf("%d", nftMR.NftMintRequest.MintSequenceNumber))

	img, err := s.uploadRawImage(req.NftOffchainUrl, folder, "offchain")
	if err != nil {
		return nil, err
	}
	nftMR, err = s.db.UpdateOffchainUrl(ctx, req.Id, img.FullSize)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.UpdateNFTOffchainUrl [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	err = s.db.IncrementUsed(ctx, req.CollectionId)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.IncrementUsed [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	return nftMR, err
}

func (s *Server) DeleteNFTOnchainUrl(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.NFTMintRequestWithStatus, error) {
	nftMR, err := s.db.DeleteIpfsUrl(ctx, req.Id)
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
