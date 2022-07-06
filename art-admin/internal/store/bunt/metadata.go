package bunt

import (
	"fmt"
)

const (
	allMetadataRequests = "metadata"
)

type MetadataStore interface {
	AddOffchainMetadata(url string) error
	GetAllOffchainMetadata() ([]string, error)
}

type metadataStore struct {
	*BuntDB
}

// MetadataStore returns a metadata store
func (bdb *BuntDB) MetadataStore() MetadataStore {
	return &metadataStore{
		BuntDB: bdb,
	}
}

func (bunt *BuntDB) AddOffchainMetadata(url string) error {
	err := bunt.SetNext(allMetadataRequests, url)
	if err != nil {
		return fmt.Errorf("AddOffchainMetadata:SetNext")
	}
	return nil
}

func (bunt *BuntDB) GetAllOffchainMetadata() ([]string, error) {
	all, err := bunt.GetAll(allMetadataRequests)
	if err != nil {
		return nil, fmt.Errorf("GetAllOffchainMetadata:GetAll")
	}
	return all, nil
}
