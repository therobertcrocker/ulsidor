package storage

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/therobertcrocker/ulsidor/internal/interfaces"
)

type JSONStorageManager struct{}

// NewJSONStorageManager returns a new instance of JSONStorageManager
func NewJSONStorageManager() *JSONStorageManager {
	return &JSONStorageManager{}
}

// LoadRepo loads a repo from storage
func (jsm *JSONStorageManager) LoadRepo(id string, repo interfaces.Repository) error {
	filepath := filepath.Join("save_data/", id+".json")
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &repo); err != nil {
		return err
	}

	return nil

}

// SaveRepo saves a repo to storage
func (jsm *JSONStorageManager) SaveRepo(id string, repo interfaces.Repository) error {
	filepath := filepath.Join("save_data/", id+".json")
	data, err := json.Marshal(repo)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath, data, 0644); err != nil {
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
