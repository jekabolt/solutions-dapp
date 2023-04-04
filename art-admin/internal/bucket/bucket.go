package bucket

import (
	"github.com/minio/minio-go"
)

type Config struct {
	S3AccessKey       string `env:"S3_ACCESS_KEY" envDefault:"xxx"`
	S3SecretAccessKey string `env:"S3_SECRET_ACCESS_KEY" envDefault:"xxx"`
	S3Endpoint        string `env:"S3_ENDPOINT" envDefault:"fra1.digitaloceanspaces.com"`
	S3BucketName      string `env:"S3_BUCKET_NAME" envDefault:"grbpwr"`
	S3BucketLocation  string `env:"S3_BUCKET_LOCATION" envDefault:"fra-1"`
	BaseFolder        string `env:"S3_BASE_FOLDER" envDefault:"solutions"`
	ImageStorePrefix  string `env:"S3_IMAGE_STORE_PREFIX" envDefault:"grbpwr-com"`
	IPFSStoragePath   string `env:"S3_IPFS_STORAGE_PATH" envDefault:""`
}

type FileStore interface {
	Image
}

type Bucket struct {
	*minio.Client
	*Config
}

type B64Image struct {
	Content     []byte
	ContentType string
}

func (c *Config) Init() (FileStore, error) {
	cli, err := minio.New(c.S3Endpoint, c.S3AccessKey, c.S3SecretAccessKey, true)
	return &Bucket{
		Client: cli,
		Config: c,
	}, err
}
