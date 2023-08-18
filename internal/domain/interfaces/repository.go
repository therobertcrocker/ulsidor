package interfaces

// Repository is the interface for a repository

type Repository interface {
	Init() error
	Serialize() ([]byte, error)
	Deserialize(data []byte) error
}

// Metadata is the struct for metadata
type Metadata struct {
	Version     string `json:"version"`
	LastUpdated string `json:"last_updated"`
}
