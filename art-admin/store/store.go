package store

const (
	BuntDBType = "bunt"
)

// type Store interface {
// 	NFTStore
// 	MetadataStore
// }

// type NFTStore interface {
// 	UpsertNFTMintRequest(p *NFTMintRequest) (*NFTMintRequest, error)
// 	GetAllNFTMintRequests() ([]*NFTMintRequest, error)
// 	DeleteNFTMintRequestById(id string) error
// 	UpdateStatusNFTMintRequest(p *NFTMintRequest, status NFTStatus) (*NFTMintRequest, error)
// 	UpsertNFT(p *NFTMintRequest) (*NFTMintRequest, error)
// 	DeleteNFT(id string) (*NFTMintRequest, error)
// 	GetAllToUpload() ([]*NFTMintRequest, error)
// }

// type MetadataStore interface {
// 	AddOffchainMetadata(url string) error
// 	GetAllOffchainMetadata() ([]string, error)
// }
