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
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft    *pb_nft.UnimplementedNftServer
	db     redis.Store
	bucket bucket.FileStore
	ipfs   ipfs.IPFS
	descs  *descriptions.Store
	c      *Config
}
type Config struct {
	NFTTotalSupply int `env:"NFT_TOTAL_SUPPLY" envDefault:"1000"`
}

func (c *Config) New(
	db redis.Store,
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
	nftMR, err := s.db.UpdateNFTOffchainUrl(ctx, req.Id, img.FullSize)
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

func (s *Server) GetAllMetadata(ctx context.Context, _ *emptypb.Empty) (*pb_nft.AllMetadataResponse, error) {
	metadata, err := s.db.GetAllOffchainMetadata(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("GetAllMetadata:s.db.GetAllOffchainMetadata [%v]", err.Error())
		return nil, fmt.Errorf("cannot get all metadata: %s", err.Error())
	}
	md := []*pb_nft.MetadataOffchainUrl{}
	for _, m := range metadata {
		md = append(md, &pb_nft.MetadataOffchainUrl{
			Url:        m.Url,
			IpfsUrl:    m.IPFSUrl,
			Ts:         int32(m.Ts),
			Processing: m.Processing,
		})
	}
	return &pb_nft.AllMetadataResponse{
		OffchainUrls: md,
	}, nil
}

func (s *Server) UploadOffchainMetadata(ctx context.Context, _ *emptypb.Empty) (*pb_nft.MetadataOffchainUrl, error) {
	mrs, err := s.db.GetAllToUpload(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:s.db.GetAllToUpload [%v]", err.Error())
		return nil, fmt.Errorf("cannot update get all to upload: %s", err.Error())
	}
	if len(mrs) == 0 {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:nothing to upload ")
		return nil, fmt.Errorf("nothing to upload")
	}

	//TODO: compose _metadata.json
	// upload to s3
	// set metadata url

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

	err = s.db.AddOffchainMetadata(ctx, url)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:AddOffchainMetadata [%v]", err.Error())
		return nil, fmt.Errorf("can't save metadata to db")
	}

	return &pb_nft.MetadataOffchainUrl{
		Url: url,
	}, nil
}

// TODO: get metadata from
func (s *Server) UploadIPFSMetadata(ctx context.Context, _ *pb_nft.UploadIPFSMetadataRequest) (*emptypb.Empty, error) {
	return nil, nil
	//TODO: get _metadata.json from given id
	// mark metadata as processing
	// upload to ipfs
	// set finish processing
	// set ipfs url
}
