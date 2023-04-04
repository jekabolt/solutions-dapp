package teststore

import (
	"context"
	"fmt"

	pb_collection "github.com/jekabolt/solutions-dapp/art-admin/proto/collection"
)

func (ts *testStore) AddCollection(ctx context.Context, name string, capacity int32) (string, error) {
	id := getId()
	ts.collections[id] = &pb_collection.Collection{
		Key:      id,
		Name:     name,
		Capacity: capacity,
		Used:     0,
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

func (ts *testStore) GetAllCollections(ctx context.Context) ([]*pb_collection.Collection, error) {
	c := []*pb_collection.Collection{}
	for _, m := range ts.collections {
		c = append(c, m)
	}
	return c, nil
}
func (ts *testStore) GetCollectionByKey(ctx context.Context, key string) (*pb_collection.Collection, error) {
	return ts.collections[key], nil
}

func (ts *testStore) IsFull(ctx context.Context, key string) (bool, error) {
	c, ok := ts.collections[key]
	if !ok {
		return false, fmt.Errorf("no such key")
	}
	return c.Used >= c.Capacity, nil
}
