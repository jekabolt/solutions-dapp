package mongo

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
			EthAddress:  "0x0",
			Description: "test",
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

	ctx := context.Background()

	c := Config{
		DSN: getMongoDSN(),
		DB:  "test",
	}
	mdb, err := c.InitDB(ctx)
	is.NoErr(err)

	defer func() {
		err = mdb.mintRequests.Drop(ctx)
		is.NoErr(err)
	}()

	ms, err := mdb.MintRequestStore(ctx)
	is.NoErr(err)

	nftMR, images := getTestNFTMintRequest(1)

	_, err = ms.New(ctx, nftMR, images)
	is.NoErr(err)

	nftMRs, err := ms.GetAll(ctx, pb_nft.Status_Unknown)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Unknown)

	_, err = ms.UpdateStatus(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		pb_nft.Status_Pending)
	is.NoErr(err)

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_Pending)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Pending)

	offchainUrl := "offchain url"

	_, err = ms.UpdateOffchainUrl(ctx, nftMRs[0].NftMintRequest.Id, offchainUrl)
	is.NoErr(err)

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_UploadedOffchain)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_UploadedOffchain)
	is.Equal(nftMRs[0].OffchainUrl, offchainUrl)

	ipfsUrl := "ipfs url"
	_, err = ms.UpdateIpfsUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		ipfsUrl)
	is.NoErr(err)

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_Uploaded)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].GetStatus(), pb_nft.Status_Uploaded)
	is.Equal(nftMRs[0].OnchainUrl, ipfsUrl)

	_, err = ms.DeleteIpfsUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.GetId()))
	is.NoErr(err)

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_UploadedOffchain)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].Status, pb_nft.Status_UploadedOffchain)
	is.Equal(nftMRs[0].OnchainUrl, "")

	_, err = ms.UpdateIpfsUrl(
		ctx,
		fmt.Sprint(nftMRs[0].NftMintRequest.Id),
		ipfsUrl)
	is.NoErr(err)

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_Uploaded)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	is.Equal(nftMRs[0].GetStatus(), pb_nft.Status_Uploaded)
	is.Equal(nftMRs[0].OnchainUrl, ipfsUrl)

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

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_Burned)
	is.NoErr(err)
	is.Equal(len(nftMRs), 1)
	fmt.Printf("%+v", nftMRs[0])
	is.Equal(nftMRs[0].Status, pb_nft.Status_Burned)
	is.Equal(nftMRs[0].Shipping.Shipping.FullName, "test")

	burnedWTrackMr, err := ms.UpdateTrackingNumber(ctx, &pb_nft.SetTrackingNumberRequest{
		Id:             nftMRs[0].NftMintRequest.GetId(),
		TrackingNumber: "testTrack",
	})
	is.NoErr(err)
	is.Equal(burnedWTrackMr.Status, pb_nft.Status_Shipped)
	is.Equal(burnedWTrackMr.Shipping.TrackingNumber, "testTrack")

	err = ms.DeleteMintById(
		ctx,
		fmt.Sprint(burnedWTrackMr.NftMintRequest.GetId()),
	)
	is.NoErr(err)

	nftMRs, err = ms.GetAll(ctx, pb_nft.Status_Any)
	is.NoErr(err)
	is.Equal(len(nftMRs), 0)

}
func TestGetAllToUpload(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	c := Config{
		DSN: getMongoDSN(),
		DB:  "test",
	}
	mdb, err := c.InitDB(ctx)
	is.NoErr(err)

	ms, err := mdb.MintRequestStore(ctx)
	is.NoErr(err)

	mrsCreated := make([]*pb_nft.NFTMintRequestWithStatus, 0)
	for i := 0; i < 100; i++ {
		nftMR, images := getTestNFTMintRequest(int32(i))
		mr, err := ms.New(ctx, nftMR, images)
		mrsCreated = append(mrsCreated, mr)
		is.NoErr(err)
	}
	is.Equal(len(mrsCreated), 100)
	defer func() {
		err = mdb.mintRequests.Drop(ctx)
		is.NoErr(err)
	}()

	mrs, err := ms.GetAll(ctx, pb_nft.Status_Any)
	is.NoErr(err)

	_, err = ms.UpdateStatus(ctx, mrs[0].NftMintRequest.Id, pb_nft.Status_Uploaded)
	is.NoErr(err)
	_, err = ms.UpdateStatus(ctx, mrs[1].NftMintRequest.Id, pb_nft.Status_UploadedOffchain)
	is.NoErr(err)
	_, err = ms.UpdateStatus(ctx, mrs[2].NftMintRequest.Id, pb_nft.Status_Burned)
	is.NoErr(err)
	_, err = ms.UpdateStatus(ctx, mrs[3].NftMintRequest.Id, pb_nft.Status_Shipped)
	is.NoErr(err)

	toUpload, err := ms.GetAllToUpload(ctx)
	is.NoErr(err)
	is.Equal(len(toUpload), 4)
	is.Equal(len(mrsCreated), 100)

}

func TestPagination(t *testing.T) {
	is := is.New(t)

	ctx := context.Background()
	c := Config{
		DSN:      "mongodb+srv://sol:H3Xw6542Mx7D08JW@mongo-sol-35718b1e.mongo.ondigitalocean.com/test?tls=true&authSource=admin&replicaSet=mongo-sol",
		DB:       "test",
		PageSize: 30,
	}
	mdb, err := c.InitDB(ctx)
	is.NoErr(err)
	ms, err := mdb.MintRequestStore(ctx)
	is.NoErr(err)

	mrsCreated := make([]*pb_nft.NFTMintRequestWithStatus, 0)
	for i := 0; i < 100; i++ {
		nftMR, images := getTestNFTMintRequest(int32(i))
		mr, err := ms.New(ctx, nftMR, images)
		mrsCreated = append(mrsCreated, mr)
		is.NoErr(err)
	}

	defer func() {
		err = mdb.mintRequests.Drop(ctx)
		is.NoErr(err)
	}()

	mrsPg1, err := ms.GetPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg1), int(c.PageSize))

	mrsPg2, err := ms.GetPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   2,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg2), int(c.PageSize))

	mrsPg3, err := ms.GetPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   3,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg3), int(c.PageSize))

	mrsPg4, err := ms.GetPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Unknown,
		Page:   4,
	})
	is.NoErr(err)
	is.Equal(len(mrsPg4), 10)

	mrsPg5, err := ms.GetPaged(ctx, &pb_nft.ListPagedRequest{
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
