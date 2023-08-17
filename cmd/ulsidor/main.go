/*
Copyright © 2023 Robert Crocker Sir.Author.Doyle@Gmail.com
*/
package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd"
	"github.com/therobertcrocker/ulsidor/internal/config"
	"github.com/therobertcrocker/ulsidor/internal/core"
)

var mainConfig *config.Config

func main() {

	loadConfiguration()

	loadLogger()

	coreInstance := core.NewCore(mainConfig)
	cmd.Execute(coreInstance)
}

func loadConfiguration() {
	var err error
	mainConfig, err = config.LoadConfig("internal/config/config.json")

	if err != nil {
		panic(fmt.Errorf("fatal error loading game data: %s", err))
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

	config.InitLogger(level)
}
