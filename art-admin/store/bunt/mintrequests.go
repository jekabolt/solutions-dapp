package bunt

import (
	"encoding/json"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/store"
	"github.com/tidwall/buntdb"
)

const (
	allNFTMintRequests = "mintrequests"
)

func (bunt *BuntDB) UpsertNFTMintRequest(p *store.NFTMintRequest) (*store.NFTMintRequest, error) {
	// new
	if p.Id == 0 && !bunt.keyUsed(allNFTMintRequests, p.Id) {
		var err error
		p.Id, err = bunt.getNextKey(allNFTMintRequests)
		if err != nil {
			return nil, fmt.Errorf("UpsertNFTMintRequest:getNextKey [%v]", err.Error())
		}
	}
	p.NFTOffchain = ""
	p.Status = store.StatusUnknown
	return p, bunt.db.Update(func(tx *buntdb.Tx) error {
		tx.Set(fmt.Sprintf("%s:%d", allNFTMintRequests, p.Id), p.String(), nil)
		return nil
	})
}

func (db *BuntDB) GetNFTMintRequestById(id string) (*store.NFTMintRequest, error) {
	prd := &store.NFTMintRequest{}
	err := db.db.View(func(tx *buntdb.Tx) error {
		NFTMintRequestStr, err := tx.Get(fmt.Sprintf("%s:%s", allNFTMintRequests, id))
		if err != nil {
			return err
		}
		return json.Unmarshal([]byte(NFTMintRequestStr), prd)
	})
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestById:db.db.View:err [%v]", err.Error())
	}
	return prd, err
}

func (bunt *BuntDB) UpdateStatusNFTMintRequest(p *store.NFTMintRequest, status store.NFTStatus) (*store.NFTMintRequest, error) {
	p.Status = status
	return p, bunt.db.Update(func(tx *buntdb.Tx) error {
		tx.Set(fmt.Sprintf("%s:%d", allNFTMintRequests, p.Id), p.String(), nil)
		return nil
	})
}

func (bunt *BuntDB) GetAllNFTMintRequests() ([]*store.NFTMintRequest, error) {
	nftMRsS := "["
	nftMRs := &[]*store.NFTMintRequest{}
	err := bunt.db.View(func(tx *buntdb.Tx) error {
		tx.Ascend(allNFTMintRequests, func(_, nftMRStr string) bool {
			nftMRsS += nftMRStr + ","
			return true
		})
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:db.db.View:err [%v]", err.Error())
	}
	if len(nftMRsS) != 1 {
		nftMRsS = nftMRsS[:len(nftMRsS)-1]
	}
	nftMRsS += "]"
	return *nftMRs, json.Unmarshal([]byte(nftMRsS), nftMRs)
}

func (bunt *BuntDB) DeleteNFTMintRequestById(id string) error {
	err := bunt.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(fmt.Sprintf("%s:%s", allNFTMintRequests, id))
		return err
	})
	return err
}
