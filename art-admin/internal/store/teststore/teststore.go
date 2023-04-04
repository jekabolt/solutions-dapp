package teststore

import (
	"math/rand"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/mongo"

	pb_collection "github.com/jekabolt/solutions-dapp/art-admin/proto/collection"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getId() string {
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

type testStore struct {
	mongo.Store

	metadata         map[string]*pb_metadata.Meta
	offchainMetadata *pb_metadata.Meta

	collections map[string]*pb_collection.Collection

	mintRequest []*pb_nft.NFTMintRequestWithStatus

	pageSize int
}

func NewTestStore(pageSize int) *testStore {
	return &testStore{
		metadata:         make(map[string]*pb_metadata.Meta),
		collections:      make(map[string]*pb_collection.Collection),
		offchainMetadata: &pb_metadata.Meta{},
		mintRequest:      []*pb_nft.NFTMintRequestWithStatus{},
		pageSize:         pageSize,
	}
}
