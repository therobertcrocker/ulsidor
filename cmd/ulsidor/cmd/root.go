package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor/cmd/quests"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
)

var coreInstance *core.Core

var RootCmd = &cobra.Command{
	Use:   "ulsidor",
	Short: "Ulsidor: A GM Tool by Gestalt",
	// ... other properties ...
}

func Execute(core *core.Core) {
	coreInstance = core
	questsCmd := quests.NewQuestsCmd(coreInstance)
	RootCmd.AddCommand(questsCmd)

	if err := RootCmd.Execute(); err != nil {
		utils.Log.Error(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
