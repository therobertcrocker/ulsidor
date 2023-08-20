/*
Copyright © 2023 Robert Crocker Sir.Author.Doyle@Gmail.com
*/
package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/data/config"
	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/data/storage"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
)

var mainConfig *config.Config
var storageManager *storage.BBoltStorageManager
var gameData *game.GameData

func main() {

	loadConfiguration()
	loadLogger()

	initStorage()
	defer storageManager.Close()
	loadGameData()

	coreInstance := core.NewCore(mainConfig, storageManager, gameData)
	cmd.Execute(coreInstance)
}

func loadConfiguration() {

	err := utils.LoadJSONData("internal/data/config/config.json", &mainConfig)
	if err != nil {
		utils.Log.Fatalf("Failed to load configuration: %v", err)
		os.Exit(1)
	}
}

func initStorage() {
	var err error
	storageManager, err = storage.NewBBoltStorageManager(mainConfig.BBoltStoragePath)
	if err != nil {
		utils.Log.Fatalf("Failed to initialize storage: %v", err)
		os.Exit(1)
	}
}

func loadGameData() {
	err := utils.LoadJSONData("internal/data/game/game_data.json", &gameData)
	if err != nil {
		utils.Log.Fatalf("Failed to load game data: %v", err)
		os.Exit(1)
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
		fmt.Printf("Invalid log level: %s\n", logLevel)
		os.Exit(1)
	}

	utils.InitLogger(level)
}
