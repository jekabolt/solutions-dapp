package ipfs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
)

const (
	uploadFolderPath = "/api/v2/ipfs/uploadFolder"
)

type UploadFolder struct {
	Path    string `json:"path"`
	Content string `json:"content,omitempty"`
}

type IpfsImage struct {
	Path           string `json:"path"`
	SequenceNumber int    `json:"sequenceNumber"`
}

// https://ipfs.moralis.io:2053/ipfs/QmPpv79r2crWAA9NXssAobbtz1frJwW9jXFM7dSJJ57E5M/0.jpg
func (ufs *UploadFolder) GetIPFSImage() (*IpfsImage, error) {
	moralisPathSplit := strings.Split(ufs.Path, "/ipfs/")
	if len(moralisPathSplit) != 2 {
		return nil, fmt.Errorf("GetIPFSImage: invalid moralisPathSplit [%s]", ufs.Path)
	}
	ipfsPath := moralisPathSplit[1]
	ipfsPathSplit := strings.Split(ipfsPath, "/")[0]
	if len(ipfsPathSplit) == 0 {
		return nil, fmt.Errorf("GetIPFSImage: invalid ipfsPathSplit [%s]", ufs.Path)
	}
	sequenceNumber, err := strconv.Atoi(ipfsPathSplit[:len(ipfsPathSplit)-3])
	if len(ipfsPathSplit) == 0 {
		return nil, fmt.Errorf("GetIPFSImage: invalid sequenceNumber [%s]", err.Error())
	}

	return &IpfsImage{
		Path:           fmt.Sprintf("/ipfs/%s", ipfsPath),
		SequenceNumber: sequenceNumber,
	}, nil
}

func mintRequestsToUpload(mrs []nft.NFTMintRequest) ([]byte, error) {
	uf := []UploadFolder{}
	for _, mr := range mrs {
		ext, err := bucket.GetExtensionFromB64String(mr.NFTOffchain)
		if err != nil {
			return nil, err
		}
		uf = append(uf, UploadFolder{
			Path:    fmt.Sprintf("%d.%s", mr.MintSequenceNumber, ext),
			Content: mr.NFTOffchain,
		})
	}
	return json.Marshal(uf)
}

func (m *Moralis) BulkUploadIPFS(mrs []nft.NFTMintRequest) (map[int]bucket.Metadata, error) {
	reqBody, err := mintRequestsToUpload(mrs)
	if err != nil {
		return nil, fmt.Errorf("BulkUploadIPFS:mintRequestsToUpload [%v]", err.Error())
	}
	ufs := []UploadFolder{}
	err = m.post(uploadFolderPath, reqBody, &ufs)

	meta := map[int]bucket.Metadata{}

	for _, uf := range ufs {
		ipfsImg, err := uf.GetIPFSImage()
		if err != nil {
			return nil, fmt.Errorf("BulkUploadIPFS:GetIPFSImage [%v]", err.Error())
		}
		meta[ipfsImg.SequenceNumber] = bucket.Metadata{
			Name:        m.desc.GetCollectionName(ipfsImg.SequenceNumber),
			Description: m.desc.GetDescription(ipfsImg.SequenceNumber),
			Image:       ipfsImg.Path,
			Edition:     ipfsImg.SequenceNumber,
			Date:        time.Now().Unix(),
		}
	}
	return meta, err
}
