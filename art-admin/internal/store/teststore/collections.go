package teststore

import (
	"context"
	"fmt"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
)

func (ts *testStore) AddCollection(ctx context.Context, name string, capacity int32) (string, error) {
	id := getId()
	ts.collections[id] = &redis.Collection{
		Key:      id,
		Ver:      1,
		Name:     name,
		Capacity: capacity,
		Ts:       time.Now().Unix(),
	}
	return id, nil
}

func (ts *testStore) DeleteCollection(ctx context.Context, key string) error {
	c, ok := ts.collections[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	if c.Used > 0 {
		return fmt.Errorf("cannot delete collection [%s] : is not empty", key)
	}
	delete(ts.collections, key)
	return nil
}
func (ts *testStore) UpdateCollectionCapacity(ctx context.Context, key string, capacity int32) error {
	c, ok := ts.collections[key]
	if !ok {
		return fmt.Errorf("no such key")
	}

	if c.Used > capacity {
		return fmt.Errorf("cannot update collection capacity: min value is %d", c.Used)
	}

	c.Capacity = capacity
	ts.collections[key] = c
	return nil
}
func (ts *testStore) UpdateCollectionName(ctx context.Context, key string, name string) error {
	c, ok := ts.collections[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	c.Name = name
	ts.collections[key] = c
	return nil
}

func (ts *testStore) IncrementUsed(ctx context.Context, key string) error {
	c, ok := ts.collections[key]
	if !ok {
		return fmt.Errorf("no such key")
	}
	col := *c

	if col.Used >= col.Capacity {
		return fmt.Errorf("IncrementUsed:collection is full")
	}
	col.Used++

	ts.collections[key] = &col
	return nil
}

func (ts *testStore) GetAllCollections(ctx context.Context) ([]*redis.Collection, error) {
	c := []*redis.Collection{}
	for _, m := range ts.collections {
		c = append(c, m)
	}
	return c, nil
}
func (ts *testStore) GetCollectionByKey(ctx context.Context, key string) (*redis.Collection, error) {
	return ts.collections[key], nil
}
