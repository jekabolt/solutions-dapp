package bunt

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/store"
	"github.com/matryer/is"
	"github.com/tidwall/buntdb"
)

func TestCreateD(t *testing.T) {
	p := &store.NFTMintRequest{}
	bs, _ := json.Marshal(p)
	fmt.Println("---", string(bs))
}

func buntFromConst() (*BuntDB, error) {
	c := &Config{
		DBPath: ":memory:",
	}
	return c.InitDB()
}
func (b *BuntDB) addTestObj(is *is.I, index string) {
	i, err := b.getNextKey(index)
	is.NoErr(err)
	fmt.Println(i)

	b.db.Update(func(tx *buntdb.Tx) error {
		is.NoErr(err)
		tx.Set(fmt.Sprintf("%s:%d", index, i), "test", nil)
		return nil
	})
}

func TestGetLastId(t *testing.T) {
	is := is.New(t)
	b, err := buntFromConst()
	is.NoErr(err)

	index := "test"

	keyN, err := b.getNextKey(index)
	is.NoErr(err)
	is.Equal(keyN, firstId)

	mTest := map[string]string{}
	for i := keyN; i < keyN+100; i++ {
		b.addTestObj(is, index)
		mTest[fmt.Sprintf("%s:%d", index, i)] = index
	}

	m := map[string]string{}
	b.db.View(func(tx *buntdb.Tx) error {
		tx.Ascend("", func(key, val string) bool {
			m[key] = val
			return true
		})
		return nil
	})
	is.Equal(m, mTest)

}

func getTestNFTMintRequest() *store.NFTMintRequest {
	return &store.NFTMintRequest{
		Id:         0,
		ETHAddress: "0x0",
		TxHash:     "0x0",
		SampleImages: []bucket.Image{
			{
				FullSize: "https://ProductImages.com/img.jpg",
			},
			{
				FullSize: "https://ProductImages2.com/img.jpg",
			},
		},
		Description: "test",
		Status:      store.StatusUnknown,
	}
}

func TestCRUDNFTMintRequest(t *testing.T) {
	is := is.New(t)
	b, err := buntFromConst()
	is.NoErr(err)

	nftMR := getTestNFTMintRequest()
	// new
	_, err = b.UpsertNFTMintRequest(nftMR)
	is.NoErr(err)

	nftMRs, err := b.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)

	// for _, p := range nftMRs {
	// 	fmt.Printf("--kekek %+v-- \n", p.Status)
	// }

	nftMR = nftMRs[0]
	nftMR.Status = store.StatusPending

	_, err = b.UpsertNFTMintRequest(nftMR)
	is.NoErr(err)

	nftMRs, err = b.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, nftMR.Status)

	nftMR.NFTOffchain = "offchain url"
	_, err = b.UpsertNFT(nftMR)
	is.NoErr(err)

	nftMRs, err = b.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, store.StatusUploadedOffchain)
	is.Equal(nftMRs[0].NFTOffchain, nftMR.NFTOffchain)

	_, err = b.DeleteNFT(fmt.Sprint(nftMR.Id))
	is.NoErr(err)

	nftMRs, err = b.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, store.StatusUnknown)
	is.Equal(nftMRs[0].NFTOffchain, "")

	err = b.DeleteNFTMintRequestById(fmt.Sprint(nftMR.Id))
	is.NoErr(err)

	nftMRs, err = b.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 0)

	n := 100

	for i := 0; i < n; i++ {
		nftMR.Id = 0
		nftMR.ETHAddress = fmt.Sprintf("0x%d", i)
		_, err := b.UpsertNFTMintRequest(nftMR)
		is.NoErr(err)
	}

	// add another type of keys
	index := "test"
	for i := 0; i < n; i++ {
		b.addTestObj(is, index)
	}

	nftMRs, err = b.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), n)

	// for _, p := range nftMRs {
	// 	fmt.Println(p.Id)
	// }

}
