package nft

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/bunt"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/matryer/is"
)

// db
func newDB(is *is.I) bunt.Store {
	c := bunt.Config{
		DBPath: ":memory:",
	}
	bdb, err := c.InitDB()
	is.NoErr(err)
	return bdb
}

// file store
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

// ipfs
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

	db := newDB(is)
	// upload mint request
	s, err := c.New(db, newFileStore(), newIpfs(c.NFTTotalSupply), newDescriptions(is))
	is.NoErr(err)
	ctx := context.Background()
	resp, err := s.UpsertNFTMintRequest(ctx, &pb_nft.NFTMintRequestToUpload{
		NftMintRequest: &pb_nft.NFTMintRequest{
			Id:                 0,
			EthAddress:         "0x0",
			TxHash:             "0x0",
			MintSequenceNumber: 1,
			Description:        "test",
		},
		SampleImages: []*pb_nft.ImageToUpload{
			{
				Raw: "some b64 image",
			},
		},
	})
	is.NoErr(err)
	is.Equal(resp.Status, bunt.StatusUnknown.String())

	resp, err = s.UpsertNFTMintRequest(ctx, &pb_nft.NFTMintRequestToUpload{
		NftMintRequest: &pb_nft.NFTMintRequest{
			Id:                 0,
			EthAddress:         "0x0",
			TxHash:             "0x0",
			MintSequenceNumber: 1,
			Description:        "test",
		},
		SampleImages: []*pb_nft.ImageToUpload{
			{
				Raw: "some b64 image",
			},
		},
	})

	is.NoErr(err)
	is.Equal(resp.Status, bunt.StatusUnknown.String())

	// list mint requests
	list, err := s.ListNFTMintRequests(ctx, nil)
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 2)

	// update mint offchain url
	resp, err = s.UpdateNFTOffchainUrl(ctx, &pb_nft.UpdateNFTOffchainUrlRequest{
		Id:             fmt.Sprint(list.NftMintRequests[0].NftMintRequest.Id),
		NftOffchainUrl: "https://example.com/offchain.jpg",
	})
	is.NoErr(err)
	is.Equal(resp.Status, bunt.StatusUploadedOffchain.String())

	// upload offchain metadata
	_, err = s.UploadOffchainMetadata(ctx, nil)
	is.NoErr(err)

	all, err := db.GetAllToUpload()
	is.NoErr(err)
	is.Equal(len(all), 1)

	// list mint requests check offchain url
	list, err = s.ListNFTMintRequests(ctx, nil)
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 2)
	is.Equal(list.NftMintRequests[0].NftOffchainUrl, "https://example.com/offchain.jpg")
	is.Equal(list.NftMintRequests[0].Status, bunt.StatusUploadedOffchain.String())

	// delete offchain url
	resp, err = s.DeleteNFTOffchainUrl(ctx, &pb_nft.DeleteId{
		Id: fmt.Sprint(list.NftMintRequests[0].NftMintRequest.Id),
	})
	is.NoErr(err)
	is.Equal(resp.Status, bunt.StatusUnknown.String())
	is.Equal(resp.NftOffchainUrl, "")

	// list mint requests check offchain url
	list, err = s.ListNFTMintRequests(ctx, nil)
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 2)
	is.Equal(list.NftMintRequests[0].NftOffchainUrl, "")
	is.Equal(list.NftMintRequests[0].Status, bunt.StatusUnknown.String())

	// delete mint request
	for _, mr := range list.NftMintRequests {
		s.DeleteNFTMintRequestById(ctx, &pb_nft.DeleteId{
			Id: fmt.Sprint(mr.NftMintRequest.Id),
		})
	}

	// check if deleted
	list, err = s.ListNFTMintRequests(ctx, nil)
	is.NoErr(err)
	is.Equal(len(list.NftMintRequests), 0)

}
