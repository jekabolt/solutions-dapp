package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/jekabolt/solutions-dapp/art-admin/auth"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/eth"
	"github.com/jekabolt/solutions-dapp/art-admin/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/store/bunt"
)

type Config struct {
	Port  string   `env:"PORT" envDefault:"8081"`
	Hosts []string `env:"HOSTS" envSeparator:"|"`

	Bunt         *bunt.Config
	Auth         *auth.Config
	Bucket       *bucket.Config
	Etherscan    *eth.Config
	IPFS         *ipfs.Config
	Descriptions *descriptions.Config

	Debug          bool `env:"DEBUG" envDefault:"false"`
	NFTTotalSupply int  `env:"NFT_TOTAL_SUPPLY" envDefault:"1000"`
}

func GetConfig() (*Config, error) {
	var err error

	cfg := &Config{
		Auth:         &auth.Config{},
		Bunt:         &bunt.Config{},
		Bucket:       &bucket.Config{},
		Etherscan:    &eth.Config{},
		IPFS:         &ipfs.Config{},
		Descriptions: &descriptions.Config{},
	}

	err = env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("GetConfig:env.Parse: [%s]", err.Error())
	}

	return cfg, nil
}
