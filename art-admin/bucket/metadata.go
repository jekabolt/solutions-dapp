package bucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/minio/minio-go"
)

type Meta interface {
	UploadMetadata(metadata map[int]Metadata) (string, error)
}

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Edition     int    `json:"edition"`
	Date        int64  `json:"date"`
}

// upload _metadata.json to bucket
func (b *Bucket) UploadMetadata(metadata map[int]Metadata) (string, error) {

	meta := []Metadata{}
	for _, m := range metadata {
		meta = append(meta, m)
	}

	metaB, err := json.Marshal(meta)
	if err != nil {
		return "", fmt.Errorf("UploadMetadata:json.Marshal [%v]", err.Error())
	}

	jsonReader := bytes.NewReader(metaB)
	url, err := b.uploadJsonToBucket(jsonReader, contentTypeJSON)
	if err != nil {
		return "", fmt.Errorf("Upload:UploadMetadata: [%v]", err.Error())
	}
	return url, nil
}

func (b *Bucket) uploadJsonToBucket(json io.Reader, contentType string) (string, error) {

	fp := b.getMetadataFullPath()

	userMetaData := map[string]string{"x-amz-acl": "public-read"} // make it public
	cacheControl := "max-age=31536000"

	bs, _ := ioutil.ReadAll(json)

	r := bytes.NewReader(bs)

	_, err := b.Client.PutObject(b.S3BucketName, fp, r, int64(len(bs)), minio.PutObjectOptions{ContentType: contentType, CacheControl: cacheControl, UserMetadata: userMetaData})
	if err != nil {
		return "", fmt.Errorf("PutObject:err [%v]", err.Error())
	}

	return b.GetCDNURL(fp), nil
}
