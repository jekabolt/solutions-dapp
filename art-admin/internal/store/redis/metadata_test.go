package redis

import (
	"context"
	"testing"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	"github.com/matryer/is"
)

func getMetadataUnits() []*metadata.MetadataUnit {
	return []*metadata.MetadataUnit{
		{
			Name:               "MetadataUnit 1 ",
			Description:        "description 1 ",
			OffchainImage:      "offchain-image",
			OnchainImage:       "onchain-image",
			Edition:            1,
			MintSequenceNumber: 1,
			Date:               int32(time.Now().Unix()),
		},
		{
			Name:               "MetadataUnit 1 ",
			Description:        "description 1 ",
			OffchainImage:      "offchain-image",
			OnchainImage:       "onchain-image",
			Edition:            1,
			MintSequenceNumber: 1,
			Date:               int32(time.Now().Unix()),
		},
	}
}

func TestMetadata(t *testing.T) {
	is := is.New(t)

	rc := getRedisAddress()
	c := Config{
		Address:  rc.Host,
		Password: rc.Password,
		CacheTTL: "1s",
		PageSize: 30,
	}
	ctx := context.Background()

	bdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := bdb.MetadataStore(ctx)
	is.NoErr(err)

	ids := []string{}
	defer func() {
		for _, id := range ids {
			err := ms.DeleteById(ctx, id)
			is.NoErr(err)
		}
	}()

	for i := 0; i < 100; i++ {
		id, err := ms.AddOffchainMetadata(ctx, getMetadataUnits())
		is.NoErr(err)
		ids = append(ids, id)
	}

	all, err := ms.GetAllOffchainMetadata(ctx)
	is.NoErr(err)
	is.Equal(len(all), 100)

	err = ms.SetProcessing(ctx, ids[0], true)
	is.NoErr(err)

	md, err := ms.GetMetadataByKey(ctx, ids[0])
	is.NoErr(err)
	is.Equal(md.Processing, true)

	err = ms.SetIPFSUrl(ctx, ids[0], "ipfs-url")
	is.NoErr(err)

	md, err = ms.GetMetadataByKey(ctx, ids[0])
	is.NoErr(err)
	is.Equal(md.IPFSUrl, "ipfs-url")

}
