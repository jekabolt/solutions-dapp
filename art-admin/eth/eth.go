package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
	"github.com/nanmu42/etherscan-api"
	"github.com/rs/zerolog/log"
)

type Config struct {
	APIKey          string `env:"ETHERSCAN_API_KEY"`
	Network         string `env:"ETHERSCAN_NETWORK" envDefault:"api-rinkeby"`
	ContractAddress string `env:"CONTRACT_ADDRESS"`
	RefreshTime     string `env:"REFRESH_TIME" envDefault:"10m"`
}

type Etherscan struct {
	nftStore nft.Store
	*Config
	*etherscan.Client
	refreshTime time.Duration
	TTLMap      map[int]int // map[buntId]count
}

const (
	maxCountTTL = 10
)

func InitEtherscan(ctx context.Context, cfg *Config, nftStore nft.Store) (*Etherscan, error) {
	tf, err := time.ParseDuration(cfg.RefreshTime)
	if err != nil && cfg.RefreshTime != "" {
		return nil, fmt.Errorf("InitEtherscan:time.ParseDuration [%s]", err.Error())
	}
	eth := Etherscan{
		nftStore:    nftStore,
		Config:      cfg,
		Client:      etherscan.New(etherscan.Network(cfg.Network), cfg.APIKey),
		TTLMap:      make(map[int]int),
		refreshTime: tf,
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

func (eth *Etherscan) getTxStatus(txHash, address string) (nft.NFTStatus, error) {
	txs, err := eth.ERC721Transfers(&txHash, &address, &startBlock, &endBlock, page, offset, desc)
	if err != nil {
		return "", err
	}

	found := false
	for _, tx := range txs {
		if tx.ContractAddress == eth.ContractAddress &&
			tx.To == address {
			found = true

			if tx.Confirmations >= 10 {

				// TODO: check for double spend
				return nft.StatusCompleted, nil
			}

			if tx.Confirmations <= 10 {

				// TODO: check for double spend
				return nft.StatusPending, nil
			}

		}

		if !found {
			return nft.StatusFailed, nil
		}
	}

	return nft.StatusUnknown, nil

}

func (eth *Etherscan) TxStatusUpdate() error {
	mrs, err := eth.nftStore.GetAllNFTMintRequests()
	if err != nil {
		return err
	}

	for _, mr := range mrs {
		if mr.Status == nft.StatusUnknown ||
			mr.Status == nft.StatusPending {

			eth.TTLMap[mr.Id]++
			status, err := eth.getTxStatus(mr.TxHash, mr.ETHAddress)
			if err != nil {
				log.Error().Msgf("TxStatusUpdate:getTxStatus [%s]", err.Error())
			}

			if ok := mr.Status != status; ok {
				log.Debug().Msgf("TxStatusUpdate:update [%s]", status)
				_, err := eth.nftStore.UpdateStatusNFTMintRequest(&mr, nft.StatusFailed)
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
			}

			if eth.TTLMap[mr.Id] >= maxCountTTL {
				log.Debug().Msgf("TxStatusUpdate:delete [%s]", nft.StatusFailed)
				_, err := eth.nftStore.UpdateStatusNFTMintRequest(&mr, nft.StatusBad)
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
			}

		}

	}
	return nil
}
