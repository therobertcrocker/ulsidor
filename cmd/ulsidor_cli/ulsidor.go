package main

import (
	cli "github.com/therobertcrocker/ulsidor/cmd/ulsidor_cli/commands"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/storage"
)

func main() {
	// TODO: configurations

	storage, err := storage.NewBboltStorageManager("ulsidor.db")
	if err != nil {
		panic(err)
	}

	defer storage.Close()

	core := core.NewCore(storage)

	// start the CLI
	cli.Execute(core)
}
