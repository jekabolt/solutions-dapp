package mongo

import (
	"context"
	"fmt"

	pb_collection "github.com/jekabolt/solutions-dapp/art-admin/proto/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	allCollections = "collections"
)

type CollectionWithId struct {
	ID         *primitive.ObjectID       `json:"ID" bson:"_id,omitempty"`
	Collection *pb_collection.Collection `json:"collection" bson:"collection"`
}

type CollectionsStore interface {
	// AddCollection adds new collection to the store returns key
	AddCollection(ctx context.Context, name string, capacity int32) (string, error)
	// DeleteCollection deletes collection by key
	DeleteCollection(ctx context.Context, key string) error
	// UpdateCollectionCapacity updates capacity of collection by key
	UpdateCollectionCapacity(ctx context.Context, key string, capacity int32) error
	// UpdateCollectionName updates name of collection by key
	UpdateCollectionName(ctx context.Context, key string, name string) error
	// IncrementUsed increments used field of collection by key
	IncrementUsed(ctx context.Context, key string) error
	// DecrementUsed decrements used field of collection by key
	GetAllCollections(ctx context.Context) ([]*pb_collection.Collection, error)
	// GetCollectionByKey returns collection by key
	GetCollectionByKey(ctx context.Context, key string) (*pb_collection.Collection, error)
	// IsFull returns true if collection is full
	IsFull(ctx context.Context, key string) (bool, error)
}

type collectionsStore struct {
	*MDB
}

// CollectionStore returns a collections store
func (mdb *MDB) CollectionStore(ctx context.Context) (CollectionsStore, error) {
	// TODO: add indexes
	return &collectionsStore{
		MDB: mdb,
	}, nil
}

func (mdb *MDB) AddCollection(ctx context.Context, name string, capacity int32) (string, error) {
	oid := primitive.NewObjectID()
	c := &CollectionWithId{
		ID: &oid,
		Collection: &pb_collection.Collection{
			Key:      oid.Hex(),
			Name:     name,
			Capacity: capacity,
			Used:     0,
		},
	}
	_, err := mdb.collections.InsertOne(ctx, c)
	if err != nil {
		return "", fmt.Errorf("AddCollection:InsertOne [%v]", err.Error())
	}
	return oid.Hex(), nil
}
func (mdb *MDB) DeleteCollection(ctx context.Context, oid string) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("DeleteCollection:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	filter := bson.M{"_id": id, "collection.used": 0}
	dr, err := mdb.collections.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("DeleteCollection:DeleteOne [%v]", err.Error())
	}
	if dr.DeletedCount == 0 {
		return fmt.Errorf("DeleteCollection:DeleteOne [%v]", "used != 0")
	}

	return nil
}
func (mdb *MDB) UpdateCollectionCapacity(ctx context.Context, oid string, capacity int32) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("UpdateCollectionCapacity:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	c := &CollectionWithId{}
	filter := bson.M{
		"_id":             id,
		"collection.used": bson.M{"$lte": capacity},
	}
	update := bson.M{"$set": bson.M{"collection.capacity": capacity}}
	err = mdb.collections.FindOneAndUpdate(ctx, filter, update).Decode(&c)
	if err != nil {
		return fmt.Errorf("UpdateCollectionCapacity:FindOneAndUpdate [%v]", err.Error())
	}
	return nil
}
func (mdb *MDB) UpdateCollectionName(ctx context.Context, oid string, name string) error {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("UpdateCollectionName:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"collection.name": name}}
	res := mdb.collections.FindOneAndUpdate(ctx, filter, update)
	if res.Err() != nil {
		return fmt.Errorf("UpdateCollectionName:FindOneAndUpdate [%v]", res.Err().Error())
	}
	return nil
}

func (mdb *MDB) IncrementUsed(ctx context.Context, oid string) error {
	c, err := mdb.GetCollectionByKey(ctx, oid)
	if err != nil {
		return fmt.Errorf("IncrementUsed:GetCollectionByKey [%v]", err.Error())
	}
	if c.Used >= c.Capacity {
		return fmt.Errorf("IncrementUsed:used [%v] >= capacity [%v]", c.Used, c.Capacity)
	}

	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return fmt.Errorf("IncrementUsed:primitive.ObjectIDFromHex [%v]", err.Error())
	}

	update := bson.M{
		"$inc": bson.M{"collection.used": 1},
	}
	_, err = mdb.collections.UpdateByID(ctx, id, update)
	if err != nil {
		return fmt.Errorf("IncrementUsed:FindOneAndUpdate [%v]", err.Error())
	}

	return nil
}

func (mdb *MDB) GetAllCollections(ctx context.Context) ([]*pb_collection.Collection, error) {
	res, err := mdb.collections.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("GetAllCollections:mdb.collections.Find [%v]", err.Error())
	}

	cs := []*pb_collection.Collection{}
	for res.Next(ctx) {
		var c CollectionWithId
		err := res.Decode(&c)
		if err != nil {
			return nil, fmt.Errorf("GetAllCollections:res.Decode [%v]", err.Error())
		}
		cs = append(cs, c.Collection)
	}

	return cs, nil
}

func (mdb *MDB) GetCollectionByKey(ctx context.Context, oid string) (*pb_collection.Collection, error) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionByKey:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	var c CollectionWithId
	err = mdb.collections.FindOne(ctx, filter).Decode(&c)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionByKey:FindOne [%v]", err.Error())
	}
	return c.Collection, nil
}

func (mdb *MDB) IsFull(ctx context.Context, oid string) (bool, error) {
	id, err := primitive.ObjectIDFromHex(oid)
	if err != nil {
		return true, fmt.Errorf("IsFull:primitive.ObjectIDFromHex [%v]", err.Error())
	}
	filter := bson.M{"_id": id}
	var c CollectionWithId
	err = mdb.collections.FindOne(ctx, filter).Decode(&c)
	if err != nil {
		return true, fmt.Errorf("IsFull:FindOne [%v]", err.Error())
	}
	return c.Collection.Used >= c.Collection.Capacity, nil
}
