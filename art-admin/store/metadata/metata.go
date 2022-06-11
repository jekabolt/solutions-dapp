package metadata

type Store interface {
	AddOffchainMetadata(url string) error
	GetAllOffchainMetadata() ([]string, error)
}
