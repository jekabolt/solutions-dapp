package ipfs

import (
	"fmt"
	"net/url"
	"time"

	"github.com/valyala/fasthttp"
)

const (
	uploadFolderPath = "/api/v2/ipfs/uploadFolder"
)

type Uploader interface {
	UploadData(data []byte) (string, error)
}

type Config struct {
	APIKey  string `env:"MORALIS_API_KEY"`
	Timeout string `env:"MORALIS_TIMEOUT" envDefault:"10s"`
	BaseURL string `env:"MORALIS_BASE_URL" envDefault:"https://deep-index.moralis.io/api/v2/"`
}

type IPFS struct {
	cli     *fasthttp.Client
	c       *Config
	BaseURL *url.URL
}

func (c *Config) New() (Uploader, error) {
	tOut, err := time.ParseDuration(c.Timeout)
	if err != nil && c.Timeout != "" {
		return nil, fmt.Errorf("init ipfs:time.ParseDuration [%s]", err.Error())
	}
	baseURL, err := url.Parse(c.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("init ipfs :url.Parse %s", err)
	}
	hc := initHTTPClient(tOut)

	return &IPFS{
		c:       c,
		cli:     hc,
		BaseURL: baseURL,
	}, nil
}

func (i IPFS) UploadData(data []byte) (string, error) {
	ufs := []UploadFolder{}
	err := i.post(uploadFolderPath, data, &ufs)
	if err != nil {
		return "", fmt.Errorf("UploadData:m.post [%v]", err.Error())
	}
	if len(ufs) == 0 {
		return "", fmt.Errorf("UploadData:empty response")
	}
	url, err := ufs[0].GetIPFSUrl()
	if err != nil {
		return "", fmt.Errorf("UploadData:ufs[0].GetIPFSImage [%v]", err.Error())
	}
	return url, nil
}

// func (i IPFS) BulkUploadIPFS(mrs []*pb_nft.NFTMintRequestWithStatus) (map[int]pb_metadata.MetadataUnit, error) {

// 	reqBody, err := mintRequestsToUpload(mrs)
// 	if err != nil {
// 		return nil, fmt.Errorf("BulkUploadIPFS:mintRequestsToUpload [%v]", err.Error())
// 	}
// 	ufs := []UploadFolder{}
// 	err = m.post(uploadFolderPath, reqBody, &ufs)

// 	meta := map[int]pb_metadata.MetadataUnit{}

// 	for _, uf := range ufs {
// 		ipfsImg, err := uf.GetIPFSImage()
// 		if err != nil {
// 			return nil, fmt.Errorf("BulkUploadIPFS:GetIPFSImage [%v]", err.Error())
// 		}
// 		meta[ipfsImg.SequenceNumber] = pb_metadata.MetadataUnit{
// 			Name:        m.desc.GetCollectionName(ipfsImg.SequenceNumber),
// 			Description: m.desc.GetDescription(ipfsImg.SequenceNumber),
// 			// TODO: option for offchain
// 			// Image:       ipfsImg.Path,
// 		}
// 	}
// 	return meta, err
// }
