package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db, err := cfg.Redis.InitDB(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init db err:[%s]", err.Error()))
	}

	desc, err := cfg.Descriptions.Init()

	b, err := cfg.Bucket.Init()
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init s3 bucket err:[%s]", err.Error()))
	}

	mrStore, err := db.MintRequestStore(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to init ETHWatcher err:[%s]", err.Error()))
	}
	eth, err := cfg.ETHWatcher.New(ctx, mrStore)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to start ETHWatcher err:[%s]", err.Error()))
	}
	eth.Run(ctx)

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

	app := cfg.Server.Init()
	err = app.Start(ctx, authS, nftS)
	if err != nil {
		log.Fatal().Err(err).Msg(fmt.Sprintf("Failed to start server:[%s]", err.Error()))
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	select {
	case s := <-sigChan:
		log.Warn().Msgf("signal received, exiting [%s]", s.String())
		app.Stop(ctx)
		eth.Stop()
		db.Close()
		log.Warn().Msg("application exited")
	case <-app.Done():
		log.Error().Msg("application exited")
	}
}
