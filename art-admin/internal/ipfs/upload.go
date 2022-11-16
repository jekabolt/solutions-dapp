package ipfs

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

const (
	uploadFolderPath = "/api/v2/ipfs/uploadFolder"
)

type IPFS interface {
	BulkUploadIPFS(mrs []*pb_nft.NFTMintRequestWithStatus) (map[int]pb_metadata.MetadataUnit, error)
}

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Edition     int    `json:"edition"`
}

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
		Path:           ipfsPath,
		SequenceNumber: sequenceNumber,
	}, nil
}

func mintRequestsToUpload(mrs []*pb_nft.NFTMintRequestWithStatus) ([]byte, error) {
	uf := []UploadFolder{}
	for _, mr := range mrs {
		// TODO: downoad image from from s3
		//
		//!!!!!!!!!!!
		// make option put images to metadata from s3
		ext, err := bucket.GetExtensionFromB64String(mr.OffchainUrl)
		if err != nil {
			return nil, fmt.Errorf("can't get file extension from offchain url")
		}
		uf = append(uf, UploadFolder{
			Path:    fmt.Sprintf("%d.%s", mr.NftMintRequest.MintSequenceNumber, ext),
			Content: mr.OffchainUrl,
		})
	}
	return json.Marshal(uf)
}

func (m *Moralis) BulkUploadIPFS(mrs []*pb_nft.NFTMintRequestWithStatus) (map[int]pb_metadata.MetadataUnit, error) {
	reqBody, err := mintRequestsToUpload(mrs)
	if err != nil {
		return nil, fmt.Errorf("BulkUploadIPFS:mintRequestsToUpload [%v]", err.Error())
	}
	ufs := []UploadFolder{}
	err = m.post(uploadFolderPath, reqBody, &ufs)

	meta := map[int]pb_metadata.MetadataUnit{}

	for _, uf := range ufs {
		ipfsImg, err := uf.GetIPFSImage()
		if err != nil {
			return nil, fmt.Errorf("BulkUploadIPFS:GetIPFSImage [%v]", err.Error())
		}
		meta[ipfsImg.SequenceNumber] = pb_metadata.MetadataUnit{
			Name:        m.desc.GetCollectionName(ipfsImg.SequenceNumber),
			Description: m.desc.GetDescription(ipfsImg.SequenceNumber),
			// TODO: option for offchain
			// Image:       ipfsImg.Path,
			Edition: int32(ipfsImg.SequenceNumber),
		}
	}
	return meta, err
}
