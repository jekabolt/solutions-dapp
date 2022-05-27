package bunt

import (
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/store"
	"github.com/tidwall/buntdb"
)

func (bunt *BuntDB) GetAllToUpload() ([]*store.NFTMintRequest, error) {
	mrs, err := bunt.GetAllNFTMintRequests()
	if err != nil {
		return nil, fmt.Errorf("GetAllToUpload:GetAllNFTMintRequests:err [%v]", err.Error())
	}
	var toUpload []*store.NFTMintRequest
	for _, mr := range mrs {
		if mr.Status == store.StatusUploaded || mr.Status == store.StatusUploadedOffchain {
			toUpload = append(toUpload, mr)
		}
	}
	return toUpload, nil
}

func (bunt *BuntDB) UpsertNFT(p *store.NFTMintRequest) (*store.NFTMintRequest, error) {
	// nonexist
	mr, err := bunt.GetNFTMintRequestById(fmt.Sprint(p.Id))
	if err != nil {
		return nil, fmt.Errorf("UpsertNFT:GetNFTMintRequestById:err [%v]", err.Error())
	}

	if p.NFTOffchain == "" {
		return nil, fmt.Errorf("UpsertNFT:GetNFTMintRequestById:NFTOffchain is empty ")
	}

	mr.Status = store.StatusUploadedOffchain
	mr.NFTOffchain = p.NFTOffchain

	return mr, bunt.db.Update(func(tx *buntdb.Tx) error {
		tx.Set(fmt.Sprintf("%s:%d", allNFTMintRequests, mr.Id), mr.String(), nil)
		return nil
	})
}

func (bunt *BuntDB) DeleteNFT(id string) (*store.NFTMintRequest, error) {
	// nonexist
	mr, err := bunt.GetNFTMintRequestById(id)
	if err != nil {
		return nil, fmt.Errorf("UpsertNFT:GetNFTMintRequestById:err [%v]", err.Error())
	}

	mr.Status = store.StatusUnknown
	mr.NFTOffchain = ""

	return mr, bunt.db.Update(func(tx *buntdb.Tx) error {
		tx.Set(fmt.Sprintf("%s:%d", allNFTMintRequests, mr.Id), mr.String(), nil)
		return nil
	})
}
