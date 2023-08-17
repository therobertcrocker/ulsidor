/*
Copyright © 2023 Robert Crocker Sir.Author.Doyle@Gmail.com
*/
package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/data/config"
	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/data/storage"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
)

var mainConfig *config.Config
var storageManager *storage.JSONStorageManager
var gameData *game.GameData

func main() {

	loadConfiguration()
	loadLogger()

	initStorage()
	loadGameData()

	coreInstance := core.NewCore(mainConfig, storageManager, gameData)
	cmd.Execute(coreInstance)
}

func loadConfiguration() {

	err := utils.LoadJSONData("internal/data/config/config.json", &mainConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to load configuration: %v", err))
	}
}

func initStorage() {
	storageManager = storage.NewJSONStorageManager(mainConfig)
}

func loadGameData() {
	err := utils.LoadJSONData("internal/data/game/game_data.json", &gameData)
	if err != nil {
		panic(fmt.Sprintf("Failed to load game data: %v", err))
	}
}

func loadLogger() {
	// Assuming gameData is the loaded configuration
	logLevel := mainConfig.LogLevel
	if logLevel == "" {
		logLevel = "info" // Default log level
	}

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		panic(fmt.Sprintf("Invalid log level: %v", err))
	}

	utils.InitLogger(level)
}
