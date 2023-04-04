package metadata

import (
	"context"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/metadata"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/mongo"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"

	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft   *pb_metadata.UnimplementedMetadataServer
	db    mongo.Store
	ipfs  ipfs.IPFS
	descs *descriptions.MintDescription
	md    *metadata.MetaManager
	c     *Config
}
type Config struct {
}

func (c *Config) New(
	db mongo.Store,
	ipfs ipfs.IPFS,
	descs *descriptions.MintDescription,
	md *metadata.MetaManager,
) *Server {
	s := &Server{
		c:     c,
		db:    db,
		ipfs:  ipfs,
		descs: descs,
		md:    md,
	}
	return s
}

func (s *Server) UploadOffchainMetadata(ctx context.Context, _ emptypb.Empty) (*pb_metadata.UploadOffchainMetadataResponse, error) {
	md, err := s.db.GetOffchainMetadata(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:s.db.GetOffchainMetadata [%v]", err.Error())
		return nil, fmt.Errorf("cannot get offchain metadata: %s", err.Error())
	}
	if md.MetaInfo.IpfsUrl == "" {
		s.md.UploadInitial()
	}

}
func (s *Server) UploadIPFSMetadata(ctx context.Context, in *pb_metadata.UploadIPFSMetadataRequest) (*emptypb.Empty, error) {

}
func (s *Server) DeleteIPFSMetadata(ctx context.Context, in *pb_metadata.DeleteIPFSMetadataRequest) (*emptypb.Empty, error) {

}
func (s *Server) GetMetadata(ctx context.Context, in *pb_metadata.GetMetadataRequest) (*GetMetadataResponse, error) {

}

func (s *Server) GetAllMetadata(ctx context.Context, _ *emptypb.Empty) (*pb_metadata.GetAllMetadataResponse, error) {
	mds, err := s.db.GetAllOffchainMetadata(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("GetAllMetadata:s.db.GetAllOffchainMetadata [%v]", err.Error())
		return nil, fmt.Errorf("cannot get all metadata: %s", err.Error())
	}
	mi := []*pb_metadata.MetaInfo{}
	for _, md := range mds {
		mi = append(mi, &pb_metadata.MetaInfo{
			IpfsUrl:    md.IPFSUrl,
			Processing: md.Processing,
			Ts:         md.Ts,
			Key:        md.Key,
		})
	}
	return &pb_metadata.GetAllMetadataResponse{
		MetaInfo: mi,
	}, nil
}

func (s *Server) UploadOffchainMetadata(ctx context.Context, _ *emptypb.Empty) (*pb_metadata.UploadOffchainMetadataResponse, error) {
	toU, err := s.db.GetAllToUpload(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:s.db.GetAllToUpload [%v]", err.Error())
		return nil, fmt.Errorf("cannot update get all to upload: %s", err.Error())
	}
	if len(toU) == 0 {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:nothing to upload ")
		return nil, fmt.Errorf("nothing to upload")
	}
	md := []*pb_metadata.MetadataUnit{}
	for _, mr := range toU {
		md = append(md, &pb_metadata.MetadataUnit{
			Name:               fmt.Sprintf("solutions art %d", mr.NftMintRequest.MintSequenceNumber),
			Description:        s.descs.GetDescription(int(mr.NftMintRequest.MintSequenceNumber)),
			OffchainImage:      mr.OffchainUrl,
			OnchainImage:       mr.OnchainUrl,
			Edition:            s.descs.GetEditionBySequence(int(mr.NftMintRequest.MintSequenceNumber)),
			MintSequenceNumber: mr.NftMintRequest.MintSequenceNumber,
		})
	}
	key, err := s.db.AddOffchainMetadata(ctx, md)
	if err != nil {
		log.Error().Err(err).Msgf("UploadOffchainMetadata:s.db.AddOffchainMetadata [%v]", err.Error())
		return nil, fmt.Errorf("cannot save metadata url: %s", err.Error())
	}

	return &pb_metadata.UploadOffchainMetadataResponse{
		Key: key,
	}, nil
}

func (s *Server) DeleteIPFSMetadata(ctx context.Context, req *pb_metadata.DeleteIPFSMetadataRequest) (*emptypb.Empty, error) {
	err := s.db.DeleteById(ctx, req.Key)
	if err != nil {
		log.Error().Err(err).Msgf("DeleteIPFSMetadata:s.db.DeleteById [%v]", err.Error())
		return nil, fmt.Errorf("cannot save metadata url: %s", err.Error())
	}
	return nil, nil
}

// TODO: get metadata from
func (s *Server) UploadIPFSMetadata(ctx context.Context, req *pb_metadata.UploadIPFSMetadataRequest) (*emptypb.Empty, error) {
	// md, err := s.db.GetMetadataByKey(ctx, req.GetKey())
	// if err != nil {
	// 	log.Error().Err(err).Msgf("UploadOffchainMetadata:s.db.AddOffchainMetadata [%v]", err.Error())
	// 	return nil, fmt.Errorf("cannot save metadata url: %s", err.Error())
	// }

	return nil, nil
	//TODO: get _metadata.json from given id
	// mark metadata as processing
	// upload to ipfs
	// set finish processing
	// set ipfs url
}
