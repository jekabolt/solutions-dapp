package mongo

import (
	"context"
	"fmt"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	allNFTMintRequests = "mintrequests"
)

type mintRequestStore struct {
	*MDB
}

type GetToUpload interface {
	// get all mint requests with 'uploaded' status and return map of sequence number to mint request
	GetAllToUpload(ctx context.Context) (map[int32]*pb_nft.NFTMintRequestWithStatus, error)
}

type MintRequestStore interface {
	// CreateNew creates new mint request with status unknown
	New(ctx context.Context, mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error)
	// GetById returns mint request by id
	GetById(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error)
	// UpdateStatus updates status of mint request
	UpdateStatus(ctx context.Context, id string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error)
	// GetAll returns all mint requests with given status
	GetAll(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error)
	// GetPaged returns paged mint requests with given status
	GetPaged(ctx context.Context, listOpts *pb_nft.ListPagedRequest) ([]*pb_nft.NFTMintRequestWithStatus, error)
	// DeleteMintById deletes mint request by id
	DeleteMintById(ctx context.Context, id string) error
	// GetAllToUpload get all mints that ready for to be uploaded to ipfs
	// possible statuses:
	// - StatusUploaded
	// - StatusBurned
	// - StatusShipped
	// used to compose _metadata.json
	GetAllToUpload(ctx context.Context) (map[int32]*pb_nft.NFTMintRequestWithStatus, error)
	// UpdateIpfsUrl updates IPFS url for mint request and sets status to Status_Uploaded
	// can be used only for:
	// - Status_UploadedOffchain
	// - Status_Uploaded
	UpdateIpfsUrl(ctx context.Context, id string, IpfsUrl string) (*pb_nft.NFTMintRequestWithStatus, error)
	// UpdateOffchainUrl updates s3 url for mint request and sets status to Status_UploadedOffchain
	// can be used only for:
	// - Status_Pending
	UpdateOffchainUrl(ctx context.Context, id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error)
	// DeleteIpfsUrl deletes IPFS url for mint request and sets status to
	DeleteIpfsUrl(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error)
	// UpdateShippingInfo updates shipping info for mint request and sets status to Status_Shipped
	UpdateShippingInfo(ctx context.Context, shippingReq *pb_nft.BurnRequest) (*pb_nft.NFTMintRequestWithStatus, error)
	// UpdateTrackingNumber updates tracking number for mint request and sets status to Status_Burned
	UpdateTrackingNumber(ctx context.Context, req *pb_nft.SetTrackingNumberRequest) (*pb_nft.NFTMintRequestWithStatus, error)
}

// MintRequestStore returns a metadata store
func (mdb *MDB) MintRequestStore(ctx context.Context) (MintRequestStore, error) {
	// TODO: add indexes
	return &metadataStore{
		MDB: mdb,
	}, nil
}

type MintRequestWithStatus struct {
	ID             *primitive.ObjectID              `json:"id" bson:"_id,omitempty"`
	MintWithStatus *pb_nft.NFTMintRequestWithStatus `json:"mint" bson:"mint"`
}

// New Create new mint request with status StatusUnknown
func (mdb *MDB) New(ctx context.Context, mr *pb_nft.NFTMintRequestToUpload, il []*pb_nft.ImageList) (*pb_nft.NFTMintRequestWithStatus, error) {
	oid := primitive.NewObjectID()
	_, err := mdb.mintRequests.InsertOne(context.Background(),
		MintRequestWithStatus{
			ID: &oid,
			MintWithStatus: &pb_nft.NFTMintRequestWithStatus{
				OffchainUrl: "",
				OnchainUrl:  "",
				Status:      pb_nft.Status_Unknown,
				NftMintRequest: &pb_nft.NFTMintRequest{
					Id:                 oid.Hex(),
					EthAddress:         mr.EthAddress,
					Description:        mr.Description,
					MintSequenceNumber: 0,
				},
				SampleImages: il,
				Collection:   "",
				Duration:     "",
				Shipping: &pb_nft.Shipping{
					Shipping: &pb_nft.ShippingTo{},
				},
			},
		})
	if err != nil {
		return nil, fmt.Errorf("New:InsertOne [%v]", err.Error())
	}
	return nil, nil
}

