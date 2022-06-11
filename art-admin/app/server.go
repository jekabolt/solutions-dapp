package app

import (
	"github.com/jekabolt/solutions-dapp/art-admin/auth"
	"github.com/jekabolt/solutions-dapp/art-admin/bucket"
	"github.com/jekabolt/solutions-dapp/art-admin/config"
	"github.com/jekabolt/solutions-dapp/art-admin/descriptions"
	"github.com/jekabolt/solutions-dapp/art-admin/eth"
	"github.com/jekabolt/solutions-dapp/art-admin/ipfs"
	"github.com/jekabolt/solutions-dapp/art-admin/store/bunt"
	"github.com/jekabolt/solutions-dapp/art-admin/store/metadata"
	"github.com/jekabolt/solutions-dapp/art-admin/store/nft"
)

type Server struct {
	nftStore      nft.Store
	metadataStore metadata.Store
	Bucket        *bucket.Bucket
	Auth          *auth.Auth
	Config        *config.Config
	eth           *eth.Etherscan
	ipfs          *ipfs.Moralis
	descs         *descriptions.Store
}

func InitServer(
	db *bunt.BuntDB,
	bucket *bucket.Bucket,
	cfg *config.Config,
	eth *eth.Etherscan,
	ipfs *ipfs.Moralis,
	descs *descriptions.Store,
) *Server {
	a := cfg.Auth.New()
	return &Server{
		nftStore:      db.NFTStore(),
		metadataStore: db.MetadataStore(),
		Bucket:        bucket,
		Auth:          a,
		Config:        cfg,
		eth:           eth,
		ipfs:          ipfs,
		descs:         descs,
	}
}
