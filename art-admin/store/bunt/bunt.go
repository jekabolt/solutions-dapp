package bunt

import (
	"fmt"

	"github.com/tidwall/buntdb"
)

type Config struct {
	DBPath string `env:"BUNT_DB_PATH" envDefault:"/tmp/storage.db"`
}

type BuntDB struct {
	db *buntdb.DB
}

func (c *Config) InitDB() (*BuntDB, error) {
	db := BuntDB{}

	var err error
	db.db, err = buntdb.Open(c.DBPath)
	if err != nil {
		return nil, err
	}

	err = db.db.CreateIndex(allNFTMintRequests, fmt.Sprintf("%s:*", allNFTMintRequests), buntdb.IndexString)
	if err != nil {
		return nil, err
	}

	return &db, nil
}
