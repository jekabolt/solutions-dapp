package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	systoken "github.com/jekabolt/solutions-dapp/art-admin/contract"
	pb_nft "github.com/jekabolt/solutions-dapp/art-admin/proto/nft"
	"github.com/rs/zerolog/log"
)

type Config struct {
	NodeAddress     string `env:"ETH_NODE_ADDRESS"`
	ContractAddress string `env:"ETH_CONTRACT_ADDRESS"`
	PrivateKey      string `env:"ETH_PRIVATE_KEY"`
	Retries         int    `env:"ETH_RETRIES"`
	CheckInterval   string `env:"ETH_CHECK_INTERVAL" envDefault:"10m"`
	WatcherDisabled bool   `env:"ETH_WATCHER_DISABLED" envDefault:"true"`
}

// Watcher interface for checking and updating statuses for mint requests once they got paid or failed
type Watcher interface {
	Run(ctx context.Context)
	IsPaid(mr *pb_nft.NFTMintRequestWithStatus) (bool, error)
	MintStatusUpdate() error
}

// UnknownUpdate db interface for getting all unknown mint requests and update status on em
type UnknownUpdate interface {
	// used fro getting all unknown mint requests
	GetAllNFTMintRequests(ctx context.Context, status pb_nft.Status) ([]*pb_nft.NFTMintRequestWithStatus, error)
	UpdateStatusNFTMintRequest(ctx context.Context, id string, status pb_nft.Status) (*pb_nft.NFTMintRequestWithStatus, error)
}

// TokenObserver interface for interacting with the token contract
type TokenObserver interface {
	IsPaid(mr *pb_nft.NFTMintRequestWithStatus) (bool, error)
}

// Client is the client for the Ethereum contract
type Client struct {
	c             *Config
	mintStore     UnknownUpdate
	tokenObserver TokenObserver
	checkInterval time.Duration
	ttlMap        map[string]int // map[redisId]count
	cancel        context.CancelFunc
}

// New creates a new client for the Ethereum contract
func (c *Config) New(ctx context.Context, mintRequestStore UnknownUpdate) (cli *Client, err error) {
	cli = &Client{
		c:         c,
		mintStore: mintRequestStore,
	}

	if c.WatcherDisabled {
		log.Info().Msgf("eth watcher disabled")
		return
	}

	ethCli, err := ethclient.Dial(c.NodeAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the Ethereum client: %v", err)
	}
	address := common.HexToAddress(c.ContractAddress)

	cli.tokenObserver, err = systoken.NewSystoken(address, ethCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a token contract: %v", err)
	}
	cli.checkInterval, err = time.ParseDuration(c.CheckInterval)
	if err != nil {
		return nil, fmt.Errorf("invalid check interval: %v", err)
	}
	cli.ttlMap = make(map[string]int)

	return cli, nil
}

// Run starts the transaction status update loop
// it will run until the context is cancelled
// once tx is paid, it will update the status to pb_nft.Status_Pending
func (cli *Client) Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	cli.cancel = cancel
	if cli.c.WatcherDisabled {
		return
	}
	t := time.NewTicker(cli.checkInterval)
	go func() {
		for {
			select {
			case <-t.C:
				log.Debug().Msg("Run: processing")
				err := cli.MintStatusUpdate(ctx)
				if err != nil {
					log.Error().Msgf("Run:MintStatusUpdate [%s]", err.Error())
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

// Stop stops the transaction status update loop
func (cli *Client) Stop() {
	log.Debug().Msg("Stop: stop")
	if cli != nil && cli.cancel != nil {
		cli.cancel()
	}
}

// MintStatusUpdate get all pb_nft.Status_Unknown mint requests and check if mint is paid
// if paid, update the status to pb_nft.Status_Pending
// if not paid, update the status to pb_nft.Status_Failed
func (cli *Client) MintStatusUpdate(ctx context.Context) error {
	mrs, err := cli.mintStore.GetAllNFTMintRequests(ctx, pb_nft.Status_Unknown)
	if err != nil {
		return err
	}

	for _, mr := range mrs {
		if mr.Status == pb_nft.Status_Unknown ||
			mr.Status == pb_nft.Status_Pending {

			cli.ttlMap[mr.NftMintRequest.Id]++

			// check if the transaction passed the max count of retries
			if cli.ttlMap[mr.NftMintRequest.Id] >= cli.c.Retries {
				// update the status to failed
				_, err := cli.mintStore.UpdateStatusNFTMintRequest(ctx, fmt.Sprint(mr.NftMintRequest.Id), pb_nft.Status_Failed)
				if err != nil {
					log.Error().Msgf("MintStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
				// delete the mint request from the map on status update
				delete(cli.ttlMap, mr.NftMintRequest.Id)
				continue
			}

			// check if mint is paid
			ok, err := cli.tokenObserver.IsPaid(mr)
			if err != nil {
				log.Debug().Msgf("MintStatusUpdate:cli.isPaid [%v]", err)
				continue
			}

			if ok {
				// update the status to pending
				// TODO:
				// TODO: update mint sequence number
				// TODO:
				// TODO:
				// TODO: upload to ipfs and update the mint request
				// TODO:
				_, err := cli.mintStore.UpdateStatusNFTMintRequest(ctx, fmt.Sprint(mr.NftMintRequest.Id), pb_nft.Status_Pending)
				if err != nil {
					log.Error().Msgf("MintStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
				// delete the mint request from the map on status update
				delete(cli.ttlMap, mr.NftMintRequest.Id)
			}

		}

	}
	return nil
}
