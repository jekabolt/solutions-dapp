package nft

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/redis"
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

// db mock
func Store() (redis.Store, error) {
	rc := getRedisAddress()
	c := redis.Config{
		Address:  rc.Host,
		Password: rc.Password,
		CacheTTL: "1s",
		PageSize: 30,
	}
	ctx := context.Background()

	rdb, err := c.InitDB(ctx)
	if err != nil {
		return nil, err
	}

	return rdb, nil
}

// file store mock
type fs struct{}

func (_ fs) UploadContentImage(rawB64Image string, pe *bucket.PathExtra) (*pb_nft.ImageList, error) {
	return &pb_nft.ImageList{
		FullSize:   "https://example.com/full.jpg",
		Compressed: "https://example.com/compressed.jpg",
	}, nil
}
func (_ fs) UploadMetadata(metadata map[int]bucket.Metadata) (string, error) {
	return "http://example.com/metadata.json", nil
}

func newFileStore() bucket.FileStore {
	return fs{}
}

// ipfs mock
type ipfsStore struct {
	NFTTotalSupply int
}

func (ipfs *ipfsStore) BulkUploadIPFS(mrs []*pb_nft.NFTMintRequestWithStatus) (map[int]bucket.Metadata, error) {
	meta := map[int]bucket.Metadata{}

	for i := 1; i <= ipfs.NFTTotalSupply; i++ {
		meta[i] = bucket.Metadata{
			Name:        fmt.Sprintf("test-%d", i),
			Description: fmt.Sprintf("description-%d", i),
			Image:       fmt.Sprintf("https://example.com/image-%d.jpg", i),
			Edition:     i,
			Date:        time.Now().Unix(),
		}
	}
	return meta, nil
}

func newIpfs(nts int) ipfs.IPFS {
	return &ipfsStore{
		NFTTotalSupply: nts,
	}
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

	db, err := Store()
	is.NoErr(err)
	// upload mint request
	s, err := c.New(db, newFileStore(), newIpfs(c.NFTTotalSupply), newDescriptions(is))
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
