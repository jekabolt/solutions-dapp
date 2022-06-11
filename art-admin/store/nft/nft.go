package nft

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
)

type Store interface {
	UpsertNFTMintRequest(p *NFTMintRequest) (*NFTMintRequest, error)
	GetNFTMintRequestById(id string) (*NFTMintRequest, error)
	UpdateStatusNFTMintRequest(p *NFTMintRequest, status NFTStatus) (*NFTMintRequest, error)
	GetAllNFTMintRequests() ([]NFTMintRequest, error)
	DeleteNFTMintRequestById(id string) error

	GetAllToUpload() ([]NFTMintRequest, error)
	UpsertNFT(p *NFTMintRequest) (*NFTMintRequest, error)
	DeleteNFT(id string) (*NFTMintRequest, error)

	GetAllTest() ([]string, error)
}

type NFTStatus string

const (
	StatusUnknown          NFTStatus = "unknown"
	StatusPending          NFTStatus = "pending"
	StatusCompleted        NFTStatus = "completed"
	StatusFailed           NFTStatus = "failed"
	StatusBad              NFTStatus = "bad"
	StatusUploadedOffchain NFTStatus = "uploadedOffchain"
	StatusUploaded         NFTStatus = "uploaded"
)

func (ns NFTStatus) IsValid() error {
	switch ns {
	case StatusPending, StatusCompleted, StatusFailed, StatusUploaded:
		return nil
	}
	return errors.New("invalid NFTStatus type")
}

type NFTMetaOffchain struct {
	Id                 int    `json:"id"`
	MintSequenceNumber int    `json:"mintSequenceNumber"`
	Image              string `json:"imageOffchain"`
}

type NFTMeta struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageIpfs   string `json:"image"`
	Image       string `json:"imageOffchain"`
	Edition     int    `json:"edition"`
	Date        int64  `json:"date"`
}

type NFTMintRequest struct {
	Id                 int            `json:"id"`
	ETHAddress         string         `json:"ethAddress,omitempty"`
	TxHash             string         `json:"txHash,omitempty"`
	MintSequenceNumber int            `json:"mintSequenceNumber,omitempty"`
	SampleImages       []bucket.Image `json:"sampleImages,omitempty"`
	Description        string         `json:"description,omitempty"`
	Status             NFTStatus      `json:"status,omitempty"`
	NFTOffchain        string         `json:"nftOffchain,omitempty"`
}

func (p *NFTMeta) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}

func (p *NFTMetaOffchain) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}

func (p *NFTMintRequest) String() string {
	bs, _ := json.Marshal(p)
	return string(bs)
}

func (mr *NFTMintRequest) Validate() error {

	if mr.NFTOffchain == "" {
		if len(mr.ETHAddress) == 0 {
			return fmt.Errorf("missing ETHAddress")
		}
		if len(mr.TxHash) == 0 {
			return fmt.Errorf("missing txHash")
		}
		if mr.MintSequenceNumber == 0 {
			return fmt.Errorf("missing MintSequenceNumber")
		}
		// if len(mr.SampleImages) == 0 {
		// 	return fmt.Errorf("missing SampleImages")
		// }
		if len(mr.Description) == 0 {
			return fmt.Errorf("missing  Description")
		}
	}

	if mr.NFTOffchain != "" {
		if mr.Id == 0 {
			return fmt.Errorf("missing id")
		}
	}

	return nil
}

func (m *NFTMeta) Validate() error {
	if len(m.Name) == 0 {
		return fmt.Errorf("missing Name")
	}
	if len(m.Description) == 0 {
		return fmt.Errorf("missing Description")
	}
	if len(m.ImageIpfs) == 0 {
		return fmt.Errorf("missing ImageIpfs")
	}
	if len(m.Image) == 0 {
		return fmt.Errorf("missing Image")
	}
	if m.Date == 0 {
		return fmt.Errorf("missing Date")
	}
	return nil
}
