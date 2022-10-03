package redis

import (
	"context"
	"fmt"
	"testing"

	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
)

func getTestNFTMintRequest(msn int32) (*pb_nft.NFTMintRequestToUpload, []*pb_nft.ImageList) {
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
				MintSequenceNumber: msn,
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

	rc := getRedisAddress()
	c := Config{
		Address:  rc.Host,
		Password: rc.Password,
		CacheTTL: "1s",
		PageSize: 30,
	}
	ctx := context.Background()

	rdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := rdb.MintRequestStore(ctx)
	is.NoErr(err)

	nftMR, images := getTestNFTMintRequest(1)

	mrCreated, err := ms.NewNFTMintRequest(ctx, nftMR, images)
	is.NoErr(err)

	ok := false
	defer func() {
		if !ok {
			err = ms.DeleteNFTMintRequestById(
				ctx,
				fmt.Sprint(mrCreated.NftMintRequest.GetId()),
			)
			is.NoErr(err)
		}
	}()

	nftMRs, err := ms.GetAllNFTMintRequests(ctx, pb_nft.Status_Unknown)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Unknown)

	_, err = ms.UpdateStatusNFTMintRequest(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		pb_nft.Status_Pending)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_Pending)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Pending)

	url := "offchain url"
	_, err = ms.UpdateNFTOffchainUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		url)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_UploadedOffchain)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].GetStatus(), pb_nft.Status_UploadedOffchain)
	is.Equal(nftMRs[0].NftOffchainUrl, url)

	_, err = ms.DeleteNFTOffchainUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.GetId()))
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_Pending)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Pending)
	is.Equal(nftMRs[0].NftOffchainUrl, "")

	_, err = ms.UpdateNFTOffchainUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		url)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_UploadedOffchain)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].GetStatus(), pb_nft.Status_UploadedOffchain)
	is.Equal(nftMRs[0].NftOffchainUrl, url)

	uploadedMr, err := ms.UpdateStatusNFTMintRequest(ctx, nftMRs[0].NftMintRequest.GetId(), pb_nft.Status_Uploaded)
	is.NoErr(err)
	is.Equal(uploadedMr.Status, pb_nft.Status_Uploaded)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_Uploaded)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].GetStatus(), pb_nft.Status_Uploaded)

	burnedMr, err := ms.UpdateShippingInfo(ctx, &pb_nft.BurnRequest{
		Id: nftMRs[0].NftMintRequest.GetId(),
		Shipping: &pb_nft.ShippingTo{
			FullName: "test",
			Address:  "addr",
			ZipCode:  "00001",
			City:     "testCity",
			Country:  "testCountry",
			Email:    "test@grbpwr.com",
		},
	})
	is.NoErr(err)
	is.True(burnedMr.Shipping != nil)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_Burned)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Burned)
	is.Equal(nftMRs[0].Shipping.Shipping.FullName, "test")

	burnedWTrackMr, err := ms.UpdateTrackingNumber(ctx, &pb_nft.SetTrackingNumberRequest{
		Id:             nftMRs[0].NftMintRequest.GetId(),
		TrackingNumber: "testTrack",
	})
	is.NoErr(err)
	is.Equal(burnedWTrackMr.Status, pb_nft.Status_Shipped)
	is.Equal(burnedWTrackMr.Shipping.TrackingNumber, "testTrack")

	err = ms.DeleteNFTMintRequestById(
		ctx,
		fmt.Sprint(mrCreated.NftMintRequest.GetId()),
	)
	is.NoErr(err)

	nftMRs, err = ms.GetAllNFTMintRequests(ctx, pb_nft.Status_Any)
	is.NoErr(err)
	is.Equal(len(nftMRs), 0)

	ok = true

}
func TestGetAllToUpload(t *testing.T) {
	is := is.New(t)

	rc := getRedisAddress()
	c := Config{
		Address:  rc.Host,
		Password: rc.Password,
		CacheTTL: "1s",
		PageSize: 30,
	}
	ctx := context.Background()

	rdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := rdb.MintRequestStore(ctx)
	is.NoErr(err)

	mrsCreated := make([]*pb_nft.NFTMintRequestWithStatus, 0)
	for i := 0; i < 100; i++ {
		nftMR, images := getTestNFTMintRequest(int32(i))
		mr, err := ms.NewNFTMintRequest(ctx, nftMR, images)
		mrsCreated = append(mrsCreated, mr)
		is.NoErr(err)
	}

	defer func() {
		for _, mr := range mrsCreated {
			err = ms.DeleteNFTMintRequestById(ctx, mr.GetNftMintRequest().GetId())
			is.NoErr(err)
		}
		rdb.Close()
	}()

	_, err = ms.UpdateStatusNFTMintRequest(ctx, mrsCreated[0].NftMintRequest.GetId(), pb_nft.Status_Uploaded)
	is.NoErr(err)
	_, err = ms.UpdateStatusNFTMintRequest(ctx, mrsCreated[1].NftMintRequest.GetId(), pb_nft.Status_UploadedOffchain)
	is.NoErr(err)
	_, err = ms.UpdateStatusNFTMintRequest(ctx, mrsCreated[2].NftMintRequest.GetId(), pb_nft.Status_Burned)
	is.NoErr(err)
	_, err = ms.UpdateStatusNFTMintRequest(ctx, mrsCreated[3].NftMintRequest.GetId(), pb_nft.Status_Shipped)
	is.NoErr(err)

	toUpload, err := ms.GetAllToUpload(ctx)
	is.NoErr(err)
	is.Equal(len(toUpload), 4)
	is.Equal(len(mrsCreated), 100)
}

func TestPagination(t *testing.T) {
	is := is.New(t)

	rc := getRedisAddress()
	c := Config{
		Address:  rc.Host,
		Password: rc.Password,
		CacheTTL: "1s",
		PageSize: 30,
	}
	ctx := context.Background()

	rdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := rdb.MintRequestStore(ctx)
	is.NoErr(err)

	mrsCreated := make([]*pb_nft.NFTMintRequestWithStatus, 0)
	for i := 0; i < 100; i++ {
		nftMR, images := getTestNFTMintRequest(int32(i))
		mr, err := ms.NewNFTMintRequest(ctx, nftMR, images)
		mrsCreated = append(mrsCreated, mr)
		is.NoErr(err)
	}

	defer func() {
		for _, mr := range mrsCreated {
			err = ms.DeleteNFTMintRequestById(ctx, mr.GetNftMintRequest().GetId())
			is.NoErr(err)
		}
		rdb.Close()
	}()

	mrsPg1, err := ms.GetNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg1), int(c.PageSize))

	mrsPg2, err := ms.GetNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   2,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg2), int(c.PageSize))

	mrsPg3, err := ms.GetNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   3,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg3), int(c.PageSize))

	mrsPg4, err := ms.GetNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   4,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg4), 10)

	mrsPg5, err := ms.GetNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   5,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg5), 0)

	mrs := append(mrsPg1, mrsPg2...)
	mrs = append(mrs, mrsPg3...)
	mrs = append(mrs, mrsPg4...)

	prev := int32(0)
	for _, mr := range mrs {
		is.True(mr.NftMintRequest.MintSequenceNumber == prev)
		prev++
	}

}
