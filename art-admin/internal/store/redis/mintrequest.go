package redis

import (
	"context"
	"fmt"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/rueian/rueidis/om"
)

const (
	allNFTMintRequests    = "mintrequests"
	statusNFTMintRequests = "mintrequestsstatus"
)

type MintRequestStore interface {
	NewNFTMintRequest(ctx context.Context, mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error)
	UpdateStatusNFTMintRequest(ctx context.Context, id string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error)
	DeleteNFTMintRequestById(ctx context.Context, id string) error

	GetAllNFTMintRequests(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error)
	GetNFTMintRequestsPaged(ctx context.Context, listOpts *pb_nft.ListPagedRequest) ([]*pb_nft.NFTMintRequestWithStatus, error)

	UpdateNFTOffchainUrl(ctx context.Context, id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error)
	DeleteNFTOffchainUrl(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error)

	UpdateShippingInfo(ctx context.Context, shippingReq *pb_nft.BurnRequest) (*pb_nft.NFTMintRequestWithStatus, error)
	UpdateTrackingNumber(ctx context.Context, req *pb_nft.SetTrackingNumberRequest) (*pb_nft.NFTMintRequestWithStatus, error)

	GetAllToUpload(ctx context.Context) ([]*pb_nft.NFTMintRequestWithStatus, error)
}

type MintRequestWithStatus struct {
	Key            string                           `json:"key" redis:",key"`
	Ver            int64                            `json:"ver" redis:",ver"`
	MintWithStatus *pb_nft.NFTMintRequestWithStatus `json:"mintWithStatus"`
}

type mintRequestStore struct {
	*RDB
}

// MintRequestStore returns a metadata store
func (rdb *RDB) MintRequestStore(ctx context.Context) (MintRequestStore, error) {
	err := rdb.mintRequests.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
		return om.Completed(schema.
			FieldName("$.mintWithStatus.status").As("status").Numeric().Build())
	})
	if err != nil && err.Error() != errIndexExists {
		return nil, fmt.Errorf("MintRequestStore:CreateIndex [%v]", err.Error())
	}
	return &mintRequestStore{
		RDB: rdb,
	}, nil
}

// NewNFTMintRequest Create new mint request with status StatusUnknown
func (rdb *RDB) NewNFTMintRequest(ctx context.Context, mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error) {

	mrNew := rdb.mintRequests.NewEntity()
	mrNew.MintWithStatus = &pb_nft.NFTMintRequestWithStatus{}
	mrNew.MintWithStatus.NftOffchainUrl = ""
	mrNew.MintWithStatus.Status = pb_nft.Status_Unknown
	mrNew.MintWithStatus.NftMintRequest = mr.NftMintRequest
	mrNew.MintWithStatus.SampleImages = il
	mrNew.MintWithStatus.NftMintRequest.Id = mrNew.Key

	err := rdb.mintRequests.Save(ctx, mrNew)
	if err != nil {
		return nil, fmt.Errorf("NewNFTMintRequest:Save [%v]", err.Error())
	}

	return mrNew.MintWithStatus, err
}

// GetNFTMintRequestById get mint by internal id
func (rdb *RDB) GetNFTMintRequestById(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestById:FetchCache [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateStatusNFTMintRequest update mint status by internal id
func (rdb *RDB) UpdateStatusNFTMintRequest(ctx context.Context, id string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UpdateStatusNFTMintRequest:FetchCache [%v]", err.Error())
	}
	mr.MintWithStatus.Status = status

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateStatusNFTMintRequest:Save [%v]", err.Error())
	}

	return mr.MintWithStatus, err
}

// GetAllNFTMintRequests get all mints
func (rdb *RDB) GetAllNFTMintRequests(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	_, records, err := rdb.mintRequests.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query(getQueryStatus(status)).Limit().OffsetNum(0, 10000).Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:Search [%v]", err.Error())
	}
	pbList := []*pb_nft.NFTMintRequestWithStatus{}
	for _, v := range records {
		pbList = append(pbList, v.MintWithStatus)
	}
	return pbList, nil

}

