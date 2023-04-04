package mongo

import (
	"context"
	"fmt"
	"testing"

	"github.com/matryer/is"
)

func TestCollections(t *testing.T) {

	is := is.New(t)

	ctx := context.Background()

	c := Config{
		DSN: getMongoDSN(),
		DB:  "test",
	}
	mdb, err := c.InitDB(ctx)
	is.NoErr(err)

	cs, err := mdb.CollectionStore(ctx)
	is.NoErr(err)

	defer func() {
		err = mdb.collections.Drop(ctx)
		is.NoErr(err)
	}()

	ids := []string{}
	defer func() {
		err := mdb.collections.Drop(ctx)
		is.NoErr(err)
	}()
	for i := 1; i <= 100; i++ {
		id, err := cs.AddCollection(ctx, fmt.Sprintf("collection_%d", i), int32(i))
		is.NoErr(err)
		ids = append(ids, id)
	}

	all, err := cs.GetAllCollections(ctx)
	is.NoErr(err)
	is.Equal(len(all), 100)

	err = cs.UpdateCollectionCapacity(ctx, ids[0], 100)
	is.NoErr(err)

	for i := 0; i < 50; i++ {
		err := cs.IncrementUsed(ctx, ids[0])
		is.NoErr(err)
	}

	err = cs.UpdateCollectionCapacity(ctx, ids[0], 40)
	is.True(err != nil)

	err = cs.UpdateCollectionCapacity(ctx, ids[0], 150)
	is.NoErr(err)

	for i := 0; i < 100; i++ {
		err := cs.IncrementUsed(ctx, ids[0])
		is.NoErr(err)
	}

	err = cs.IncrementUsed(ctx, ids[0])
	is.True(err != nil)

	err = cs.UpdateCollectionName(ctx, ids[0], "new_name")
	is.NoErr(err)

	col, err := cs.GetCollectionByKey(ctx, ids[0])
	is.NoErr(err)

	is.Equal(col.Capacity, int32(150))
	is.Equal(col.Name, "new_name")
	is.Equal(col.Used, int32(150))

	err = cs.DeleteCollection(ctx, ids[0])
	is.True(err != nil)

	err = cs.DeleteCollection(ctx, ids[1])
	is.NoErr(err)

	col, err = cs.GetCollectionByKey(ctx, ids[1])
	is.True(err != nil)

	all, err = cs.GetAllCollections(ctx)
	is.NoErr(err)
	is.True(all[2].Capacity == 4)

	full, err := cs.IsFull(ctx, ids[2])
	is.NoErr(err)
	is.True(!full)

	for i := 1; i < 4; i++ {
		err := cs.IncrementUsed(ctx, ids[2])
		is.NoErr(err)
	}
	full, err = cs.IsFull(ctx, ids[2])
	is.NoErr(err)
	is.True(full)

}
