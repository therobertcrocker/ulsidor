// internal/config/loader.go

package config

import (
	"encoding/json"
	"os"
)

func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		Log.Errorf("Failed to load config from file: %v", err)
		return nil, err
	}
	defer file.Close()

	config := &Config{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		Log.Errorf("Failed to decode config: %v", err)
		return nil, err
	}

	return config, nil
}
