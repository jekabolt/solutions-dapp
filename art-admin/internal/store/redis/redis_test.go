package redis

import (
	"context"
	"os"
	"testing"

	"github.com/matryer/is"
)

func getRedisAddress() RedisConf {
	if os.Getenv("REDIS_HOST") == "" {
		return RedisConf{
			Host: "localhost:9999",
		}
	}
	return RedisConf{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}
func TestCreateD(t *testing.T) {
	t.Logf("redis address: %s", getRedisAddress())
	rc := getRedisAddress()
	c := Config{
		Address:  rc.Host,
		Password: rc.Password,
		CacheTTL: "1s",
		PageSize: 30,
	}

	is := is.New(t)
	rdb, err := c.InitDB(context.Background())
	is.NoErr(err)

	rdb.B().Ping()
}