// GetById get mint by internal id
func (mdb *MDB) GetById(ctx context.Context, oid string) (*pb_nft.NFTMintRequestWithStatus, error) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("GetById:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	var mr MintRequestWithStatus
	err = mdb.mintRequests.FindOne(ctx, filter).Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("GetById:FindOne [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateStatus update mint status by internal id
func (mdb *MDB) UpdateStatus(ctx context.Context, oid string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("UpdateStatus:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	var mr MintRequestWithStatus
	filter := bson.M{
		"_id": id,
	}
	update := bson.M{"$set": bson.M{"mint.status": status}}
	err = mdb.mintRequests.FindOneAndUpdate(ctx, filter, update).Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateStatus:FindOneAndUpdate [%v]", err.Error())
	}
	mr.MintWithStatus.Status = status
	return mr.MintWithStatus, err
}

// GetAll get all mints
func (mdb *MDB) GetAll(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	res, err := mdb.mintRequests.Find(ctx, getQueryStatus(status))
	if err != nil {
		return nil, fmt.Errorf("GetAllCollections:mdb.mintRequests.Find [%v]", err.Error())
	}
	defer res.Close(ctx)

	var mrs []*pb_nft.NFTMintRequestWithStatus
	for res.Next(ctx) {
		var mr MintRequestWithStatus
		err := res.Decode(&mr)
		if err != nil {
			return nil, fmt.Errorf("GetAllCollections:res.Decode [%v]", err.Error())
		}
		mrs = append(mrs, mr.MintWithStatus)
	}
	return mrs, nil
}

// GetPagedNFTMintRequests get all mints paged which selected status
func (mdb *MDB) GetPaged(ctx context.Context, listOpts *pb_nft.ListPagedRequest) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	if listOpts.Page <= 0 {
		return nil, fmt.Errorf("GetPaged:page must be greater than 0")
	}
	res, err := mdb.mintRequests.Find(ctx, getQueryStatus(listOpts.Status),
		options.Find().
			SetSkip(int64((listOpts.Page-1)*mdb.pageSize)).
			SetLimit(int64(mdb.pageSize)).
			SetSort(bson.M{"_id": 1}),
	)
	if err != nil {
		return nil, fmt.Errorf("GetPaged:mdb.mintRequests.Find [%v]", err.Error())
	}
	defer res.Close(ctx)

	pbList := []*pb_nft.NFTMintRequestWithStatus{}
	for res.Next(ctx) {
		var mr MintRequestWithStatus
		err := res.Decode(&mr)
		if err != nil {
			return nil, fmt.Errorf("GetPaged:res.Decode [%v]", err.Error())
		}
		pbList = append(pbList, mr.MintWithStatus)
	}
	return pbList, nil
}

// DeleteMintById delete mint by internal id
func (mdb *MDB) DeleteMintById(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("DeleteMintById:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": oid}
	_, err = mdb.mintRequests.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("DeleteMintById:mintRequests.DeleteOne [%v]", err.Error())
	}
	return nil
}

