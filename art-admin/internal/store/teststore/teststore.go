package teststore

import (
	"math/rand"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"

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
	redis.Store
	metadata    map[string]*redis.Metadata
	mintRequest []*pb_nft.NFTMintRequestWithStatus
	pageSize    int
}

func NewTestStore(pageSize int) *testStore {
	return &testStore{
		metadata:    make(map[string]*redis.Metadata),
		mintRequest: []*pb_nft.NFTMintRequestWithStatus{},
		pageSize:    pageSize,
	}
}
