package nft

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/teststore"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/matryer/is"
)

func getRedisAddress() redis.RedisConf {
	if os.Getenv("REDIS_HOST") == "" {
		return redis.RedisConf{
			Host: "localhost:6379",
		}
	}
	return redis.RedisConf{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
	}
}

// file store mock
type fs struct{}

func (_ fs) UploadContentImage(rawB64Image string, pe *bucket.PathExtra) (*pb_nft.ImageList, error) {
	return &pb_nft.ImageList{
		FullSize:   "https://example.com/full.jpg",
		Compressed: "https://example.com/compressed.jpg",
	}, nil
}

func newFileStore() bucket.FileStore {
	return fs{}
}

// descriptions

func newDescriptions(is *is.I) *descriptions.Store {
	c := descriptions.Config{
		Path:           "../../etc/descriptions.json",
		CollectionName: "test",
	}
	ds, err := c.Init()
	is.NoErr(err)
	return ds
}

func TestNft(t *testing.T) {
	is := is.New(t)

	c := Config{
		NFTTotalSupply: 100,
	}

	db := teststore.NewTestStore(30)
	// upload mint request
	s, err := c.New(db, newFileStore())
	is.NoErr(err)
	ctx := context.Background()
	resp, err := s.NewNFTMintRequest(ctx, &pb_nft.NFTMintRequestToUpload{
		NftMintRequest: &pb_nft.NFTMintRequest{
			Id:                 "",
			EthAddress:         "0x0",
			TxHash:             "0x0",
			MintSequenceNumber: 1,
			Description:        "test",
		},
		SampleImages: []*pb_nft.ImageToUpload{
			{
				Raw: "https://grbpwr.com/img/small-logo.png",
			},
		},
	})

	defer func() {
		err = db.DeleteNFTMintRequestById(ctx, resp.GetNftMintRequest().GetId())
		is.NoErr(err)
	}()
	is.NoErr(err)
	is.Equal(resp.Status, pb_nft.Status_Unknown)
	time.Sleep(time.Second)
	// list mint requests
	list, err := s.ListNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: resp.Status,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 1)

	_, err = s.db.UpdateStatusNFTMintRequest(
		ctx,
		list.NftMintRequests[0].NftMintRequest.Id,
		pb_nft.Status_Pending)
	is.NoErr(err)

	list, err = s.ListNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Pending,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 1)

	// update mint offchain url
	resp, err = s.UpdateNFTOffchainUrl(ctx, &pb_nft.UpdateNFTOffchainUrlRequest{
		Id: list.NftMintRequests[0].NftMintRequest.Id,
		NftOffchainUrl: &pb_nft.ImageToUpload{
			Raw: "https://example.com/offchain.jpg",
		},
	})
	is.NoErr(err)
	is.Equal(resp.Status, pb_nft.Status_UploadedOffchain)

	list, err = s.ListNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_UploadedOffchain,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 1)

	_, err = s.db.UpdateStatusNFTMintRequest(
		ctx,
		list.NftMintRequests[0].NftMintRequest.Id,
		pb_nft.Status_Uploaded)
	is.NoErr(err)

	list, err = s.ListNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Uploaded,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 1)

	_, err = s.Burn(ctx, &pb_nft.BurnRequest{
		Id: list.NftMintRequests[0].NftMintRequest.Id,
		Shipping: &pb_nft.ShippingTo{
			FullName: "test",
		},
	})
	is.NoErr(err)

	list, err = s.ListNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Burned,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 1)
	is.True(list.NftMintRequests[0].Shipping.Shipping.FullName == "test")

	_, err = s.SetTrackingNumber(ctx, &pb_nft.SetTrackingNumberRequest{
		Id:             list.NftMintRequests[0].NftMintRequest.Id,
		TrackingNumber: "test",
	})
	is.NoErr(err)

	list, err = s.ListNFTMintRequestsPaged(ctx, &pb_nft.ListPagedRequest{
		Status: pb_nft.Status_Shipped,
		Page:   1,
	})
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 1)
	is.True(list.NftMintRequests[0].Shipping.TrackingNumber == "test")

}
