package redis

import (
	"context"
	"fmt"
	"time"

	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	"github.com/rueian/rueidis/om"
)

const (
	allMetadataRequests = "metadata"
)

type Metadata struct {
	Key        string                      `redis:",key" json:"key"`
	Ver        int64                       `redis:",ver"`
	Meta       []*pb_metadata.MetadataUnit `json:"info"`
	IPFSUrl    string                      `json:"ipfs_url"`
	Processing bool                        `json:"processing"`
	Ts         int64                       `json:"ts"`
}

type MetadataStore interface {
	AddOffchainMetadata(ctx context.Context, md []*pb_metadata.MetadataUnit) (string, error)
	SetIPFSUrl(ctx context.Context, key string, IPFSUrl string) error
	SetProcessing(ctx context.Context, key string, processing bool) error
	GetAllOffchainMetadata(ctx context.Context) ([]*Metadata, error)
	GetMetadataByKey(ctx context.Context, key string) (*Metadata, error)
	DeleteById(ctx context.Context, key string) error
}

type metadataStore struct {
	*RDB
}

// MetadataStore returns a metadata store
func (rdb *RDB) MetadataStore(ctx context.Context) (MetadataStore, error) {
	rdb.metadata.DropIndex(ctx)
	err := rdb.metadata.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
		return om.Completed(schema.
			FieldName("ts").Text().Build())
	})
	if err != nil {
		return nil, fmt.Errorf("MetadataStore:CreateIndex [%v]", err.Error())
	}
	return &metadataStore{
		RDB: rdb,
	}, nil
}

// TODO: test
func (rdb *RDB) AddOffchainMetadata(ctx context.Context, md []*pb_metadata.MetadataUnit) (string, error) {
	mdE := rdb.metadata.NewEntity()
	mdE.Meta = md
	mdE.Ts = time.Now().Unix()
	err := rdb.metadata.Save(ctx, mdE)
	if err != nil {
		return "", fmt.Errorf("AddOffchainMetadata:Save [%v]", err.Error())
	}
	return mdE.Key, nil
}

// TODO: test
func (rdb *RDB) SetIPFSUrl(ctx context.Context, key string, IPFSUrl string) error {
	md, err := rdb.metadata.Fetch(ctx, key)
	if err != nil {
		return fmt.Errorf("no such metadata with provided for key %s", key)
	}

	// set ipfs url and set and terminate processing
	md.IPFSUrl = IPFSUrl
	md.Processing = false
	md.Ts = time.Now().Unix()

	err = rdb.metadata.Save(ctx, md)
	if err != nil {
		return fmt.Errorf("SetIPFSUrl:Save [%v]", err.Error())
	}
	return nil
}

// TODO: test
func (rdb *RDB) SetProcessing(ctx context.Context, key string, processing bool) error {
	md, err := rdb.metadata.Fetch(ctx, key)
	if err != nil {
		return fmt.Errorf("no such metadata with provided key %s", key)
	}
	md.Processing = processing
	err = rdb.metadata.Save(ctx, md)
	if err != nil {
		return fmt.Errorf("SetProcessing:Save [%v]", err.Error())
	}
	return nil
}

func (rdb *RDB) GetAllOffchainMetadata(ctx context.Context) ([]*Metadata, error) {
	_, records, err := rdb.metadata.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query("*").Limit().OffsetNum(0, 1000000).Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:Search [%v]", err.Error())
	}
	return records, nil
}

// TODO: test
func (rdb *RDB) GetMetadataByKey(ctx context.Context, key string) (*Metadata, error) {
	md, err := rdb.metadata.Fetch(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("GetOffchainMetadataById:rdb.metadata.Fetch [%v]", err.Error())
	}
	return md, nil
}

// TODO: test
func (rdb *RDB) DeleteById(ctx context.Context, key string) error {
	err := rdb.metadata.Remove(ctx, key)
	if err != nil {
		return fmt.Errorf("DeleteById:rdb.metadata.Fetch [%v]", err.Error())
	}
	return nil
}
