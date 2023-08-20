package config

type Config struct {
	LogLevel          string `json:"log_level"`
	EntityStoragePath string `json:"entity_storage_path"`
	BBoltStoragePath  string `json:"bbolt_storage_path"`
}
