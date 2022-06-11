package bunt

import (
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
)

const (
	allNFTMintRequests = "mintrequests"
)

type mintRequestStore struct {
	*BuntDB
}

// NftStore returns a metadata store
func (bdb *BuntDB) NFTStore() nft.Store {
	return &mintRequestStore{
		BuntDB: bdb,
	}
}

func (bunt *BuntDB) UpsertNFTMintRequest(mr *nft.NFTMintRequest) (*nft.NFTMintRequest, error) {
	if mr.Id == 0 {
		var err error
		mr.Id, err = bunt.GetNextKey(allNFTMintRequests)
		if err != nil {
			return nil, fmt.Errorf("UpsertNFTMintRequest:getNextKey [%v]", err.Error())
		}
	}
	mr.NFTOffchain = ""
	mr.Status = nft.StatusUnknown

	err := bunt.Set(allNFTMintRequests, fmt.Sprint(mr.Id), mr.String())
	if err != nil {
		return nil, fmt.Errorf("UpsertNFTMintRequest:Set [%v]", err.Error())
	}
	return mr, nil
}

func (bunt *BuntDB) GetNFTMintRequestById(id string) (*nft.NFTMintRequest, error) {
	prd := nft.NFTMintRequest{}
	err := bunt.GetJSONById(allNFTMintRequests, id, &prd)
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestById:GetJSONById [%v]", err.Error())
	}
	return &prd, err
}

func (bunt *BuntDB) UpdateStatusNFTMintRequest(p *nft.NFTMintRequest, status nft.NFTStatus) (*nft.NFTMintRequest, error) {
	p.Status = status
	return p, bunt.Set(allNFTMintRequests, fmt.Sprint(p.Id), p.String())
}

func (bunt *BuntDB) GetAllNFTMintRequests() ([]nft.NFTMintRequest, error) {
	nftMRs := []nft.NFTMintRequest{}
	err := bunt.GetAllJSON(allNFTMintRequests, &nftMRs)
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:GetAllJSON [%v]", err.Error())
	}

	return nftMRs, err
}

func (bunt *BuntDB) GetAllTest() ([]string, error) {
	ss, err := bunt.GetAll(allNFTMintRequests)
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:GetAllJSON [%v]", err.Error())
	}

	return ss, err
}

func (bunt *BuntDB) DeleteNFTMintRequestById(id string) error {
	err := bunt.Delete(allNFTMintRequests, id)
	if err != nil {
		return fmt.Errorf("DeleteNFTMintRequestById:GetAllJSON [%v]", err.Error())
	}
	return err
}