// // GetAllToUpload get all mints that ready for to be uploaded to ipfs
// // possible statuses StatusUploaded, StatusBurned, StatusShipped
// // used to compose _metadata.json
func (mdb *MDB) GetAllToUpload(ctx context.Context) (map[int32]*pb_nft.NFTMintRequestWithStatus, error) {
	res, err := mdb.mintRequests.Find(ctx, bson.M{
		"mint.status": bson.M{
			"$in": []pb_nft.Status{
				pb_nft.Status_Uploaded,
				pb_nft.Status_Burned,
				pb_nft.Status_Shipped,
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllToUpload:mdb.mintRequests.Find [%v]", err.Error())
	}
	defer res.Close(ctx)

	mrs := map[int32]*pb_nft.NFTMintRequestWithStatus{}
	for res.Next(ctx) {
		var mr MintRequestWithStatus
		err := res.Decode(&mr)
		if err != nil {
			return nil, fmt.Errorf("GetAllToUpload:res.Decode [%v]", err.Error())
		}
		mrs[mr.MintWithStatus.NftMintRequest.MintSequenceNumber] = mr.MintWithStatus
	}
	return mrs, nil
}

// UpdateIpfsUrl set final art image url uploaded to s3 and update status to StatusUploadedOffchain
func (mdb *MDB) UpdateIpfsUrl(ctx context.Context, id string, onchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error) {
	if onchainUrl == "" {
		return nil, fmt.Errorf("UpdateIpfsUrl:offchainUrl is empty")
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("UpdateIpfsUrl:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	// TODO: test this i.e if status is already uploaded offchain, then it should not update
	res := mdb.mintRequests.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id": oid,
			"$or": []bson.M{
				{"mint.status": pb_nft.Status_UploadedOffchain},
				{"mint.status": pb_nft.Status_Uploaded},
			},
		},
		bson.M{"$set": bson.M{"mint.status": pb_nft.Status_Uploaded, "mint.onchainUrl": onchainUrl}},
	)

	var mr MintRequestWithStatus
	err = res.Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateIpfsUrl:res.Decode [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateOffchainUrl set final art image url uploaded to s3 and update status to StatusUploadedOffchain
func (mdb *MDB) UpdateOffchainUrl(ctx context.Context, id string, offchainUrl string) (*pb_nft.NFTMintRequestWithStatus, error) {
	if offchainUrl == "" {
		return nil, fmt.Errorf("UpdateOffchainUrl:offchainUrl is empty")
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainUrl:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	// TODO: test this i.e if status is already uploaded offchain, then it should not update
	res := mdb.mintRequests.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id": oid,
			"$or": []bson.M{
				{"mint.status": pb_nft.Status_Pending},
			},
		},
		bson.M{"$set": bson.M{"mint.status": pb_nft.Status_UploadedOffchain, "mint.offchainUrl": offchainUrl}},
	)

	var mr MintRequestWithStatus
	err = res.Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainUrl:res.Decode [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// DeleteNFTOffchainUrl delete url for mint and set status to StatusPending
func (mdb *MDB) DeleteIpfsUrl(ctx context.Context, id string) (*pb_nft.NFTMintRequestWithStatus, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, fmt.Errorf("DeleteNFTOffchainUrl:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	res := mdb.mintRequests.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id":         oid,
			"mint.status": pb_nft.Status_Uploaded,
		},
		bson.M{"$set": bson.M{"mint.status": pb_nft.Status_UploadedOffchain, "mint.onchainUrl": ""}},
	)
	var mr MintRequestWithStatus
	err = res.Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainUrl:res.Decode [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateShippingInfo sets shipping info and status StatusBurned
func (mdb *MDB) UpdateShippingInfo(ctx context.Context, req *pb_nft.BurnRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("UpdateShippingInfo:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	res := mdb.mintRequests.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id":         oid,
			"mint.status": pb_nft.Status_Uploaded,
		},
		bson.M{"$set": bson.M{
			"mint.status": pb_nft.Status_Burned,
			"mint.shipping": &pb_nft.Shipping{
				Shipping:       req.Shipping,
				TrackingNumber: "",
			},
		}},
	)
	var mr MintRequestWithStatus
	err = res.Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainUrl:res.Decode [%v]", err.Error())
	}
	return mr.MintWithStatus, err
}

// UpdateTrackingNumber updates shipping tracking number and set status StatusShipped
func (mdb *MDB) UpdateTrackingNumber(ctx context.Context, req *pb_nft.SetTrackingNumberRequest) (*pb_nft.NFTMintRequestWithStatus, error) {
	oid, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, fmt.Errorf("UpdateShippingInfo:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	res := mdb.mintRequests.FindOneAndUpdate(
		ctx,
		bson.M{
			"_id":         oid,
			"mint.status": pb_nft.Status_Burned,
		},
		bson.M{"$set": bson.M{
			"mint.status":                  pb_nft.Status_Shipped,
			"mint.shipping.trackingNumber": req.TrackingNumber,
		}},
	)
	var mr MintRequestWithStatus
	err = res.Decode(&mr)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainUrl:res.Decode [%v]", err.Error())
	}
	mr.MintWithStatus.Status = pb_nft.Status_Shipped
	mr.MintWithStatus.Shipping.TrackingNumber = req.TrackingNumber
	return mr.MintWithStatus, err
}

func getQueryStatus(status pb_nft.Status) bson.M {
	if status == pb_nft.Status_Any {
		return bson.M{}
	}
	return bson.M{"mint.status": status}
}
