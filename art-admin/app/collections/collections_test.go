package collections

import (
	"context"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/teststore"
	pb_collection "github.com/jekabolt/solutions-dapp/art-admin/proto/collection"

	"github.com/matryer/is"
)

func TestCollections(t *testing.T) {
	is := is.New(t)

	cfg := Config{}

	db := teststore.NewTestStore(30)

	s := cfg.New(db)
	ctx := context.Background()

	cr, err := s.CreateNewCollection(ctx, &pb_collection.CreateNewCollectionRequest{
		Name:     "test crud",
		Capacity: 100,
	})
	is.NoErr(err)

	cByKey, err := s.GetCollectionByKey(ctx, &pb_collection.GetCollectionByKeyRequest{
		Key: cr.Key,
	})
	is.NoErr(err)

	all, err := s.GetAllCollections(ctx, nil)
	is.NoErr(err)

	is.Equal(len(all.Collections), 1)
	is.Equal(all.Collections[0].Key, cByKey.Key)

	_, err = s.UpdateCollectionCapacity(ctx, &pb_collection.UpdateCollectionCapacityRequest{
		Key:      cByKey.Key,
		Capacity: 200,
	})

	_, err = s.UpdateCollectionName(ctx, &pb_collection.UpdateCollectionNameRequest{
		Key:  cByKey.Key,
		Name: "test crud updated",
	})

	cByKey, err = s.GetCollectionByKey(ctx, &pb_collection.GetCollectionByKeyRequest{
		Key: cByKey.Key,
	})

	is.Equal(cByKey.Name, "test crud updated")
	is.Equal(cByKey.Capacity, int32(200))

	_, err = s.DeleteCollection(ctx, &pb_collection.DeleteCollectionRequest{
		Key: cr.Key,
	})
	is.NoErr(err)

	///

	///

	cr2, err := s.CreateNewCollection(ctx, &pb_collection.CreateNewCollectionRequest{
		Name:     "test errors",
		Capacity: 100,
	})
	is.NoErr(err)

	err = s.db.IncrementUsed(ctx, cr2.Key)
	is.NoErr(err)

	_, err = s.DeleteCollection(ctx, &pb_collection.DeleteCollectionRequest{
		Key: cr2.Key,
	})
	is.True(err != nil) // should be error because collection is not empty

	for i := 0; i < 49; i++ { // 49 + 1 = 50
		err := s.db.IncrementUsed(ctx, cr2.Key)
		is.NoErr(err)
	}

	_, err = s.UpdateCollectionCapacity(ctx, &pb_collection.UpdateCollectionCapacityRequest{
		Key:      cr2.Key,
		Capacity: 40,
	})
	is.True(err != nil) // should be error because used > new capacity

}
