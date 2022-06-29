package bunt

import (
	"encoding/json"
	"fmt"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

const (
	allBurns = "burn"
)

type BurnStore interface {
	BurnNft(req *pb_nft.BurnRequest) error
	GetBurned() ([]*pb_nft.BurnShippingInfo, error)
	GetBurnedPending() ([]*pb_nft.BurnShippingInfo, error)
	GetBurnedErrors() ([]*pb_nft.BurnShippingInfo, error)
	UpdateShippingStatus(req *pb_nft.ShippingStatusUpdateRequest) error
}

type burnStore struct {
	*BuntDB
}

// MetadataStore returns a metadata store
func (bdb *BuntDB) BurnStore() BurnStore {
	return &burnStore{
		BuntDB: bdb,
	}
}

func (bunt *BuntDB) BurnNft(req *pb_nft.BurnRequest) error {
	id, err := bunt.GetNextKey(allBurns)
	if err != nil {
		return fmt.Errorf("BurnNft:json.Marshal [%v]", err.Error())
	}
	si := pb_nft.BurnShippingInfo{
		Id:     id,
		Burn:   req,
		Status: &pb_nft.ShippingStatus{},
	}
	bs, err := json.Marshal(si)
	if err != nil {
		return fmt.Errorf("BurnNft:json.Marshal [%v]", err.Error())
	}
	err = bunt.Set(allBurns, fmt.Sprint(id), string(bs))
	if err != nil {
		return fmt.Errorf("BurnNft:SetNext %v", err)
	}
	return nil
}

func (bunt *BuntDB) GetBurned() ([]*pb_nft.BurnShippingInfo, error) {
	bsi := []*pb_nft.BurnShippingInfo{}
	err := bunt.GetAllJSON(allBurns, &bsi)
	if err != nil {
		return nil, fmt.Errorf("GetBurned:GetAllJSON [%v]", err.Error())
	}
	return bsi, err
}

func (bunt *BuntDB) GetBurnedPending() ([]*pb_nft.BurnShippingInfo, error) {
	sis := []*pb_nft.BurnShippingInfo{}
	err := bunt.GetAllJSON(allBurns, &sis)
	if err != nil {
		return nil, fmt.Errorf("GetBurnedPending:GetAllJSON [%v]", err.Error())
	}
	pending := []*pb_nft.BurnShippingInfo{}
	for _, si := range sis {
		if !si.Status.Success && si.Status.Error == "" {
			pending = append(pending, si)
		}
	}
	return pending, err
}

func (bunt *BuntDB) GetBurnedErrors() ([]*pb_nft.BurnShippingInfo, error) {
	sis := []*pb_nft.BurnShippingInfo{}
	err := bunt.GetAllJSON(allBurns, &sis)
	if err != nil {
		return nil, fmt.Errorf("GetBurnedErrors:GetAllJSON [%v]", err.Error())
	}
	errors := []*pb_nft.BurnShippingInfo{}
	for _, si := range sis {
		if si.Status.Error != "" {
			errors = append(errors, si)
		}
	}
	return errors, err
}

func (bunt *BuntDB) UpdateShippingStatus(req *pb_nft.ShippingStatusUpdateRequest) error {
	bsi := pb_nft.BurnShippingInfo{}
	err := bunt.GetJSONById(allBurns, req.Id, &bsi)
	if err != nil {
		return fmt.Errorf("UpdateShippingStatus:GetJSONById [%v]", err.Error())
	}
	bsi.Status = req.Status
	bs, err := json.Marshal(&bsi)
	if err != nil {
		return fmt.Errorf("UpdateShippingStatus:json.Marshal [%v]", err.Error())
	}
	err = bunt.Set(allBurns, req.Id, string(bs))
	if err != nil {
		return fmt.Errorf("UpdateShippingStatus:bunt.Set [%v]", err.Error())
	}
	return nil
}
