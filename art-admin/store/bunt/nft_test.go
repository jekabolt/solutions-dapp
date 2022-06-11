package bunt

import (
	"fmt"
	"testing"

	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
	"github.com/matryer/is"
)

func getTestNFTMintRequest() *nft.NFTMintRequest {
	return &nft.NFTMintRequest{
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
		Status:      nft.StatusUnknown,
	}
}

func TestNFT(t *testing.T) {
	is := is.New(t)

	c := Config{
		DBPath: ":memory:",
	}

	bdb, err := c.InitDB()
	is.NoErr(err)

	ms := bdb.NFTStore()

	nftMR := getTestNFTMintRequest()

	_, err = ms.UpsertNFTMintRequest(nftMR)
	is.NoErr(err)

	nftMRs, err := ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)

	nftMR = &nftMRs[0]
	nftMR.Status = nft.StatusPending

	_, err = ms.UpsertNFTMintRequest(nftMR)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, nftMR.Status)

	nftMR.NFTOffchain = "offchain url"
	_, err = ms.UpsertNFT(nftMR)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, nft.StatusUploadedOffchain)
	is.Equal(nftMRs[0].NFTOffchain, nftMR.NFTOffchain)

	_, err = ms.DeleteNFT(fmt.Sprint(nftMR.Id))
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, nft.StatusUnknown)
	is.Equal(nftMRs[0].NFTOffchain, "")

	err = ms.DeleteNFTMintRequestById(fmt.Sprint(nftMR.Id))
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(len(nftMRs), 0)

	n := 100

	for i := 0; i < n; i++ {
		nftMR.Id = 0
		nftMR.ETHAddress = fmt.Sprintf("0x%d", i)
		_, err := ms.UpsertNFTMintRequest(nftMR)
		is.NoErr(err)
	}

}
