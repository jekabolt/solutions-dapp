package bunt

import (
	"fmt"
	"testing"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
	"github.com/tidwall/buntdb"
)

func getTestNFTMintRequest() (*pb_nft.NFTMintRequestToUpload, []*pb_nft.ImageList) {
	return &pb_nft.NFTMintRequestToUpload{
			SampleImages: []*pb_nft.ImageToUpload{
				{
					Raw: "https://ProductImages.com/img.jpg",
				},
				{
					Raw: "https://ProductImages2.com/img.jpg",
				},
			},
			NftMintRequest: &pb_nft.NFTMintRequest{
				Id:                 0,
				EthAddress:         "0x0",
				TxHash:             "0x0",
				MintSequenceNumber: 3,
				Description:        "test",
			},
		}, []*pb_nft.ImageList{
			{
				FullSize: "https://ProductImages.com/img.jpg",
			},
			{
				FullSize: "https://ProductImages2.com/img.jpg",
			},
		}
}

func TestNFT(t *testing.T) {
	is := is.New(t)

	c := Config{
		DBPath: ":memory:",
	}

	bdb, err := c.InitDB()
	is.NoErr(err)

	ms := bdb.MintRequestStore()

	nftMR, images := getTestNFTMintRequest()

	_, err = ms.UpsertNFTMintRequest(nftMR, images)
	is.NoErr(err)

	nftMRs, err := ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, StatusUnknown.String())

	_, err = ms.UpdateStatusNFTMintRequest(fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		StatusPending)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, StatusPending.String())

	url := "offchain url"
	upd, err := ms.UpdateNFTOffchainUrl(fmt.Sprint(nftMRs[0].NftMintRequest.Id), url)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, upd.Status)
	is.Equal(nftMRs[0].NftOffchainUrl, url)

	_, err = ms.DeleteNFTOffchainUrl(fmt.Sprint(nftMRs[0].NftMintRequest.GetId()))
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, StatusUnknown.String())
	is.Equal(nftMRs[0].NftOffchainUrl, "")

	err = ms.DeleteNFTMintRequestById(fmt.Sprint(nftMRs[0].NftMintRequest.GetId()))
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 0)

	// n := 100

	// for i := 0; i < n; i++ {
	// 	nftMR.Id = 0
	// 	nftMR.ETHAddress = fmt.Sprintf("0x%d", i)
	// 	_, err := ms.UpsertNFTMintRequest(nftMR)
	// 	is.NoErr(err)
	// }

}

func TestNFTPages(t *testing.T) {
	is := is.New(t)

	c := Config{
		DBPath: ":memory:",
	}

	bdb, err := c.InitDB()
	is.NoErr(err)

	ms := bdb.MintRequestStore()

	for i := 0; i < 100; i++ {
		nftMR, images := getTestNFTMintRequest()
		fmt.Println(bdb.GetNextKey(allNFTMintRequests))
		ms.UpsertNFTMintRequest(nftMR, images)
	}

	mrs, err := ms.GetAllNFTMintRequests()
	is.NoErr(err)

	fmt.Println("len: ", len(mrs))
	ok := true
	for _, mr := range mrs {
		if ok {
			_, err := ms.UpdateStatusNFTMintRequest(fmt.Sprint(mr.NftMintRequest.Id), StatusPending)
			is.NoErr(err)
			ok = false
			continue
		}
		ok = true
	}

	pgRange, err := bdb.getPagesRange(statusNFTMintRequests, "unknown", 1)
	is.NoErr(err)

	fmt.Println("pgRange: ", pgRange)

	mrJsonSS := make([]string, 0)
	err = bdb.db.View(func(tx *buntdb.Tx) error {
		tx.AscendEqual(statusNFTMintRequests, statusJSON("unknown"), func(_, mrJsonStr string) bool {
			if len(mrJsonSS) < pgRange.From {

			}
			mrJsonSS = append(mrJsonSS, mrJsonStr)
			return true
		})
		return err
	})

	fmt.Println("---- ", len(mrJsonSS))

}
