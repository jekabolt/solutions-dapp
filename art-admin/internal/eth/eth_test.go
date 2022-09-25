package eth

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	systoken "github.com/jekabolt/solutions-dapp/art-admin/contract"
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

type testStore struct {
	mrs map[string]*pb_nft.NFTMintRequestWithStatus
}

func (ts *testStore) GetAllNFTMintRequests(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error) {
	var res []*pb_nft.NFTMintRequestWithStatus
	for _, v := range ts.mrs {
		res = append(res, v)
	}
	return res, nil
}

func (ts *testStore) UpdateStatusNFTMintRequest(ctx context.Context, id string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error) {
	ts.mrs[id].Status = status
	return ts.mrs[id], nil
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
				Id:                 fmt.Sprint(mintSequenceNumber),
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

	mr1Up, imgs1 := getTestNFTMintRequest(1)
	mr2Up, imgs2 := getTestNFTMintRequest(2)
	ts := &testStore{
		mrs: map[string]*pb_nft.NFTMintRequestWithStatus{
			"1": {
				NftMintRequest: mr1Up.NftMintRequest,
				Status:         pb_nft.Status_Pending,
				SampleImages:   imgs1,
			},
			"2": {
				NftMintRequest: mr2Up.NftMintRequest,
				Status:         pb_nft.Status_Pending,
				SampleImages:   imgs2,
			},
		},
	}
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))

	cli := &Client{
		c: &Config{
			Retries: 3,
		},
		tokenObserver: &testObserver{},
		mintStore:     ts,
		checkInterval: interval,
		ttlMap:        make(map[string]int),
	}

	cli.Run(context.Background())

	tick := time.NewTicker(time.Millisecond * 1000)

loop:
	for {
		select {
		case <-tick.C:
			nftMRs, err := ts.GetAllNFTMintRequests(ctx, pb_nft.Status_Any)
			is.NoErr(err)
			is.True(len(nftMRs) == 2)
			if nftMRs[0].Status == pb_nft.Status_Pending &&
				nftMRs[1].Status == pb_nft.Status_Failed {
				cli.Stop()
				tick.Stop()
				break loop
			}
		case <-ctx.Done():
			break loop
		}
	}

	nftMRs, err := ts.GetAllNFTMintRequests(ctx, pb_nft.Status_Any)
	is.NoErr(err)
	is.Equal(nftMRs[0].Status, pb_nft.Status_Pending)
	is.Equal(nftMRs[1].Status, pb_nft.Status_Failed)
	t.Logf("---- nftMRs %+v", nftMRs)
	cancel()
}
