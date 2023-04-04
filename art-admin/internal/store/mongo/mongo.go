package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mintRequestsCollection = "mintrequests"
	metadataCollection     = "metadata"
	collectionsCollection  = "collections"
)

type Config struct {
	DSN      string `env:"MONGO_DB_DSN" envDefault:"mongodb+srv://sol:password@mongo-sol-35718b1e.mongo.ondigitalocean.com/test?replicaSet=mongo-sol&tls=true&authSource=admin"`
	DB       string `env:"MONGO_DB_NAME" envDefault:"test"`
	PageSize int    `env:"MONGO_DB_PAGE_SIZE" envDefault:"30"`
}

type Store interface {
	MintRequestStore
	MetadataStore
	CollectionsStore
}

type MDB struct {
	*mongo.Client
	mintRequests *mongo.Collection
	metadata     *mongo.Collection
	collections  *mongo.Collection
	pageSize     int32
}

func (mdb *MDB) Close(ctx context.Context) error {
	return mdb.Client.Disconnect(ctx)
}

func (c *Config) InitDB(ctx context.Context) (*MDB, error) {

	opts := options.Client().ApplyURI(c.DSN)
	opts.SetServerSelectionTimeout(time.Second * 3)

	mgo, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("InitDB:mongo.Connect %v", err)
	}

	return &MDB{
		Client:       mgo,
		mintRequests: mgo.Database(c.DB).Collection(mintRequestsCollection),
		metadata:     mgo.Database(c.DB).Collection(metadataCollection),
		collections:  mgo.Database(c.DB).Collection(collectionsCollection),
		pageSize:     int32(c.PageSize),
	}, nil
}
