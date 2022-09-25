package redis

type Config struct {
	Address    string `env:"REDIS_DB_ADDRESS" envDefault:"localhost:6379"`
	Password   string `env:"REDIS_DB_PASSWORD" envDefault:""`
	Username   string `env:"REDIS_DB_USERNAME" envDefault:""`
	ClientName string `env:"REDIS_DB_CLIENT_NAME" envDefault:""`
	CacheTTL   string `env:"REDIS_DB_CACHE_TTL" envDefault:"1m"`
	DB         int    `env:"REDIS_DB_NUMBER" envDefault:"0"`
	PoolSize   int    `env:"REDIS_DB_POOL_SIZE" envDefault:"10"`
	PageSize   int32  `env:"REDIS_DB_PAGE_SIZE" envDefault:"30"`
}
