package collections

import (
	"context"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
	pb_collection "github.com/jekabolt/solutions-dapp/art-admin/proto/collection"
	"github.com/rs/zerolog/log"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Server struct {
	nft *pb_collection.UnimplementedCollectionsServer
	db  redis.CollectionsStore
	c   *Config
}
type Config struct {
}

func (c *Config) New(
	db redis.Store,
) *Server {
	s := &Server{
		c:  c,
		db: db,
	}
	return s
}

func (s *Server) CreateNewCollection(ctx context.Context, in *pb_collection.CreateNewCollectionRequest) (*pb_collection.CreateNewCollectionResponse, error) {
	k, err := s.db.AddCollection(ctx, in.GetName(), in.GetCapacity())
	if err != nil {
		log.Error().Err(err).Msgf("CreateNewCollection:s.db.AddCollection [%v]", err.Error())
		return nil, fmt.Errorf("collection not found by key %s: %s", k, err.Error())
	}
	return &pb_collection.CreateNewCollectionResponse{
		Key: k,
	}, nil
}
func (s *Server) DeleteCollection(ctx context.Context, in *pb_collection.DeleteCollectionRequest) (*pb_collection.DeleteCollectionResponse, error) {
	err := s.db.DeleteCollection(ctx, in.GetKey())
	if err != nil {
		log.Error().Err(err).Msgf("DeleteCollection:s.db.DeleteByKey [%v]", err.Error())
		return nil, fmt.Errorf("cannot delete collection %s : %s", in.GetKey(), err.Error())
	}
	return &pb_collection.DeleteCollectionResponse{
		Key: in.Key,
	}, nil
}
func (s *Server) UpdateCollectionCapacity(ctx context.Context, in *pb_collection.UpdateCollectionCapacityRequest) (*pb_collection.UpdateCollectionCapacityResponse, error) {
	err := s.db.UpdateCollectionCapacity(ctx, in.GetKey(), in.GetCapacity())
	if err != nil {
		log.Error().Err(err).Msgf("UpdateCollectionCapacity:s.db.UpdateCollectionCapacity [%v]", err.Error())
		return nil, fmt.Errorf("cannot update collection capacity collections: %s", err.Error())
	}
	return &pb_collection.UpdateCollectionCapacityResponse{
		Key: in.Key,
	}, nil
}
func (s *Server) UpdateCollectionName(ctx context.Context, in *pb_collection.UpdateCollectionNameRequest) (*pb_collection.UpdateCollectionCapacityResponse, error) {
	err := s.db.UpdateCollectionName(ctx, in.GetKey(), in.GetName())
	if err != nil {
		log.Error().Err(err).Msgf("UpdateCollectionName:s.db.UpdateCollectionName [%v]", err.Error())
		return nil, fmt.Errorf("cannot update collection name: %s", err.Error())
	}
	return &pb_collection.UpdateCollectionCapacityResponse{
		Key: in.Key,
	}, nil
}
func (s *Server) GetAllCollections(ctx context.Context, _ *emptypb.Empty) (*pb_collection.GetAllCollectionsResponse, error) {
	cs, err := s.db.GetAllCollections(ctx)
	if err != nil {
		log.Error().Err(err).Msgf("GetAllCollections:s.db.GetAllCollections [%v]", err.Error())
		return nil, fmt.Errorf("cannot get all collections: %s", err.Error())
	}
	collections := []*pb_collection.Collection{}
	for _, c := range cs {
		collections = append(collections, &pb_collection.Collection{
			Key:      c.Key,
			Name:     c.Name,
			Capacity: c.Capacity,
			Used:     c.Used,
		})
	}

	return &pb_collection.GetAllCollectionsResponse{
		Collections: collections,
	}, nil
}
func (s *Server) GetCollectionByKey(ctx context.Context, in *pb_collection.GetCollectionByKeyRequest) (*pb_collection.Collection, error) {
	c, err := s.db.GetCollectionByKey(ctx, in.GetKey())
	if err != nil {
		log.Error().Err(err).Msgf("GetCollectionByKey:s.db.GetCollectionByKey [%v]", err.Error())
		return nil, fmt.Errorf("collection not found by key %s: %s", in.GetKey(), err.Error())
	}
	return &pb_collection.Collection{
		Key:      c.Key,
		Name:     c.Name,
		Capacity: c.Capacity,
		Used:     c.Used,
	}, nil
}
