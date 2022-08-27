package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/jekabolt/solutions-dapp/art-admin/app"
	"github.com/jekabolt/solutions-dapp/art-admin/app/auth"
	"github.com/jekabolt/solutions-dapp/art-admin/app/nft"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/eth"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/internal/store/bunt"
)

type Config struct {
	Bunt         *bunt.Config
	Bucket       *bucket.Config
	ETHWatcher   *eth.Config
	IPFS         *ipfs.Config
	Descriptions *descriptions.Config
	Server       *app.Config
	Nft          *nft.Config
	Auth         *auth.Config

	Debug bool `env:"DEBUG" envDefault:"false"`
}

func GetConfig() (*Config, error) {
	var err error

	cfg := &Config{
		Bunt:         &bunt.Config{},
		Bucket:       &bucket.Config{},
		ETHWatcher:   &eth.Config{},
		IPFS:         &ipfs.Config{},
		Descriptions: &descriptions.Config{},
		Server:       &app.Config{},
		Nft:          &nft.Config{},
		Auth:         &auth.Config{},
	}

	err = env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("GetConfig:env.Parse: [%s]", err.Error())
	}

	return cfg, nil
}
