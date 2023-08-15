/*
Copyright © 2023 Robert Crocker Sir.Author.Doyle@Gmail.com
*/
package main

import (
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd"
	_ "github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd/quests"
	"github.com/therobertcrocker/ulsidor/internal/core"
)

func main() {
	coreInstance := core.NewCore()
	cmd.Execute(coreInstance)
}
