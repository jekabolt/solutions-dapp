package bunt

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
	"github.com/matryer/is"
)

func TestCreateD(t *testing.T) {
	p := &nft.NFTMintRequest{}
	bs, _ := json.Marshal(p)
	fmt.Println("---", string(bs))
}

func TestBuntDBI(t *testing.T) {
	is := is.New(t)

	c := Config{
		DBPath: ":memory:",
	}

	bdb, err := c.InitDB()
	is.NoErr(err)

	ok := bdb.KeyUsed(allMetadataRequests, firstId)
	is.True(!ok)

	testVal := "testVal-%s"
	for i := 0; i < 100; i++ {
		err = bdb.SetNext(allMetadataRequests, fmt.Sprintf(testVal, strconv.Itoa(i)))
		is.NoErr(err)
	}
	all, err := bdb.GetAll(allMetadataRequests)
	is.NoErr(err)
	is.Equal(len(all), 100)

	for i := 0; i < 100; i++ {
		ok = bdb.KeyUsed(allMetadataRequests, firstId+i)
		is.True(ok) // key should be used
	}

	for i := 0; i < 100; i++ {
		err = bdb.Delete(allMetadataRequests, strconv.Itoa(i+firstId))
		is.NoErr(err)
	}

	all, err = bdb.GetAll(allMetadataRequests)
	is.NoErr(err)
	is.Equal(len(all), 0)

	for i := 0; i < 100; i++ {
		err = bdb.SetNext(allMetadataRequests, fmt.Sprintf(`{"id": %d}`, i))
	}
	type test struct {
		Id int `json:"id"`
	}

	all, err = bdb.GetAll(allMetadataRequests)
	is.NoErr(err)
	is.Equal(len(all), 100)

	for i := 0; i < 100; i++ {
		ts := test{}
		err = bdb.GetJSONById(allMetadataRequests, strconv.Itoa(i+firstId), &ts)
		is.NoErr(err)
		is.Equal(ts.Id, i)
	}

	ts := []test{}
	bdb.GetAllJSON(allMetadataRequests, &ts)
	is.Equal(len(all), 100)

	for i := 0; i < 100; i++ {
		err = bdb.Delete(allMetadataRequests, strconv.Itoa(i+firstId))
		is.NoErr(err)
	}

	id := firstId + 1
	err = bdb.Set(allMetadataRequests, fmt.Sprint(id), "test")
	is.NoErr(err)

	ok = bdb.KeyUsed(allMetadataRequests, id)
	is.True(ok)

	ok = bdb.KeyUsed(allMetadataRequests, id+1)
	is.True(!ok)

}
