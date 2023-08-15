package config

type Config struct {
	LogLevel string    `json:"log_level"`
	BasePath string    `json:"base_path"`
	GameData *GameData `json:"game_data"`
}

type GameData struct {
	PartyLevel int `json:"party_level"`
}
