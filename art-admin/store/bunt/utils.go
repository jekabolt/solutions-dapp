package bunt

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tidwall/buntdb"
)

const (
	firstId = 1000001
)

func (db *BuntDB) getNextKey(index string) (int, error) {
	last := firstId
	err := db.db.View(func(tx *buntdb.Tx) error {
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
		return 1, fmt.Errorf("getLastId:db.db.View:err [%v]", err.Error())
	}
	return last, nil
}

func (db *BuntDB) keyUsed(index string, key int) bool {
	ok := false
	db.db.View(func(tx *buntdb.Tx) error {
		_, err := tx.Get(fmt.Sprintf("%s:%d", index, key))
		if err != buntdb.ErrNotFound {
			ok = true
		}
		return nil
	})
	return ok
}
