package metadata

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/teststore"
	pb_metadata "github.com/jekabolt/solutions-dapp/art-admin/proto/metadata"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"

	"github.com/matryer/is"
)

// ipfs mock
type ipfsStore struct {
	NFTTotalSupply int
}

func (ipfs *ipfsStore) BulkUploadIPFS(mrs []*pb_nft.NFTMintRequestWithStatus) (map[int]pb_metadata.MetadataUnit, error) {
	meta := map[int]pb_metadata.MetadataUnit{}

	for i := 1; i <= ipfs.NFTTotalSupply; i++ {
		meta[i] = pb_metadata.MetadataUnit{
			Name:          fmt.Sprintf("test-%d", i),
			Description:   fmt.Sprintf("description-%d", i),
			OffchainImage: fmt.Sprintf("https://example.com/offchain-image-%d.jpg", i),
			OnchainImage:  fmt.Sprintf("https://example.com/onchain-image-%d.jpg", i),
			Edition:       int32(i),
			Date:          int32(time.Now().Unix()),
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

func TestMetadata(t *testing.T) {
	is := is.New(t)

	c := Config{}

	db := teststore.NewTestStore(30)
	db.AddMockData([]pb_nft.Status{
		pb_nft.Status_UploadedOffchain,
	}, 10)

	// upload mint request
	s := c.New(db, newIpfs(10000), newDescriptions(is))
	ctx := context.Background()

	_, err := s.UploadOffchainMetadata(ctx, nil)
	is.NoErr(err)

	md, err := s.GetAllMetadata(ctx, nil)
	is.NoErr(err)
	is.True(len(md.MetaInfo) == 1)
	is.True(md.MetaInfo[0].IpfsUrl == "")

	_, err = s.UploadIPFSMetadata(ctx, &pb_metadata.UploadIPFSMetadataRequest{
		Key: md.MetaInfo[0].Key,
	})
	is.NoErr(err)

	_, err = s.DeleteIPFSMetadata(ctx, &pb_metadata.DeleteIPFSMetadataRequest{
		Key: md.MetaInfo[0].Key,
	})
	is.NoErr(err)
}
