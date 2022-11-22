package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/rueian/rueidis/om"
)

const (
	allCollectionsRequests = "collections"
)

type Collection struct {
	Key      string `redis:",key" json:"key"`
	Ver      int64  `redis:",ver"`
	Name     string `json:"name"`
	Capacity int32  `json:"capacity"`
	Used     int32  `json:"used"`
	Ts       int64  `json:"ts"`
}

type CollectionsStore interface {
	AddCollection(ctx context.Context, name string, capacity int32) (string, error)
	DeleteCollection(ctx context.Context, key string) error
	UpdateCollectionCapacity(ctx context.Context, key string, capacity int32) error
	UpdateCollectionName(ctx context.Context, key string, name string) error
	IncrementUsed(ctx context.Context, key string) error
	GetAllCollections(ctx context.Context) ([]*Collection, error)
	GetCollectionByKey(ctx context.Context, key string) (*Collection, error)
}

type collectionsStore struct {
	*RDB
}

// CollectionStore returns a collections store
func (rdb *RDB) CollectionStore(ctx context.Context) (CollectionsStore, error) {
	rdb.collections.DropIndex(ctx)
	err := rdb.collections.CreateIndex(ctx, func(schema om.FtCreateSchema) om.Completed {
		return om.Completed(schema.
			FieldName("name").Text().Build())
	})
	if err != nil {
		return nil, fmt.Errorf("CollectionStore:CreateIndex [%v]", err.Error())
	}
	return &collectionsStore{
		RDB: rdb,
	}, nil
}

func (rdb *RDB) AddCollection(ctx context.Context, name string, capacity int32) (string, error) {
	cE := rdb.collections.NewEntity()
	cE.Name = name
	cE.Capacity = capacity
	cE.Ts = time.Now().Unix()
	err := rdb.collections.Save(ctx, cE)
	if err != nil {
		return "", fmt.Errorf("AddCollection:Save [%v]", err.Error())
	}
	return cE.Key, nil
}
func (rdb *RDB) DeleteCollection(ctx context.Context, key string) error {
	c, err := rdb.collections.Fetch(ctx, key)
	if err != nil {
		return fmt.Errorf("DeleteCollection:no such collection with provided for key %s", key)
	}

	if c.Used > 0 {
		return fmt.Errorf("cannot delete collection [%s] : is not empty", key)
	}

	err = rdb.collections.Remove(ctx, key)
	if err != nil {
		return fmt.Errorf("DeleteCollection:rdb.collections.Remove [%v]", err.Error())
	}
	return nil
}
func (rdb *RDB) UpdateCollectionCapacity(ctx context.Context, key string, capacity int32) error {
	c, err := rdb.collections.Fetch(ctx, key)
	if err != nil {
		return fmt.Errorf("UpdateCollectionCapacity:no such collection with provided for key %s", key)
	}

	if c.Used > capacity {
		return fmt.Errorf("cannot update collection capacity: min value is %d", c.Used)
	}

	c.Capacity = capacity

	err = rdb.collections.Save(ctx, c)
	if err != nil {
		return fmt.Errorf("UpdateCollectionCapacity:Save [%v]", err.Error())
	}
	return nil
}
func (rdb *RDB) UpdateCollectionName(ctx context.Context, key string, name string) error {
	c, err := rdb.collections.Fetch(ctx, key)
	if err != nil {
		return fmt.Errorf("UpdateCollectionName:no such collection with provided for key %s", key)
	}

	c.Name = name

	err = rdb.collections.Save(ctx, c)
	if err != nil {
		return fmt.Errorf("UpdateCollectionName:Save [%v]", err.Error())
	}
	return nil
}

func (rdb *RDB) IncrementUsed(ctx context.Context, key string) error {
	c, err := rdb.collections.Fetch(ctx, key)
	if err != nil {
		return fmt.Errorf("IncrementUsed:no such collection with provided for key %s", key)
	}

	c.Used++

	if c.Used >= c.Capacity {
		return fmt.Errorf("IncrementUsed:collection is full")
	}

	err = rdb.collections.Save(ctx, c)
	if err != nil {
		return fmt.Errorf("IncrementUsed:Save [%v]", err.Error())
	}
	return nil
}

func (rdb *RDB) GetAllCollections(ctx context.Context) ([]*Collection, error) {
	_, records, err := rdb.collections.Search(ctx, func(search om.FtSearchIndex) om.Completed {
		return search.Query("*").Limit().OffsetNum(0, 1000000).Build()
	})
	if err != nil {
		return nil, fmt.Errorf("GetAllCollections:Search [%v]", err.Error())
	}
	return records, nil
}

func (rdb *RDB) GetCollectionByKey(ctx context.Context, key string) (*Collection, error) {
	md, err := rdb.collections.Fetch(ctx, key)
	if err != nil {
		return nil, fmt.Errorf("GetCollectionByKey:rdb.collections.Fetch [%v]", err.Error())
	}
	return md, nil
}
