package bunt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/buntdb"
	"github.com/tidwall/gjson"
)

const (
	firstId    = 1000001
	DBPageSize = 30
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
	db       *buntdb.DB
	pageSize int32
}

func (c *Config) InitDB() (*BuntDB, error) {
	bunt := BuntDB{}

	var err error
	bunt.db, err = buntdb.Open(c.DBPath)
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

	err = bunt.db.CreateIndex(allNFTMintRequests, fmt.Sprintf("%s:*", allNFTMintRequests), buntdb.IndexJSON("id"))
	if err != nil {
		return nil, err
	}

	// new indexes
	err = bunt.db.CreateIndex(statusNFTMintRequests, fmt.Sprintf("%s:*", allNFTMintRequests),
		buntdb.IndexJSON("id"),
		IndexJSONOptional("status"),
	)
	if err != nil {
		return nil, err
	}
	//TODO: to env config
	bunt.pageSize = DBPageSize

	return &bunt, nil
}

func IndexJSONOptional(path string) func(a, b string) bool {
	return func(a, b string) bool {
		opa := gjson.Get(a, path)
		opb := gjson.Get(b, path)

		if !opa.Exists() || !opb.Exists() {
			return false
		}

		return opa.Less(opb, false)
	}
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

func (bunt *BuntDB) GetNextKeyWStatus(index, status string) (int32, error) {
	last := firstId

	err := bunt.db.View(func(tx *buntdb.Tx) error {
		tx.AscendEqual("name", statusJSON(status), func(key, _ string) bool {
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
		return 1, fmt.Errorf("GetNextKeyWStatus:bunt.db.View:err [%v]", err.Error())
	}
	return int32(last), nil
}

func statusJSON(v string) string {
	return fmt.Sprintf("{\"status\":\"%v\"}", v)
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

func (bunt *BuntDB) GetAllJSONStatus(index string, dst interface{}) error {
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

func (db *BuntDB) lastKey(index, status string) (int32, error) {
	next, err := db.GetNextKeyWStatus(index, status)
	if err != nil {
		return 0, err
	}
	if next-firstId == 1 {
		return next, nil
	}
	return next, nil
}

type pagesCount struct {
	lastKey int
	count   int
}

func (db *BuntDB) getPagesCount(index, status string) (int32, error) {
	lastAbsolute, err := db.lastKey(index, status)
	if err != nil {
		return 0, err
	}
	last := lastAbsolute - firstId

	if last < db.pageSize {
		return 1, nil
	}

	pagesF := float64(last) / float64(db.pageSize)
	pagesI := last / db.pageSize

	if pagesF > float64(pagesI) {
		return pagesI + 1, nil
	}
	return pagesI, nil

}

type pgRange struct {
	greaterThan string
	lessOrEqual string
	lastPage    int32
}

func (db *BuntDB) getPagesRange(index, status string, page int32) (*pgRange, error) {
	totalPages, err := db.getPagesCount(index, status)
	if err != nil {
		return nil, err
	}
	pgr := &pgRange{
		lastPage:    totalPages,
		greaterThan: fmt.Sprintf("%s:%d", index, 0),
		lessOrEqual: fmt.Sprintf("%s:%d", index, 0),
	}

	// out of pages range
	if page > totalPages {
		return pgr, nil
	}

	// last page
	if page == totalPages {
		pgr.lessOrEqual = fmt.Sprintf("%s:%d", index, (db.pageSize*page)+firstId)
		pgr.greaterThan = fmt.Sprintf("%s:%d", index, (db.pageSize*(page-1))+firstId)
	}

	return pgr, nil
}
