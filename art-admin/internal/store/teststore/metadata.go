package teststore

import (
	"context"
	"fmt"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
)

func (ts *testStore) AddOffchainMetadata(ctx context.Context, md []*pb_metadata.MetadataUnit) (string, error) {
	id := getId()
	ts.metadata[id] = &redis.Metadata{
		Key:  id,
		Ver:  1,
		Meta: md,
		Ts:   time.Now().Unix(),
	}
	return id, nil
}

func (ts *testStore) SetIPFSUrl(ctx context.Context, key string, IPFSUrl string) error {
	md, ok := ts.metadata[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	md.Processing = false
	md.IPFSUrl = IPFSUrl
	ts.metadata[key] = md
	return nil
}

func (ts *testStore) SetProcessing(ctx context.Context, key string, processing bool) error {
	md, ok := ts.metadata[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	md.Processing = processing
	ts.metadata[key] = md
	return nil
}
func (ts *testStore) GetAllOffchainMetadata(ctx context.Context) ([]*redis.Metadata, error) {
	md := []*redis.Metadata{}
	for _, m := range ts.metadata {
		md = append(md, m)
	}
	return md, nil
}
func (ts *testStore) GetMetadataByKey(ctx context.Context, key string) (*redis.Metadata, error) {
	return ts.metadata[key], nil
}
func (ts *testStore) DeleteById(ctx context.Context, key string) error {
	delete(ts.metadata, key)
	return nil
}
