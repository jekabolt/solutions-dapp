package mongo

import (
	"context"
	"fmt"
	"time"

	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	allMetadataRequests = "metadata"
)

type MetadataWithId struct {
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
	Meta     *pb_metadata.Meta   `json:"meta" bson:"meta"`
	Offchain bool                `json:"offchain" bson:"offchain"`
}

type metadataStore struct {
	*MDB
}

type MetadataStore interface {
	// GetOffchainMetadata returns initial metadata with all offchain images
	GetOffchainMetadata(ctx context.Context) (*pb_metadata.Meta, error)
	// UpdateOffchainMetadataAttributes updates attributes of offchain metadata for corresponding mint sequence number
	UpdateOffchainMetadataAttributes(ctx context.Context, mintSeq int32, attr []*pb_metadata.Attributes) (*pb_metadata.Meta, error)
	// UpdateOffchainMetadataImage updates image of offchain metadata for corresponding mint sequence number
	UpdateOffchainMetadataImage(ctx context.Context, mintSeq int32, imageUrl string) (*pb_metadata.Meta, error)
	// AddMetadata adds offchain metadata to the store returns id
	AddMetadata(ctx context.Context, md []*pb_metadata.MetadataUnit) (*pb_metadata.MetaInfo, error)
	// SetIPFSUrl sets IPFS url for metadata
	SetIPFSUrl(ctx context.Context, id string, IPFSUrl string) error
	// SetProcessing sets processing flag for metadata need while uploading to IPFS
	SetProcessing(ctx context.Context, id string, processing bool) error
	// SetOffchain sets offchain flag for metadata
	SetOffchain(ctx context.Context, id string) error
	// GetAllMetadata returns all metadata including offchain and onchain
	GetAllMetadata(ctx context.Context) ([]*pb_metadata.Meta, error)
	// GetMetadataById returns metadata by id
	GetMetadataById(ctx context.Context, id string) (*pb_metadata.Meta, error)
	// DeleteById deletes metadata by id
	DeleteById(ctx context.Context, id string) error
}

// MetadataStore returns a metadata store
func (mdb *MDB) MetadataStore() (MetadataStore, error) {
	// TODO: add indexes
	return &metadataStore{
		MDB: mdb,
	}, nil
}

func (mdb *MDB) GetOffchainMetadata(ctx context.Context) (*pb_metadata.Meta, error) {
	var meta MetadataWithId
	err := mdb.metadata.FindOne(context.Background(), bson.M{"offchain": true}).Decode(&meta)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, fmt.Errorf("GetOffchainMetadata:FindOne [%v]", err.Error())
	}
	return meta.Meta, nil
}

func (mdb *MDB) UpdateOffchainMetadataImage(ctx context.Context, mintSeq int32, imageUrl string) (*pb_metadata.Meta, error) {
	var meta MetadataWithId
	err := mdb.metadata.FindOne(context.Background(), bson.M{"offchain": true}).Decode(&meta)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainMetadataImage:FindOne [%v]", err.Error())
	}
	for _, md := range meta.Meta.Metadata {
		if md.MintSequenceNumber == mintSeq {
			md.Image = imageUrl
			break
		}
	}
	filter := bson.M{"offchain": true}
	update := bson.M{"$set": bson.M{"meta": meta.Meta}}
	_, err = mdb.metadata.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainMetadataImage:UpdateOne [%v]", err.Error())
	}
	return meta.Meta, nil
}

func (mdb *MDB) UpdateOffchainMetadataAttributes(ctx context.Context, mintSeq int32, attr []*pb_metadata.Attributes) (*pb_metadata.Meta, error) {
	var meta MetadataWithId
	err := mdb.metadata.FindOne(context.Background(), bson.M{"offchain": true}).Decode(&meta)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainMetadataAttributes:FindOne [%v]", err.Error())
	}
	for _, md := range meta.Meta.Metadata {
		if md.MintSequenceNumber == mintSeq {
			md.Attributes = attr
			break
		}
	}

	filter := bson.M{"offchain": true}
	update := bson.M{"$set": bson.M{"meta": meta.Meta}}
	_, err = mdb.metadata.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, fmt.Errorf("UpdateOffchainMetadataAttributes:UpdateOne [%v]", err.Error())
	}
	return meta.Meta, nil

}

func (mdb *MDB) AddMetadata(ctx context.Context, md []*pb_metadata.MetadataUnit) (*pb_metadata.MetaInfo, error) {
	oid := primitive.NewObjectID()
	meta := MetadataWithId{
		ID: &oid,
		Meta: &pb_metadata.Meta{
			Metadata: md,
			MetaInfo: &pb_metadata.MetaInfo{
				Ts:         time.Now().Unix(),
				IpfsUrl:    "",
				Processing: true,
				Id:         oid.Hex(),
			},
		},
	}
	_, err := mdb.metadata.InsertOne(context.Background(), meta)
	if err != nil {
		return nil, fmt.Errorf("AddOffchainMetadata:InsertOne [%v]", err.Error())
	}
	return meta.Meta.MetaInfo, nil
}

// TODO: test
func (mdb *MDB) SetIPFSUrl(ctx context.Context, oid string, IPFSUrl string) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("SetIPFSUrl:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"meta.metaInfo.ipfsUrl": IPFSUrl}}
	res := mdb.metadata.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return fmt.Errorf("SetIPFSUrl:FindOneAndUpdate [%v]", res.Err().Error())
	}
	return nil
}

func (mdb *MDB) SetProcessing(ctx context.Context, oid string, processing bool) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("SetProcessing:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"meta.metaInfo.processing": processing}}
	res := mdb.metadata.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return fmt.Errorf("SetProcessing:FindOneAndUpdate [%v]", res.Err().Error())
	}
	return nil
}

// SetOffchain sets offchain flag for metadata
func (mdb *MDB) SetOffchain(ctx context.Context, oid string) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("SetOffchain:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	// TODO: check if metadata offchain is not already set

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"offchain": true}}
	res := mdb.metadata.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return fmt.Errorf("SetOffchain:FindOneAndUpdate [%v]", res.Err().Error())
	}
	return nil
}

func (mdb *MDB) GetAllMetadata(ctx context.Context) ([]*pb_metadata.Meta, error) {

	res, err := mdb.metadata.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllMetadata:mdb.metadata.Find [%v]", err.Error())
	}

	metas := []*pb_metadata.Meta{}
	for res.Next(ctx) {
		var meta MetadataWithId
		err := res.Decode(&meta)
		if err != nil {
			return nil, fmt.Errorf("GetAllMetadata:res.Decode [%v]", err.Error())
		}
		metas = append(metas, meta.Meta)
	}

	return metas, nil
}

// TODO: test
func (mdb *MDB) GetMetadataById(ctx context.Context, oid string) (*pb_metadata.Meta, error) {

	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("GetMetadataById:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	var meta MetadataWithId
	err = mdb.metadata.FindOne(ctx, filter).Decode(&meta)
	if err != nil {
		return nil, fmt.Errorf("GetMetadataById:FindOne [%v]", err.Error())
	}
	return meta.Meta, nil
}

// TODO: test
func (mdb *MDB) DeleteById(ctx context.Context, oid string) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("DeleteById:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	res := mdb.metadata.FindOneAndDelete(ctx, filter)
	if res.Err() != nil {
		return fmt.Errorf("DeleteById:FindOneAndDelete [%v]", res.Err().Error())
	}
	return nil
}
