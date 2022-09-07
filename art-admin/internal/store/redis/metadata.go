package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/rueian/rueidis/om"
)

const (
	allMetadataRequests = "metadata"
)

type Metadata struct {
	Key string `redis:",key"`
	Ver int64  `redis:",ver"`
	Ts  int64  `redis:",ts"`
	Url string `redis:",url"`
}

type MetadataStore interface {
	AddOffchainMetadata(ctx context.Context, url string) error
	GetAllOffchainMetadata(ctx context.Context) ([]*Metadata, error)
}

type metadataStore struct {
	*RDB
}

// MetadataStore returns a metadata store
func (rdb *RDB) MetadataStore(ctx context.Context) (MetadataStore, error) {
	fmt.Println("--- rdb.metadata.IndexName() ", rdb.metadata.IndexName())
	rdb.metadata.DropIndex(ctx)
	err := rdb.metadata.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
		return om.Completed(schema.
			FieldName("$.ts").As("ts").Text().Build())
	})
	if err != nil {
		return nil, fmt.Errorf("MetadataStore:CreateIndex [%v]", err.Error())
	}
	return &metadataStore{
		RDB: rdb,
	}, nil
}

func (rdb *RDB) AddOffchainMetadata(ctx context.Context, url string) error {
	md := rdb.metadata.NewEntity()
	md.Url = url
	md.Ts = time.Now().Unix()

	err := rdb.metadata.Save(ctx, md)
	if err != nil {
		return fmt.Errorf("AddOffchainMetadata:Save [%v]", err.Error())
	}
	return nil
}

func (rdb *RDB) GetAllOffchainMetadata(ctx context.Context) ([]*Metadata, error) {
	_, records, err := rdb.metadata.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query("*").Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:Search [%v]", err.Error())
	}
	return records, nil
}
