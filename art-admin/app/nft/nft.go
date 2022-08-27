package nft

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/bunt"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft    *pb_nft.UnimplementedNftServer
	db     bunt.Store
	bucket bucket.FileStore
	ipfs   ipfs.IPFS
	descs  *descriptions.Store
	c      *Config
}
type Config struct {
	NFTTotalSupply int `env:"NFT_TOTAL_SUPPLY" envDefault:"1000"`
}

func (c *Config) New(
	db bunt.Store,
	bucket bucket.FileStore,
	ipfs ipfs.IPFS,
	descs *descriptions.Store,
) (*Server, error) {
	s := &Server{
		c:      c,
		db:     db,
		bucket: bucket,
		ipfs:   ipfs,
		descs:  descs,
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

func (s *Server) UpsertNFTMintRequest(ctx context.Context, req *pb_nft.NFTMintRequestToUpload) (*pb_nft.NFTMintRequestWithStatus, error) {
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

	mr, err := s.db.UpsertNFTMintRequest(req, sampleImages)
	if err != nil {
		log.Error().Err(err).Msgf("upsertNFTMintRequest:s.DB.UpsertNFTMintRequest [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert mint request: %s", err.Error())
	}
	return mr, nil
}

func (s *Server) ListNFTMintRequests(ctx context.Context, _ *emptypb.Empty) (*pb_nft.NFTMintRequestListArray, error) {
	nftMRs, err := s.db.GetAllNFTMintRequests()
	if err != nil {
		log.Error().Err(err).Msgf("ListNFTMintRequests:s.DB.GetAllNFTMintRequests [%v]", err.Error())
		return nil, fmt.Errorf("cannot upsert list mint request: %s", err.Error())
	}
	return &pb_nft.NFTMintRequestListArray{
		NftMintRequests: nftMRs,
	}, err
}

func (s *Server) DeleteNFTMintRequestById(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.DeleteStatus, error) {
	if err := s.db.DeleteNFTMintRequestById(req.Id); err != nil {
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
	nftMR, err := s.db.UpdateNFTOffchainUrl(req.Id, img.FullSize)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.UpdateNFTOffchainUrl [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	return nftMR, err
}

func (s *Server) DeleteNFTOffchainUrl(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.NFTMintRequestWithStatus, error) {
	nftMR, err := s.db.DeleteNFTOffchainUrl(req.Id)
	if err != nil {
		log.Error().Err(err).Msgf("DeleteNFOffchainUrl:s.db.DeleteNFTMintRequestById [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	return nftMR, err
}

func (s *Server) UploadOffchainMetadata(ctx context.Context, _ *emptypb.Empty) (*pb_nft.MetadataOffchainUrl, error) {
	mrs, err := s.db.GetAllToUpload()
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:s.db.GetAllToUpload [%v]", err.Error())
		return nil, fmt.Errorf("cannot update get all to upload: %s", err.Error())
	}
	if len(mrs) == 0 {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:nothing to upload ")
		return nil, fmt.Errorf("nothing to upload")
	}

	// TODO: QUEUE THIS SHIT
	// TODO: QUEUE THIS SHIT
	// TODO: QUEUE THIS SHIT
	metadata, err := s.ipfs.BulkUploadIPFS(mrs)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:ipfs.BulkUploadIPFS [%v]", err.Error())
		return nil, fmt.Errorf("can't upload to ipfs")
	}

	for i := 0; i < s.c.NFTTotalSupply; i++ {
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

	url, err := s.bucket.UploadMetadata(metadata)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:s.bucket.UploadMetadata [%v]", err.Error())
		return nil, fmt.Errorf("can't upload metadata to bucket")
	}

	err = s.db.AddOffchainMetadata(url)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:AddOffchainMetadata [%v]", err.Error())
		return nil, fmt.Errorf("can't save metadata to db")
	}

	return &pb_nft.MetadataOffchainUrl{
		Url: url,
	}, nil
}

func (s *Server) Burn(ctx context.Context, req *pb_nft.BurnRequest) (*emptypb.Empty, error) {
	err := s.db.BurnNft(req)
	if err != nil {
		log.Error().Err(err).Msgf("Burn:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot submit burn data: %s", err.Error())
	}
	return nil, err
}

func (s *Server) GetAllBurned(ctx context.Context, _ *emptypb.Empty) (*pb_nft.BurnList, error) {
	all, err := s.db.GetBurned()
	if err != nil {
		log.Error().Err(err).Msgf("GetAllBurned:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot get burn data: %s", err.Error())
	}
	return &pb_nft.BurnList{
		Data: all,
	}, nil
}

func (s *Server) GetAllBurnedPending(ctx context.Context, _ *emptypb.Empty) (*pb_nft.BurnList, error) {
	pending, err := s.db.GetBurnedPending()
	if err != nil {
		log.Error().Err(err).Msgf("GetAllBurned:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot get burn data: %s", err.Error())
	}
	return &pb_nft.BurnList{
		Data: pending,
	}, nil
}

func (s *Server) GetAllBurnedError(ctx context.Context, _ *emptypb.Empty) (*pb_nft.BurnList, error) {
	errors, err := s.db.GetBurnedErrors()
	if err != nil {
		log.Error().Err(err).Msgf("GetAllBurned:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot get burn data: %s", err.Error())
	}
	return &pb_nft.BurnList{
		Data: errors,
	}, nil
}

func (s *Server) UpdateBurnShippingStatus(ctx context.Context, req *pb_nft.ShippingStatusUpdateRequest) (*emptypb.Empty, error) {
	err := s.db.UpdateShippingStatus(req)
	if err != nil {
		log.Error().Err(err).Msgf("GetAllBurned:s.db.BurnNft [%v]", err.Error())
		return nil, fmt.Errorf("cannot update shipping status: %s", err.Error())
	}
	return nil, err
}

// TODO: get metadata from
func (s *Server) UploadIPFSMetadata(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}
