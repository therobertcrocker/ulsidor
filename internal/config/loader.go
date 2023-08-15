// internal/config/loader.go

package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(filepath string) (*GameData, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	config := &GameData{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
