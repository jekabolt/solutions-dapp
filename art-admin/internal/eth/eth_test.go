package eth

import (
	"context"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	systoken "github.com/jekabolt/solutions-dapp/art-admin/contract"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/bunt"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/matryer/is"
)

const (
	contractAddress    = "0x9f7bdb481eacaa02089dd35269084a6355158192"
	MintSequenceNumber = 1
)

func skipCI(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping testing in CI environment")
	}
}

func TestConnect(t *testing.T) {
	skipCI(t)
	is := is.New(t)
	client, err := ethclient.Dial("ws://127.0.0.1:8545")
	is.NoErr(err)

	// ctx := context.Background()

	address := common.HexToAddress(contractAddress)
	instance, err := systoken.NewSystoken(address, client)
	is.NoErr(err)

	ts, err := instance.OwnerById(nil, big.NewInt(3))
	is.NoErr(err)
	t.Log(ts)
}

type testObserver struct {
	paid bool
}

func (c *testObserver) IsPaid(mr *pb_nft.NFTMintRequestWithStatus) (bool, error) {
	if mr.NftMintRequest.MintSequenceNumber == 1 {
		if !c.paid {
			c.paid = true
			return false, nil
		}
		return c.paid, nil
	}
	return false, nil
}

func getTestNFTMintRequest(mintSequenceNumber int32) (*pb_nft.NFTMintRequestToUpload, []*pb_nft.ImageList) {
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
				MintSequenceNumber: mintSequenceNumber,
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

func TestWatcher(t *testing.T) {
	is := is.New(t)

	interval, err := time.ParseDuration("100ms")
	is.NoErr(err)

	bc := bunt.Config{
		DBPath: ":memory:",
	}

	bdb, err := bc.InitDB()
	is.NoErr(err)

	mrStore := bdb.MintRequestStore()

	_, err = mrStore.UpsertNFTMintRequest(getTestNFTMintRequest(1))
	is.NoErr(err)

	_, err = mrStore.UpsertNFTMintRequest(getTestNFTMintRequest(2))
	is.NoErr(err)

	cli := &Client{
		c: &Config{
			Retries: 3,
		},
		tokenObserver: &testObserver{},
		mintStore:     mrStore,
		checkInterval: interval,
		ttlMap:        make(map[int]int),
	}

	cli.Run(context.Background())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	tick := time.NewTicker(time.Millisecond * 1000)

loop:
	for {
		select {
		case <-tick.C:
			nftMRs, err := mrStore.GetAllNFTMintRequests()
			is.NoErr(err)
			is.True(len(nftMRs) == 2)
			if nftMRs[0].Status == bunt.StatusPending.String() &&
				nftMRs[1].Status == bunt.StatusFailed.String() {
				cli.Stop()
				cancel()
				tick.Stop()
				break loop
			}
		case <-ctx.Done():
			cancel()
			break loop
		}
	}

	nftMRs, err := mrStore.GetAllNFTMintRequests()
	is.NoErr(err)
	is.Equal(nftMRs[0].Status, bunt.StatusPending.String())
	is.Equal(nftMRs[1].Status, bunt.StatusFailed.String())
	t.Logf("---- nftMRs %+v", nftMRs)
}
