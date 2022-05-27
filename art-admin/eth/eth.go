package eth

import (
	"context"
	"fmt"
	"time"

	"github.com/jekabolt/solutions-dapp/art-admin/store"
	"github.com/nanmu42/etherscan-api"
	"github.com/rs/zerolog/log"
)

type EtherscanConfig struct {
	APIKey          string `env:"ETHERSCAN_API_KEY"`
	Network         string `env:"ETHERSCAN_NETWORK" envDefault:"api-rinkeby"`
	ContractAddress string `env:"CONTRACT_ADDRESS"`
	RefreshTime     string `env:"REFRESH_TIME" envDefault:"10m"`
}

type Etherscan struct {
	*EtherscanConfig
	*etherscan.Client
	refreshTime time.Duration
	TTLMap      map[int]int // map[buntId]count
}

const (
	maxCountTTL = 10
)

func InitEtherscan(ctx context.Context, cfg *EtherscanConfig, bunt store.Store) (*Etherscan, error) {
	tf, err := time.ParseDuration(cfg.RefreshTime)
	if err != nil && cfg.RefreshTime != "" {
		return nil, fmt.Errorf("InitEtherscan:time.ParseDuration [%s]", err.Error())
	}
	eth := Etherscan{
		EtherscanConfig: cfg,
		Client:          etherscan.New(etherscan.Network(cfg.Network), cfg.APIKey),
		TTLMap:          make(map[int]int),
		refreshTime:     tf,
	}
	eth.StartTxStatusUpdate(ctx, bunt)
	return &eth, nil
}

func (eth *Etherscan) StartTxStatusUpdate(ctx context.Context, bunt store.Store) {
	t := time.NewTicker(eth.refreshTime)
	go func() {
		for {
			select {
			case <-t.C:
				log.Debug().Msg("TxStatusUpdate: start")
				err := eth.TxStatusUpdate(bunt)
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

func (eth *Etherscan) getTxStatus(txHash, address string, bunt store.Store) (store.NFTStatus, error) {
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
				return store.StatusCompleted, nil
			}

			if tx.Confirmations <= 10 {

				// TODO: check for double spend
				return store.StatusPending, nil
			}

		}

		if !found {
			return store.StatusFailed, nil
		}
	}

	return store.StatusUnknown, nil

}

func (eth *Etherscan) TxStatusUpdate(bunt store.Store) error {
	mrs, err := bunt.GetAllNFTMintRequests()
	if err != nil {
		return err
	}

	for _, mr := range mrs {
		if mr.Status == store.StatusUnknown ||
			mr.Status == store.StatusPending {

			eth.TTLMap[mr.Id]++
			status, err := eth.getTxStatus(mr.TxHash, mr.ETHAddress, bunt)
			if err != nil {
				log.Error().Msgf("TxStatusUpdate:getTxStatus [%s]", err.Error())
			}

			if ok := mr.Status != status; ok {
				log.Debug().Msgf("TxStatusUpdate:update [%s]", status)
				_, err := bunt.UpdateStatusNFTMintRequest(mr, store.StatusFailed)
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
			}

			if eth.TTLMap[mr.Id] >= maxCountTTL {
				log.Debug().Msgf("TxStatusUpdate:delete [%s]", store.StatusFailed)
				_, err := bunt.UpdateStatusNFTMintRequest(mr, store.StatusBad)
				if err != nil {
					log.Error().Msgf("TxStatusUpdate:UpdateStatusNFTMintRequest [%s]", err.Error())
				}
			}

		}

	}
	return nil
}
