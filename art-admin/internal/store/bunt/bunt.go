package bunt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/buntdb"
)

const (
	firstId = 1000001
)

type Store interface {
	MintRequestStore
	MetadataStore
	BurnStore
}

type Config struct {
	DBPath string `env:"BUNT_DB_PATH" envDefault:"/tmp/storage.db"`
}

type dbi interface {
	GetNextKey(index string) (int32, error)
	KeyUsed(index string, key int) bool
	SetNext(index, value string) error
	Set(index, key, value string) error
	Delete(index, id string) error
	GetAllJSON(index string, dst interface{}) error
	GetJSONById(index, id string, dst interface{}) error
	GetAll(index string) ([]string, error)
}

type BuntDB struct {
	db *buntdb.DB
}

func (c *Config) InitDB() (*BuntDB, error) {
	bunt := BuntDB{}

	var err error
	bunt.db, err = buntdb.Open(c.DBPath)
	if err != nil {
		return nil, err
	}

	err = bunt.db.CreateIndex(allNFTMintRequests, fmt.Sprintf("%s:*", allNFTMintRequests), buntdb.IndexJSON("id"))
	if err != nil {
		return nil, err
	}
	err = bunt.db.CreateIndex(allBurns, fmt.Sprintf("%s:*", allBurns), buntdb.IndexString)
	if err != nil {
		return nil, err
	}

	err = bunt.db.CreateIndex(allMetadataRequests, fmt.Sprintf("%s:*", allMetadataRequests), buntdb.IndexString)
	if err != nil {
		return nil, err
	}

	return &bunt, nil
}

func (bunt *BuntDB) GetNextKey(index string) (int32, error) {
	last := firstId

	err := bunt.db.View(func(tx *buntdb.Tx) error {
		tx.DescendKeys(fmt.Sprintf("%s:*", index), func(key, _ string) bool {
			id := strings.Trim(key, index+":")
			var err error
			last, err = strconv.Atoi(id)
			last++
			if err != nil {
				last = firstId
			}
			return false
		})
		return nil
	})
	if err != nil {
		return 1, fmt.Errorf("GetNextKey:db.db.View:err [%v]", err.Error())
	}
	return int32(last), nil
}

func (bunt *BuntDB) KeyUsed(index string, key int) bool {
	ok := false
	bunt.db.View(func(tx *buntdb.Tx) error {
		_, err := tx.Get(fmt.Sprintf("%s:%d", index, key))
		if err != buntdb.ErrNotFound {
			ok = true
		}
		return nil
	})
	return ok
}

func (bunt *BuntDB) SetNext(index, value string) error {
	key, err := bunt.GetNextKey(index)
	if err != nil {
		return fmt.Errorf("SetNext:bunt.db.GetNextKey [%v]", err.Error())
	}
	err = bunt.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(fmt.Sprintf("%s:%d", index, key), value, nil)
		return err
	})
	if err != nil {
		return fmt.Errorf("SetNext:bunt.db.Update [%s]", err.Error())
	}
	return nil
}

func (bunt *BuntDB) Set(index, key, value string) error {
	err := bunt.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(fmt.Sprintf("%s:%s", index, key), value, nil)
		return err
	})
	if err != nil {
		return fmt.Errorf("Set:bunt.db.Update [%s]", err.Error())
	}
	return nil
}

func (bunt *BuntDB) Delete(index, id string) error {
	err := bunt.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(fmt.Sprintf("%s:%s", index, id))
		return err
	})
	if err != nil {
		return fmt.Errorf("Delete:bunt.db.Update [%s]", err.Error())
	}
	return nil
}

func (bunt *BuntDB) GetAllJSON(index string, dst interface{}) error {
	jsonStr := "["
	err := bunt.db.View(func(tx *buntdb.Tx) error {
		tx.Ascend(index, func(_, value string) bool {
			jsonStr += value + ","

			return true
		})
		return nil
	})
	if err != nil {
		return fmt.Errorf("GetAllJSON:bunt.db.View [%v]", err.Error())
	}
	if len(jsonStr) != 1 {
		jsonStr = jsonStr[:len(jsonStr)-1]
	}
	jsonStr += "]"
	return json.Unmarshal([]byte(jsonStr), &dst)
}

func (bunt *BuntDB) GetJSONById(index, id string, dst interface{}) error {
	err := bunt.db.View(func(tx *buntdb.Tx) error {
		val, err := tx.Get(fmt.Sprintf("%s:%s", index, id))
		if err != nil {
			return err
		}
		return json.Unmarshal([]byte(val), &dst)
	})
	if err != nil {
		return fmt.Errorf("GetJSONById:bunt.db.View [%v]", err.Error())
	}
	return nil
}

func (bunt *BuntDB) GetAll(index string) ([]string, error) {
	dst := []string{}
	err := bunt.db.View(func(tx *buntdb.Tx) error {
		return tx.Ascend(index, func(k, val string) bool {
			dst = append(dst, val)
			return true
		})
	})
	if err != nil {
		return dst, fmt.Errorf("GetAll:bunt.db.View [%s]", err.Error())
	}
	return dst, nil
}
