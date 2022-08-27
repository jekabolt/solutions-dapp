package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jekabolt/solutions-dapp/art-admin/config"
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

	desc, err := cfg.Descriptions.Init()

	b, err := cfg.Bucket.Init()
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init s3 bucket err:[%s]", err.Error()))
	}

	eth, err := cfg.ETHWatcher.New(context.Background(), db.MintRequestStore())
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init ETHWatcher err:[%s]", err.Error()))
	}
	eth.Run(context.TODO())
	defer eth.Stop()

	ipfs, err := cfg.IPFS.Init(desc)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init ipfs uploader err:[%s]", err.Error()))
	}

	nftS, err := cfg.Nft.New(db, b, ipfs, desc)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to create new nft server:[%s]", err.Error()))
	}
	authS, err := cfg.Auth.New()
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to create new auth server:[%s]", err.Error()))
	}
	s := cfg.Server.Init()
	s.Start(context.TODO(), authS, nftS)

	c := make(chan struct{})
	<-c
}
