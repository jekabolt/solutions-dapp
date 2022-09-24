package redis

import (
	"context"
	"os"
	"testing"

	"github.com/matryer/is"
)

func getRedisAddress() string {
	if os.Getenv("REDIS_HOST") == "" {
		return "localhost:6379"
		// return "redis.sys.solutions:6379"
	}
	return os.Getenv("REDIS_HOST")
}

func TestCreateD(t *testing.T) {
	t.Logf("redis address: %s", getRedisAddress())
	c := Config{
		Address:  getRedisAddress(),
		CacheTTL: "1s",
	}

	is := is.New(t)
	rdb, err := c.InitDB(context.Background())
	is.NoErr(err)

	rdb.B().Ping()
}
