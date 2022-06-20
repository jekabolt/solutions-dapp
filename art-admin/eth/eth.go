package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/store/bunt"
	etherscan "github.com/nanmu42/etherscan-api"
	"github.com/rs/zerolog/log"
)

type Config struct {
	APIKey          string `env:"ETHERSCAN_API_KEY"`
	Network         string `env:"ETHERSCAN_NETWORK" envDefault:"api-rinkeby"`
	ContractAddress string `env:"CONTRACT_ADDRESS"`
	RefreshTime     string `env:"REFRESH_TIME" envDefault:"10m"`
}

type ETHCli interface {
	ERC721Transfers(contractAddress, address *string, startBlock *int, endBlock *int, page int, offset int, desc bool) (txs []etherscan.ERC721Transfer, err error)
}

type Etherscan struct {
	mintRequestStore bunt.MintRequestStore
	*Config
	ethCli      ETHCli
	refreshTime time.Duration
	TTLMap      map[int]int // map[buntId]count
}

const (
	maxCountTTL = 10
)

func (c *Config) InitEtherscan(ctx context.Context, mintRequestStore bunt.MintRequestStore) (*Etherscan, error) {
	tf, err := time.ParseDuration(c.RefreshTime)
	if err != nil && c.RefreshTime != "" {
		return nil, fmt.Errorf("InitEtherscan:time.ParseDuration [%s]", err.Error())
	}
	eth := Etherscan{
		mintRequestStore: mintRequestStore,
		Config:           c,
		ethCli:           etherscan.New(etherscan.Network(c.Network), c.APIKey),
		TTLMap:           make(map[int]int),
		refreshTime:      tf,
	}
	eth.StartTxStatusUpdate(ctx)
	return &eth, nil
}

func (eth *Etherscan) StartTxStatusUpdate(ctx context.Context) {
	t := time.NewTicker(eth.refreshTime)
	go func() {
		for {
			select {
			case <-t.C:
				log.Debug().Msg("TxStatusUpdate: start")
				err := eth.TxStatusUpdate()
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:TxStatusUpdate [%s]", err.Error())
				}
			case <-ctx.Done():
				return
			}
		}
	}()
}

var (
	startBlock = 0
	endBlock   = 9999999999
	page       = 1
	offset     = 100
	desc       = false
)

func (e *Etherscan) getTxStatus(txHash, address string) (bunt.NFTStatus, error) {
	txs, err := e.ethCli.ERC721Transfers(&txHash, &address, &startBlock, &endBlock, page, offset, desc)
	if err != nil {
		return "", err
	}

	found := false
	for _, tx := range txs {
		if tx.ContractAddress == e.ContractAddress &&
			tx.To == address {
			found = true

			if tx.Confirmations >= 10 {

				// TODO: check for double spend
				return bunt.StatusCompleted, nil
			}

			if tx.Confirmations <= 10 {

				// TODO: check for double spend
				return bunt.StatusPending, nil
			}

		}

		if !found {
			return bunt.StatusFailed, nil
		}
	}

	return bunt.StatusUnknown, nil

}

func (eth *Etherscan) TxStatusUpdate() error {
	mrs, err := eth.mintRequestStore.GetAllNFTMintRequests()
	if err != nil {
		return err
	}

	for _, mr := range mrs {
		if mr.Status == bunt.StatusUnknown.String() ||
			mr.Status == bunt.StatusPending.String() {

			eth.TTLMap[int(mr.NftMintRequest.Id)]++
			status, err := eth.getTxStatus(mr.NftMintRequest.TxHash, mr.NftMintRequest.EthAddress)
			if err != nil {
				log.Error().Msgf("TxStatusUpdate:getTxStatus [%s]", err.Error())
			}

			if ok := mr.Status != status.String(); ok {
				log.Debug().Msgf("TxStatusUpdate:update [%s]", status)
				_, err := eth.mintRequestStore.UpdateStatusNFTMintRequest(fmt.Sprint(mr.NftMintRequest.Id), bunt.StatusFailed)
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
			}

			if eth.TTLMap[int(mr.NftMintRequest.Id)] >= maxCountTTL {
				log.Debug().Msgf("TxStatusUpdate:delete [%s]", bunt.StatusFailed)
				_, err := eth.mintRequestStore.UpdateStatusNFTMintRequest(fmt.Sprint(mr.NftMintRequest.Id), bunt.StatusBad)
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
			}

		}

	}
	return nil
}
