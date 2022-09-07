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
	UpdateStatusNFTMintRequest(ctx context.Context, id string, status NFTStatus) (*pb_nft.NFTMintRequestWithStatus, error)
	DeleteNFTMintRequestById(ctx context.Context, id string) error

	GetAllNFTMintRequests(ctx context.Context) ([]*pb_nft.NFTMintRequestWithStatus, error)
	GetNFTMintRequestsPaged(ctx context.Context, status NFTStatus, page int) ([]*pb_nft.NFTMintRequestWithStatusPaged, error)

	UpdateNFTOffchainUrl(ctx context.Context, id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error)
	DeleteNFTOffchainUrl(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error)

	UpdateShippingInfo(ctx context.Context, id string, shipping *pb_nft.ShippingTo) (*pb_nft.NFTMintRequestWithStatus, error)
	UpdateTrackingNumber(ctx context.Context, id, trackingNumber string) (*pb_nft.NFTMintRequestWithStatus, error)

	GetAllToUpload(ctx context.Context) ([]*pb_nft.NFTMintRequestWithStatus, error)
}

type NFTStatus string

func (ns NFTStatus) String() string {
	return string(ns)
}

const (
	StatusUnknown          NFTStatus = "unknown"
	StatusPending          NFTStatus = "pending"
	StatusFailed           NFTStatus = "failed"
	StatusUploadedOffchain NFTStatus = "offchain"
	StatusUploaded         NFTStatus = "uploaded"
	StatusBurned           NFTStatus = "burned"
	StatusShipped          NFTStatus = "shipped"
)

type MintRequestWithStatus struct {
	Key            string                           `redis:",key"`
	Ver            int64                            `redis:",ver"`
	Status         string                           `redis:",status"`
	MintWithStatus *pb_nft.NFTMintRequestWithStatus `redis:",mintWithStatus"`
}

type mintRequestStore struct {
	*RDB
}

// MintRequestStore returns a metadata store
func (rdb *RDB) MintRequestStore(ctx context.Context) (MintRequestStore, error) {
	err := rdb.mintRequests.DropIndex(ctx)
	if err != nil {
		return nil, fmt.Errorf("MintRequestStore:DropIndex [%v]", err.Error())
	}
	err = rdb.mintRequests.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
		return om.Completed(schema.
			FieldName("$.status").As("status").Text().Build())
	})
	if err != nil {
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
	mrNew.MintWithStatus.Status = StatusUnknown.String()
	mrNew.Status = StatusUnknown.String()
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
func (rdb *RDB) UpdateStatusNFTMintRequest(ctx context.Context, id string, status NFTStatus) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UpdateStatusNFTMintRequest:FetchCache [%v]", err.Error())
	}
	mr.MintWithStatus.Status = status.String()
	mr.Status = status.String()

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateStatusNFTMintRequest:Save [%v]", err.Error())
	}

	return mr.MintWithStatus, err
}

// GetAllNFTMintRequests get all mints
func (rdb *RDB) GetAllNFTMintRequests(ctx context.Context) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	// cursor, err := rdb.mintRequests.Aggregate(context.Background(), func(search om.FtAggregateIndex) om.Completed {
	// 	return search.Query("@loc:{1}").
	// 		Groupby(1).Property("@id").Reduce("MIN").Nargs(1).Arg("@count").As("minCount").
	// 		Sortby(2).Property("@minCount").Asc().Build()
	// })

	_, records, err := rdb.mintRequests.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query("*").Limit().OffsetNum(0, 100000).Build()
		// return search.Query("@status:").
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
func (rdb *RDB) GetNFTMintRequestsPaged(ctx context.Context, status NFTStatus, page int) ([]*pb_nft.NFTMintRequestWithStatusPaged, error) {
	cursor, err := rdb.mintRequests.Aggregate(context.Background(), func(search om.FtAggregateIndex) om.Completed {
		return search.Query(fmt.Sprintf("*")).Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllNFTMintRequests:Search [%v]", err.Error())
	}

	fmt.Println("cursor", cursor.Total())
	m, err := cursor.Read(ctx)
	fmt.Println("cursor m err", m, err)
	return nil, nil

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
		if mr.Status == StatusUploaded.String() ||
			mr.Status == StatusUploadedOffchain.String() ||
			mr.Status == StatusBurned.String() ||
			mr.Status == StatusShipped.String() {
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

	if !((mr.Status == StatusPending.String()) ||
		(mr.Status == StatusUploadedOffchain.String()) ||
		(mr.Status == StatusUploaded.String())) {
		return nil, fmt.Errorf("UpdateNFTOffchainUrl:can change url only for pending and uploaded")
	}

	mr.MintWithStatus.NftOffchainUrl = offchainUrl
	mr.MintWithStatus.Status = StatusUploadedOffchain.String()
	mr.Status = StatusUploadedOffchain.String()

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
	mr.MintWithStatus.Status = StatusPending.String()
	mr.Status = StatusPending.String()

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("NewNFTMintRequest:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateShippingInfo sets shipping info and status StatusBurned
func (rdb *RDB) UpdateShippingInfo(ctx context.Context, id string, shipping *pb_nft.ShippingTo) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UpdateShippingInfo:Fetch [%v]", err.Error())
	}
	mr.MintWithStatus.Shipping.Shipping = shipping
	mr.MintWithStatus.Status = StatusBurned.String()
	mr.Status = StatusBurned.String()

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateShippingInfo:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateTrackingNumber updates shipping tracking number and set status StatusShipped
func (rdb *RDB) UpdateTrackingNumber(ctx context.Context, id, trackingNumber string) (*pb_nft.NFTMintRequestWithStatus, error) {
	mr, err := rdb.mintRequests.Fetch(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("UpdateTrackingNumber:Fetch [%v]", err.Error())
	}
	mr.MintWithStatus.Shipping.TrackingNumber = trackingNumber
	mr.MintWithStatus.Status = StatusShipped.String()
	mr.Status = StatusShipped.String()

	err = rdb.mintRequests.Save(ctx, mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateTrackingNumber:Save [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}
