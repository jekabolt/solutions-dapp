package app

import (
	"github.com/jekabolt/solutions-dapp/art-admin/auth"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/config"
	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/eth"
	"github.com/jekabolt/solutions-dapp/art-admin/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/store"
)

type Server struct {
	DB     store.Store
	Bucket *bucket.Bucket
	Auth   *auth.Auth
	Config *config.Config
	eth    *eth.Etherscan
	ipfs   *ipfs.Moralis
	descs  *descriptions.Store
}

func InitServer(
	db store.Store,
	bucket *bucket.Bucket,
	cfg *config.Config,
	eth *eth.Etherscan,
	ipfs *ipfs.Moralis,
	descs *descriptions.Store,
) *Server {
	a := cfg.Auth.New()
	return &Server{
		DB:     db,
		Bucket: bucket,
		Auth:   a,
		Config: cfg,
		eth:    eth,
		ipfs:   ipfs,
		descs:  descs,
	}
}
