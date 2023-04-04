package ipfs

import (
	"fmt"
	"strconv"
	"strings"
)

// const (
// 	uploadFolderPath = "/api/v2/ipfs/uploadFolder"
// )

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

// https://ipfs.moralis.io:2053/ipfs/QmPpv79r2crWAA9NXssAobbtz1frJwW9jXFM7dSJJ57E5M/0.jpg
func (ufs *UploadFolder) GetIPFSUrl() (string, error) {
	moralisPathSplit := strings.Split(ufs.Path, "/ipfs/")
	if len(moralisPathSplit) != 2 {
		return "", fmt.Errorf("GetIPFSImage: invalid moralisPathSplit [%s]", ufs.Path)
	}
	ipfsUrlWithPath := moralisPathSplit[1]
	ipfsUrl := strings.Split(ipfsUrlWithPath, "/")[0] // strip 0.jpg
	if len(ipfsUrl) == 0 {
		return "", fmt.Errorf("GetIPFSImage: invalid ipfsPathSplit [%s]", ufs.Path)
	}
	return fmt.Sprintf("ipfs://%s", ipfsUrl), nil
}
