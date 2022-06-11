package bunt

import (
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
)

func (bunt *BuntDB) GetAllToUpload() ([]nft.NFTMintRequest, error) {
	mrs, err := bunt.GetAllNFTMintRequests()
	if err != nil {
		return nil, fmt.Errorf("GetAllToUpload:GetAllNFTMintRequests:err [%v]", err.Error())
	}
	var toUpload []nft.NFTMintRequest
	for _, mr := range mrs {
		if mr.Status == nft.StatusUploaded || mr.Status == nft.StatusUploadedOffchain {
			toUpload = append(toUpload, mr)
		}
	}
	return toUpload, nil
}

func (bunt *BuntDB) UpsertNFT(p *nft.NFTMintRequest) (*nft.NFTMintRequest, error) {
	// not exist
	mr, err := bunt.GetNFTMintRequestById(fmt.Sprint(p.Id))
	if err != nil {
		return nil, fmt.Errorf("UpsertNFT:GetNFTMintRequestById:err [%v]", err.Error())
	}

	if p.NFTOffchain == "" {
		return nil, fmt.Errorf("UpsertNFT:GetNFTMintRequestById:NFTOffchain is empty ")
	}

	mr.Status = nft.StatusUploadedOffchain
	mr.NFTOffchain = p.NFTOffchain

	err = bunt.Set(allNFTMintRequests, fmt.Sprint(mr.Id), mr.String())
	if err != nil {
		return nil, fmt.Errorf("UpsertNFT:Set [%v]", err.Error())
	}
	return mr, err
}

func (bunt *BuntDB) DeleteNFT(id string) (*nft.NFTMintRequest, error) {
	// not exist
	mr, err := bunt.GetNFTMintRequestById(id)
	if err != nil {
		return nil, fmt.Errorf("UpsertNFT:GetNFTMintRequestById:err [%v]", err.Error())
	}

	mr.Status = nft.StatusUnknown
	mr.NFTOffchain = ""

	err = bunt.Set(allNFTMintRequests, fmt.Sprint(mr.Id), mr.String())

	if err != nil {
		return nil, fmt.Errorf("UpsertNFT:Set [%v]", err.Error())
	}
	return mr, err
}
