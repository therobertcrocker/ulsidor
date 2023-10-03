package main

import (
	"fmt"
	"os"

	cli "github.com/therobertcrocker/ulsidor/cmd/ulsidor_cli/commands"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/storage"
	"github.com/therobertcrocker/ulsidor/utils"
)

const (
	DATABASE_PATH = "data/ulsidor.db"
)

func main() {

	displayBanner()

	// TODO: configurations

	// initialize the storage manager
	utils.Log.WithField("database_path", DATABASE_PATH).Debug("initializing storage manager")
	storage, err := storage.NewBboltStorageManager(DATABASE_PATH)
	if err != nil {
		panic(err)
	}

	defer func() {
		storage.Close()

		utils.Log.Debug("---------- End of Run --------------")
	}()

	// initialize the core
	utils.Log.Debug("initializing core")
	core := core.NewCore(storage)
	err = core.Init()
	if err != nil {
		panic(err)
	}

	// start the CLI
	utils.Log.Debug("starting CLI")
	cli.Execute(core)
}

func displayBanner() {
	data, err := os.ReadFile("assets/banner.txt")
	if err != nil {
		utils.Log.WithError(err).Error("failed to read banner.txt")
	}
	fmt.Println(string(data))

}
