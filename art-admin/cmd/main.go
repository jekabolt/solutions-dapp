package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/app"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/config"
	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/eth"
	"github.com/jekabolt/solutions-dapp/art-admin/ipfs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse env variables")
	}

	if cfg.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	confBytes, _ := json.Marshal(cfg)
	log.Info().Str("config:", "").Msg(string(confBytes))

	db, err := cfg.Bunt.InitDB()
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init db err:[%s]", err.Error()))
	}

	desc, err := descriptions.Init(cfg.Descriptions)

	b, err := bucket.InitBucket(cfg.Bucket)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init s3 bucket err:[%s]", err.Error()))
	}

	eth, err := eth.InitEtherscan(context.Background(), cfg.Etherscan, db)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init etherscan err:[%s]", err.Error()))
	}

	ipfs, err := ipfs.InitMoralis(cfg.IPFS, desc)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init etherscan err:[%s]", err.Error()))
	}

	s := app.InitServer(db, b, cfg, eth, ipfs, desc)

	log.Fatal().Err(s.Serve()).Msg("InitServer")

}
