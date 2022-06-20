package nft

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/eth"
	"github.com/jekabolt/solutions-dapp/art-admin/ipfs"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/jekabolt/solutions-dapp/art-admin/store/bunt"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft    *pb_nft.UnimplementedNftServer
	db     bunt.Store
	bucket *bucket.Bucket
	eth    *eth.Etherscan
	ipfs   *ipfs.Moralis
	descs  *descriptions.Store
	c      *Config
}
type Config struct {
	NFTTotalSupply int `env:"NFT_TOTAL_SUPPLY" envDefault:"1000"`
}

func (c *Config) New(db bunt.Store,
	bucket *bucket.Bucket,
	eth *eth.Etherscan,
	ipfs *ipfs.Moralis,
	descs *descriptions.Store) (*Server, error) {
	s := &Server{
		c:      c,
		db:     db,
		bucket: bucket,
		eth:    eth,
		ipfs:   ipfs,
		descs:  descs,
	}

	return s, nil
}

func (s *Server) upsertNFTMintRequest(ctx context.Context, req *pb_nft.NFTMintRequestToUpload) (*pb_nft.NFTMintRequestWithStatus, error) {

	sampleImages := []*pb_nft.ImageList{}
	for i, si := range req.SampleImages {
		if si.Raw != "" &&
			!strings.Contains(si.Raw, "https://") {

			img, err := s.bucket.UploadContentImage(si.Raw, &bucket.PathExtra{
				TxHash:       req.NftMintRequest.TxHash,
				EthAddr:      req.NftMintRequest.EthAddress,
				MintSequence: fmt.Sprintf("%d-%d", req.NftMintRequest.MintSequenceNumber, i),
			})
			if err != nil {
				log.Error().Err(err).Msgf("upsertNFTMintRequest:s.Bucket.UploadContentImage [%v]", err.Error())
				return nil, fmt.Errorf("cannot upload raw image: %s", err.Error())
			}
			sampleImages = append(sampleImages, img)
		} else {
			sampleImages = append(sampleImages, &pb_nft.ImageList{
				FullSize:   si.Raw,
				Compressed: si.Raw,
			})
		}
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
	nftMR, err := s.db.UpdateNFTOffchainUrl(req.Id, req.NftOffchainUrl)
	if err != nil {
		log.Error().Err(err).Msgf("UpdateNFTOffchainUrl:s.db.UpdateNFTOffchainUrl [%v]", err.Error())
		return nil, fmt.Errorf("cannot update nft offchain url request: %s", err.Error())
	}
	return nftMR, err
}

func (s *Server) DeleteNFOffchainUrl(ctx context.Context, req *pb_nft.DeleteId) (*pb_nft.NFTMintRequestWithStatus, error) {
	nftMR, err := s.db.DeleteNFOffchainUrl(req.Id)
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

// TODO: get metadata from
func (s *Server) UploadIPFSMetadata(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, nil
}
