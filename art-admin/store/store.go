package store

const (
	BuntDBType = "bunt"
)

type Store interface {
	NFTMintRequestCRUD
}

type NFTMintRequestCRUD interface {
	UpsertNFTMintRequest(p *NFTMintRequest) (*NFTMintRequest, error)
	GetAllNFTMintRequests() ([]*NFTMintRequest, error)
	DeleteNFTMintRequestById(id string) error
	UpdateStatusNFTMintRequest(p *NFTMintRequest, status NFTStatus) (*NFTMintRequest, error)
	UpsertNFT(p *NFTMintRequest) (*NFTMintRequest, error)
	DeleteNFT(id string) (*NFTMintRequest, error)
	GetAllToUpload() ([]*NFTMintRequest, error)
}
