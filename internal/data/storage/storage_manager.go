package storage

import (
	"os"
	"path/filepath"

	"github.com/therobertcrocker/ulsidor/internal/interfaces"
)

type JSONStorageManager struct {
}

// NewJSONStorageManager returns a new instance of JSONStorageManager
func NewJSONStorageManager() *JSONStorageManager {
	return &JSONStorageManager{}
}

// LoadRepo loads data from storage into a repository.
func (jsm *JSONStorageManager) LoadRepo(id string, repo interfaces.Repository) error {
	storagePath := filepath.Join("internal/storage/save_data/" + id + ".json")
	fileData, err := os.ReadFile(storagePath)
	if err != nil {
		return err
	}
	return repo.Deserialize(fileData)
}

// SaveRepo saves data from a repository to storage.
func (jsm *JSONStorageManager) SaveRepo(id string, repo interfaces.Repository) error {
	storagePath := filepath.Join("internal/storage/save_data/" + id + ".json")

	fileData, err := repo.Serialize()
	if err != nil {
		return err
	}

	if err := os.WriteFile(storagePath, fileData, 0644); err != nil {
		return err
	}
	return nil
}
