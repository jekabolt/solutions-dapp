package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/rueian/rueidis"
	"github.com/rueian/rueidis/om"
)

type Store interface {
	MintRequestStore
	MetadataStore
}

type RDB struct {
	rueidis.Client
	ttl          time.Duration
	mintRequests om.Repository[MintRequestWithStatus]
	metadata     om.Repository[Metadata]
	pageSize     int
}

func (c *Config) InitDB(ctx context.Context) (rdb *RDB, err error) {
	rdb = &RDB{}
	rdb.ttl, err = time.ParseDuration(c.CacheTTL)
	if err != nil {
		return nil, fmt.Errorf("InitDB:time.ParseDuration %v", err)
	}
	rdb.pageSize = c.PageSize
	rdb.Client, err = rueidis.NewClient(rueidis.ClientOption{
		InitAddress: []string{c.Address},
		Username:    c.Username,
		Password:    c.Password,
		ClientName:  c.ClientName,
		SelectDB:    c.DB,
	})
	if err != nil {
		return nil, fmt.Errorf("InitDB:rueidis.NewClient %v", err)
	}

	rdb.mintRequests = om.NewJSONRepository(allNFTMintRequests, MintRequestWithStatus{}, rdb.Client)
	rdb.metadata = om.NewHashRepository(allMetadataRequests, Metadata{}, rdb.Client)

	return
}
