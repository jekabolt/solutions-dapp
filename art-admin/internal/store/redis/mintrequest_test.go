package redis

import (
	"context"
	"fmt"
	"testing"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
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
				Id:                 "0x1",
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
		Address:  "localhost:6379",
		CacheTTL: "1s",
	}
	ctx := context.Background()

	rdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := rdb.MintRequestStore(ctx)
	is.NoErr(err)

	nftMR, images := getTestNFTMintRequest()

	_, err = ms.NewNFTMintRequest(ctx, nftMR, images)
	is.NoErr(err)

	nftMRs, err := ms.GetAllNFTMintRequests(ctx)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, StatusUnknown.String())

	_, err = ms.UpdateStatusNFTMintRequest(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		StatusPending)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, StatusPending.String())

	url := "offchain url"
	upd, err := ms.UpdateNFTOffchainUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		url)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, upd.Status)
	is.Equal(nftMRs[0].NftOffchainUrl, url)

	_, err = ms.DeleteNFTOffchainUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.GetId()))
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, StatusPending.String())
	is.Equal(nftMRs[0].NftOffchainUrl, "")

	err = ms.DeleteNFTMintRequestById(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.GetId()),
	)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx)
	is.NoErr(err)
	is.Equal(len(nftMRs), 0)

}

func TestPagination(t *testing.T) {
	is := is.New(t)

	c := Config{
		Address:  "localhost:6379",
		CacheTTL: "1s",
	}
	ctx := context.Background()

	rdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := rdb.MintRequestStore(ctx)
	is.NoErr(err)

	nftMR, images := getTestNFTMintRequest()

	for i := 0; i < 100; i++ {
		_, err = ms.NewNFTMintRequest(ctx, nftMR, images)
		is.NoErr(err)
	}
	// time.Sleep(1 * time.Second)

	mrs, err := ms.GetNFTMintRequestsPaged(ctx, StatusPending, 1)
	is.NoErr(err)
	fmt.Println("--- ", len(mrs))

}