// GetPagedNFTMintRequests get all mints paged which selected status
func (rdb *RDB) GetNFTMintRequestsPaged(ctx context.Context, listOpts *pb_nft.ListPagedRequest) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	if listOpts.Page <= 0 {
		listOpts.Page = 1
	}

	offset := (listOpts.Page - 1) * rdb.pageSize
	_, records, err := rdb.mintRequests.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query(getQueryStatus(listOpts.Status)).
			Limit().
			OffsetNum(int64(offset), int64(rdb.pageSize)).
			Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestsPaged:Search [%v]", err.Error())
	}
	pbList := []*pb_nft.NFTMintRequestWithStatus{}
	for _, v := range records {
		pbList = append(pbList, v.MintWithStatus)
	}
	return pbList, nil
}

// DeleteNFTMintRequestById delete mint by internal id
func (rdb *RDB) DeleteNFTMintRequestById(ctx context.Context, id string) error {
	return rdb.mintRequests.Remove(ctx, id)
}

// GetAllToUpload get all mints that ready for to be uploaded to ipfs
// possible statuses StatusUploaded, StatusUploadedOffchain, StatusBurned, StatusShipped
// used to compose _metadata.json
func (rdb *RDB) GetAllToUpload(ctx context.Context) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	_, records, err := rdb.mintRequests.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query("*").Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllToUpload:Search [%v]", err.Error())
	}
	toUpload := []*pb_nft.NFTMintRequestWithStatus{}
	for _, mr := range records {
		if mr.MintWithStatus.Status == pb_nft.Status_Uploaded ||
			mr.MintWithStatus.Status == pb_nft.Status_UploadedOffchain ||
			mr.MintWithStatus.Status == pb_nft.Status_Burned ||
			mr.MintWithStatus.Status == pb_nft.Status_Shipped {
			toUpload = append(toUpload, mr.MintWithStatus)
		}
	}
	return toUpload, nil
}

// UpdateNFTOffchainUrl set final art url to mint request and update status to StatusUploadedOffchain
func (rdb *RDB) UpdateNFTOffchainUrl(ctx context.Context, id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error) {
	if offchainUrl == "" {
		return nil, fmt.Errorf("UpdateNFTOffchainUrl:offchainUrl is empty")
	}
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UpdateNFTOffchainUrl:FetchCache [%v]", err.Error())
	}

	if !((mr.MintWithStatus.Status == pb_nft.Status_Pending) ||
		(mr.MintWithStatus.Status == pb_nft.Status_UploadedOffchain) ||
		(mr.MintWithStatus.Status == pb_nft.Status_Uploaded)) {
		return nil, fmt.Errorf("UpdateNFTOffchainUrl:can change url only for pending and uploaded")
	}

	mr.MintWithStatus.NftOffchainUrl = offchainUrl
	mr.MintWithStatus.Status = pb_nft.Status_UploadedOffchain

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateNFTOffchainUrl:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// DeleteNFTOffchainUrl delete url for mint and set status to StatusPending
func (rdb *RDB) DeleteNFTOffchainUrl(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("GetNFTMintRequestById:FetchCache [%v]", err.Error())
	}
	mr.MintWithStatus.NftOffchainUrl = ""
	mr.MintWithStatus.Status = pb_nft.Status_Pending

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("NewNFTMintRequest:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateShippingInfo sets shipping info and status StatusBurned
func (rdb *RDB) UpdateShippingInfo(ctx context.Context, shippingReq *pb_nft.BurnRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, shippingReq.Id)
	if err != nil {
		return nil, fmt.Errorf("UpdateShippingInfo:Fetch [%v]", err.Error())
	}
	mr.MintWithStatus.Shipping = &pb_nft.Shipping{}
	mr.MintWithStatus.Shipping.Shipping = shippingReq.Shipping
	mr.MintWithStatus.Status = pb_nft.Status_Burned

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateShippingInfo:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateTrackingNumber updates shipping tracking number and set status StatusShipped
func (rdb *RDB) UpdateTrackingNumber(ctx context.Context, req *pb_nft.SetTrackingNumberRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, req.Id)
	if err != nil {
		return nil, fmt.Errorf("UpdateTrackingNumber:Fetch [%v]", err.Error())
	}
	mr.MintWithStatus.Shipping.TrackingNumber = req.TrackingNumber
	mr.MintWithStatus.Status = pb_nft.Status_Shipped

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateTrackingNumber:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

func getQueryStatus(status pb_nft.Status) (query string) {
	query = fmt.Sprintf("@status:[%d %d]", status.Number(), status.Number())
	if status.Number() == pb_nft.Status_Any.Number() {
		query = "*"
	}
	return
}
