package redis

import (
	"context"
	"testing"

	"github.com/matryer/is"
)

func TestCreateD(t *testing.T) {
	c := Config{
		Address: "localhost:6379",
	}

	is := is.New(t)
	rdb, err := c.InitDB(context.Background())
	is.NoErr(err)

	rdb.B().Ping()
}
