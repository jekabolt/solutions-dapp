package bunt

import (
	"encoding/json"
	"fmt"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
)

const (
	allNFTMintRequests = "mintrequests"
)

type MintRequestStore interface {
	UpsertNFTMintRequest(mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error)
	GetNFTMintRequestById(id string) (*pb_nft.NFTMintRequestWithStatus, error)
	UpdateStatusNFTMintRequest(id string, status NFTStatus) (*pb_nft.NFTMintRequestWithStatus, error)
	GetAllNFTMintRequests() ([]*pb_nft.NFTMintRequestWithStatus, error)
	DeleteNFTMintRequestById(id string) error

	GetAllToUpload() ([]*pb_nft.NFTMintRequestWithStatus, error)
	UpdateNFTOffchainUrl(id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error)
	DeleteNFTOffchainUrl(id string) (*pb_nft.NFTMintRequestWithStatus, error)
}

type NFTStatus string

func (ns NFTStatus) String() string {
	return string(ns)
}

const (
	StatusUnknown          NFTStatus = "unknown"
	StatusPending          NFTStatus = "pending"
	StatusCompleted        NFTStatus = "completed"
	StatusFailed           NFTStatus = "failed"
	StatusBad              NFTStatus = "bad"
	StatusUploadedOffchain NFTStatus = "uploadedOffchain"
	StatusUploaded         NFTStatus = "uploaded"
)

type mintRequestStore struct {
	*BuntDB
}

// MintRequestStore returns a metadata store
func (bdb *BuntDB) MintRequestStore() MintRequestStore {
	return &mintRequestStore{
		BuntDB: bdb,
	}
}

func (bunt *BuntDB) UpsertNFTMintRequest(mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error) {
	mrUpsert := &pb_nft.NFTMintRequestWithStatus{
		NftMintRequest: mr.NftMintRequest,
	}
	if mr.NftMintRequest.Id == 0 {
		var err error
		mrUpsert.NftMintRequest.Id, err = bunt.GetNextKey(allNFTMintRequests)
		if err != nil {
			return nil, fmt.Errorf("UpsertNFTMintRequest:getNextKey [%v]", err.Error())
		}
	}
	mrUpsert.NftOffchainUrl = ""
	mrUpsert.Status = StatusUnknown.String()
	mrUpsert.NftMintRequest = mr.NftMintRequest
	mrUpsert.SampleImages = il

	bs, err := json.Marshal(mrUpsert)
	if err != nil {
		return nil, fmt.Errorf("UpsertNFTMintRequest:protojson.Marshal [%v]", err.Error())
	}

	err = bunt.Set(allNFTMintRequests, fmt.Sprint(mrUpsert.NftMintRequest.Id), string(bs))
	if err != nil {
		return nil, fmt.Errorf("UpsertNFTMintRequest:Set [%v]", err.Error())
	}
	return mrUpsert, nil
}

func (bunt *BuntDB) GetNFTMintRequestById(id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr := pb_nft.NFTMintRequestWithStatus{}
	err := bunt.GetJSONById(allNFTMintRequests, id, &mr)
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestById:GetJSONById [%v]", err.Error())
	}
	return &mr, err
}

func (bunt *BuntDB) UpdateStatusNFTMintRequest(id string, status NFTStatus) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr := pb_nft.NFTMintRequestWithStatus{}
	err := bunt.GetJSONById(allNFTMintRequests, id, &mr)
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestById:GetJSONById [%v]", err.Error())
	}
	mr.Status = status.String()
	bs, err := json.Marshal(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpsertNFTMintRequest:protojson.Marshal [%v]", err.Error())
	}
	return &mr, bunt.Set(allNFTMintRequests, fmt.Sprint(id), string(bs))
}

func (bunt *BuntDB) GetAllNFTMintRequests() ([]*pb_nft.NFTMintRequestWithStatus, error) {
	nftMRs := []*pb_nft.NFTMintRequestWithStatus{}
	err := bunt.GetAllJSON(allNFTMintRequests, &nftMRs)
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:GetAllJSON [%v]", err.Error())
	}
	return nftMRs, err
}

func (bunt *BuntDB) DeleteNFTMintRequestById(id string) error {
	err := bunt.Delete(allNFTMintRequests, id)
	if err != nil {
		return fmt.Errorf("DeleteNFTMintRequestById:GetAllJSON [%v]", err.Error())
	}
	return err
}

func (bunt *BuntDB) GetAllToUpload() ([]*pb_nft.NFTMintRequestWithStatus, error) {
	mrs, err := bunt.GetAllNFTMintRequests()
	if err != nil {
		return nil, fmt.Errorf("GetAllToUpload:GetAllNFTMintRequests:err [%v]", err.Error())
	}
	var toUpload []*pb_nft.NFTMintRequestWithStatus
	for _, mr := range mrs {
		if mr.Status == StatusUploaded.String() ||
			mr.Status == StatusUploadedOffchain.String() {
			toUpload = append(toUpload, mr)
		}
	}
	return toUpload, nil
}

func (bunt *BuntDB) UpdateNFTOffchainUrl(id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error) {
	// not exist
	mr, err := bunt.GetNFTMintRequestById(id)
	if err != nil {
		return nil, fmt.Errorf("UpdateNFTOffchain:GetNFTMintRequestById:err [%v]", err.Error())
	}

	if offchainUrl == "" {
		return nil, fmt.Errorf("UpdateNFTOffchain:GetNFTMintRequestById:NFTOffchain is empty ")
	}

	mr.Status = StatusUploadedOffchain.String()
	mr.NftOffchainUrl = offchainUrl

	bs, err := json.Marshal(mr)
	if err != nil {
		return nil, fmt.Errorf("UpsertNFTMintRequest:protojson.Marshal [%v]", err.Error())
	}

	err = bunt.Set(allNFTMintRequests, fmt.Sprint(id), string(bs))
	if err != nil {
		return nil, fmt.Errorf("UpdateNFTOffchain:Set [%v]", err.Error())
	}
	return mr, err
}

func (bunt *BuntDB) DeleteNFTOffchainUrl(id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	// not exist
	mr, err := bunt.GetNFTMintRequestById(id)
	if err != nil {
		return nil, fmt.Errorf("DeleteNFOffchain:GetNFTMintRequestById:err [%v]", err.Error())
	}

	mr.Status = StatusUnknown.String()
	mr.NftOffchainUrl = ""

	bs, err := json.Marshal(mr)
	if err != nil {
		return nil, fmt.Errorf("UpsertNFTMintRequest:protojson.Marshal [%v]", err.Error())
	}
	err = bunt.Set(allNFTMintRequests, fmt.Sprint(id), string(bs))

	if err != nil {
		return nil, fmt.Errorf("DeleteNFOffchain:Set [%v]", err.Error())
	}
	return mr, err
}
