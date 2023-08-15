package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/therobertcrocker/ulsidor/internal/config"
	"github.com/therobertcrocker/ulsidor/internal/interfaces"
)

type JSONStorageManager struct {
	storageConfig *config.Config
}

// NewJSONStorageManager returns a new instance of JSONStorageManager
func NewJSONStorageManager(conf *config.Config) *JSONStorageManager {
	return &JSONStorageManager{
		storageConfig: conf,
	}
}

// LoadRepo loads a repo from storage
func (jsm *JSONStorageManager) LoadRepo(id string, repo interfaces.Repository) error {
	baseDir := jsm.storageConfig.BasePath
	storagePath := filepath.Join(baseDir, "internal/storage/save_data/"+id+".json")
	data, err := os.ReadFile(storagePath)
	if err != nil {
		config.Log.Errorf("Failed to load repo from file: %v", err)
		return err
	}
	if err := json.Unmarshal(data, &repo); err != nil {
		config.Log.Errorf("Failed to unmarshall json: %v", err)
		return err
	}

	return nil

}

// SaveRepo saves a repo to storage
func (jsm *JSONStorageManager) SaveRepo(id string, repo interfaces.Repository) error {
	baseDir := jsm.storageConfig.BasePath
	storagePath := filepath.Join(baseDir, "internal/storage/save_data/"+id+".json")
	data, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	if err := os.WriteFile(storagePath, data, 0644); err != nil {
		return err
	}
	return nil
}

// LoadGameData loads game data from storage
func (jsm *JSONStorageManager) LoadGameData() (map[string]interface{}, error) {
	data, err := os.ReadFile("game_data/game_data.json")
	if err != nil {
		return nil, err
	}
	var gameData map[string]interface{}
	if err := json.Unmarshal(data, &gameData); err != nil {
		return nil, err
	}
	return gameData, nil

}
