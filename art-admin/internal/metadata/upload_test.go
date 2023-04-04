package metadata

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/teststore"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

const (
	baseExternalUrl = "https://nft.sys.solutions/nft/%d"
	baseImageUrl    = "https://nft.sys.solutions/nft/%d/image"
	totalQuantity   = 10000
	namePrefix      = "Solutions art #%d"

	MoralisApiKey  = "YNK1oJejsgzJ1L5Gxaszqd1fOH5t5h595ksVu5bvE6nyCDl9Bb7WD7N18Gb6mglz"
	MoralisTimeout = "50s"
	MoralisBaseURL = "https://deep-index.moralis.io/api/v2/"

	ipfsUrl = "ipfs://test"
)

type uploader struct{}

func (u *uploader) UploadData(data []byte) (string, error) {
	return ipfsUrl, nil
}

func TestInitialUpload(t *testing.T) {
	skipCI(t)
	is := is.New(t)

	cd := descriptions.Config{
		CollectionName:  "test",
		CountPerEdition: 3,
		TotalCount:      totalQuantity,
		RandSeed:        1337,
	}
	d := cd.New(totalQuantity)

	s := teststore.NewTestStore(30)

	// ci := ipfs.Config{
	// 	APIKey:  MoralisApiKey,
	// 	Timeout: MoralisTimeout,
	// 	BaseURL: MoralisBaseURL,
	// }
	// i, err := ci.New()

	c := Config{
		BaseExternalUrl: baseExternalUrl,
		BaseImageUrl:    baseImageUrl,
		TotalQuantity:   totalQuantity,
		NamePrefix:      namePrefix,
		DefaultAuthor:   "some author",
		UploadRetries:   3,
	}
	u := &uploader{}

	mm, err := c.New(d, s, s, u)
	is.NoErr(err)

	ctx := context.Background()

	// need to upload initial metadata before uploading ipfs
	mi, err := mm.UploadIpfs(ctx)
	is.True(err != nil)

	mi, err = mm.UploadInitial(ctx)
	is.NoErr(err)
	is.True(mi.Processing)

	time.Sleep(time.Second / 2)

	md, err := s.GetAllMetadata(ctx)
	is.NoErr(err)
	is.True(!md[0].MetaInfo.Processing)
	is.True(md[0].MetaInfo.IpfsUrl == ipfsUrl)

	// should fail because nothing to upload
	mi, err = mm.UploadIpfs(ctx)
	is.True(err != nil)

	s.AddMockData([]pb_nft.Status{
		pb_nft.Status_Uploaded,
	}, 39)

	mi, err = mm.UploadIpfs(ctx)
	is.NoErr(err)
	is.True(mi.Processing)

	time.Sleep(time.Second / 2)

	md, err = s.GetAllMetadata(ctx)
	is.NoErr(err)
	is.True(!md[0].MetaInfo.Processing)
	is.True(md[0].MetaInfo.IpfsUrl == ipfsUrl)

}
