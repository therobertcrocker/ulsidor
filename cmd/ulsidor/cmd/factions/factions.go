package factions

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
)

var coreInstance *core.Core

func NewFactionsCmd(core *core.Core) *cobra.Command {
	coreInstance = core

	factionsCmd := &cobra.Command{
		Use:   "factions",
		Short: "Commands for managing factions",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := coreInstance.InitFactionComponents()
			if err != nil {
				utils.Log.Error(err)
			}
		},
	}

	createCmd := NewCreateCmd()
	factionsCmd.AddCommand(createCmd)

	return factionsCmd
}
