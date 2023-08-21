package interfaces

import (
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces/changelog"
)

// Repository is the interface for a repository

type Repository interface {
	Init() error
	LogChange(entry changelog.LogEntry) error
	Serialize() (map[string][]byte, error)
	Deserialize(collection map[string][]byte) error
}

// Metadata is the struct for metadata
type Metadata struct {
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}
