package redis

import (
	"context"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestMetadata(t *testing.T) {
	is := is.New(t)

	c := Config{
		Address:  "localhost:6379",
		CacheTTL: "1s",
	}
	ctx := context.Background()

	bdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := bdb.MetadataStore(ctx)
	is.NoErr(err)

	for i := 0; i < 100; i++ {
		err := ms.AddOffchainMetadata(ctx, fmt.Sprint(i))
		is.NoErr(err)
	}

	all, err := ms.GetAllOffchainMetadata(ctx)
	is.NoErr(err)
	is.Equal(len(all), 100)

}
