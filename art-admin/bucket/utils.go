package bucket

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	contentTypeJPEG = "image/jpeg"
	contentTypePNG  = "image/png"
	contentTypeJSON = "application/json"
)

type FileType struct {
	Extension string
	MIMEType  string
}

func fileExtensionFromContentType(contentType string) string {
	switch contentType {
	case contentTypeJPEG:
		return "jpg"
	case contentTypePNG:
		return "png"
	case contentTypeJSON:
		return "json"
	default:
		ss := strings.Split(contentType, "/")
		if len(ss) > 0 {
			return ss[1]
		}
		return contentType
	}
}

func (b *Bucket) getImageFullPath(filenameExtension, prefix string, pe *PathExtra) string {
	now := time.Now()
	if pe == nil {
		return fmt.Sprintf("%d/%s/%d-%s.%s", now.Year(), now.Month().String(), now.UnixNano(), prefix, filenameExtension)
	}
	if prefix == "" || filenameExtension == "" || pe.EthAddr == "" || pe.TxHash == "" {
		return fmt.Sprintf("%d/%s/%d-%s.%s", now.Year(), now.Month().String(), now.UnixNano(), prefix, filenameExtension)
	}
	return fmt.Sprintf("%s/%s/%s/%s-%s-%s.%s", b.ImageStorePrefix, b.BaseFolder, pe.TxHash, prefix, pe.EthAddr, pe.MintSequence, filenameExtension)
}

func (b *Bucket) getMetadataFullPath() string {
	return fmt.Sprintf("%s/%s/%s/%s/_metadata.json",
		b.BaseFolder,
		b.MetadataStorePrefix,
		time.Now().Format("2006-01-02"),
		time.Now().Format(time.Kitchen))
}

func (b *Bucket) GetCDNURL(path string) string {
	return fmt.Sprintf("https://%s.%s/%s", b.S3BucketName, b.S3Endpoint, path)
}

type rawImage struct {
	B64Image  string `json:"b64Image"`
	MIMEType  string `json:"mimeType"`
	Extension string `json:"Extension"`
}

func GetExtensionFromB64String(b64 string) (string, error) {
	// data:image/jpeg;base64,/9j/2wCEA...

	if strings.HasPrefix(b64, "data:") {
		ss := strings.Split(strings.Trim(b64, "data:"), ";")
		if len(ss) > 0 {
			return fileExtensionFromContentType(ss[0]), nil
		}
	}
	return "", fmt.Errorf("GetExtensionFromB64String: bad b64 string: [%s]", b64)
}

// image URL to base64 string
func getImageB64(url string) (*rawImage, error) {

	// data:image/jpeg;base64

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get: url: [%s] err: [%v]", url, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("url: [%s] statusCode: [%d]", url, resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAl: url: [%s] err: [%v]", url, err.Error())
	}

	mimeType := http.DetectContentType(body)

	var base64Encoding string

	base64Encoding += fmt.Sprintf("data:%s;base64,", mimeType)

	// Append the base64 encoded output
	base64Encoding += base64.StdEncoding.EncodeToString(body)

	return &rawImage{
		B64Image:  base64Encoding,
		MIMEType:  mimeType,
		Extension: fileExtensionFromContentType(mimeType),
	}, nil
}
