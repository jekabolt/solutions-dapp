package bucket

import (
	"github.com/minio/minio-go"
)

type S3BucketConfig struct {
	S3AccessKey         string `env:"S3_ACCESS_KEY" envDefault:"xxx"`
	S3SecretAccessKey   string `env:"S3_SECRET_ACCESS_KEY" envDefault:"xxx"`
	S3Endpoint          string `env:"S3_ENDPOINT" envDefault:"fra1.digitaloceanspaces.com"`
	S3BucketName        string `env:"S3_BUCKET_NAME" envDefault:"grbpwr"`
	S3BucketLocation    string `env:"S3_BUCKET_LOCATION" envDefault:"fra-1"`
	BaseFolder          string `env:"S3_BASE_FOLDER" envDefault:"solutions"`
	ImageStorePrefix    string `env:"IMAGE_STORE_PREFIX" envDefault:"grbpwr-com"`
	MetadataStorePrefix string `env:"METADATA_STORE_PREFIX" envDefault:"metadata"`
	IPFSStoragePath     string `env:"IPFS_STORAGE_PATH" envDefault:""`
}

type Bucket struct {
	*minio.Client
	*S3BucketConfig
}

type B64Image struct {
	Content     []byte
	ContentType string
}

func InitBucket(bc *S3BucketConfig) (*Bucket, error) {
	cli, err := minio.New(bc.S3Endpoint, bc.S3AccessKey, bc.S3SecretAccessKey, true)
	return &Bucket{
		Client:         cli,
		S3BucketConfig: bc,
	}, err
}
