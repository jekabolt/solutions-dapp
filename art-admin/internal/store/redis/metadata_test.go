package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

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

	defer func() {
		for i := 0; i < 100; i++ {
			err := ms.AddOffchainMetadata(ctx, fmt.Sprint(i))
			is.NoErr(err)
		}

		all, err := ms.GetAllOffchainMetadata(ctx)
		is.NoErr(err)
		is.Equal(len(all), 100)

		for i := 0; i < 100; i++ {
			err := ms.DeleteMetadataById(ctx, fmt.Sprint(i))
			is.NoErr(err)
		}
	}()
}
