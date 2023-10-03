package factions

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor_cli/commands/factions/assets"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/internal/ui"
	"github.com/therobertcrocker/ulsidor/utils"
)

var (
	coreInstance *core.Core
)

func FactionsCmd(core *core.Core) *cobra.Command {
	coreInstance = core
	factionsCmd := &cobra.Command{
		Use:   "factions",
		Short: "Commands for managing factions",
		Run: func(cmd *cobra.Command, args []string) {
			utils.PrintTitle("Machinations - Faction Management")
		},
	}

	AddCmd := AddFactionCmd()
	AssetsCmd := assets.AssetsCmd(coreInstance)
	factionsCmd.AddCommand(AddCmd)
	factionsCmd.AddCommand(AssetsCmd)

	return factionsCmd

}

/*
===============================================================================

	FACTION COMMANDs

===============================================================================
*/
func AddFactionCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new faction",
		Run: func(cmd *cobra.Command, args []string) {
			utils.PrintTitle("Machinations - Create Faction")

			factionInput, err := ui.CollectFactionInput()
			if err != nil {
				utils.Log.WithError(err).Error("failed to collect faction input")
			}

			if err := coreInstance.Factions.Codex.AddFaction(factionInput); err != nil {
				utils.Log.WithError(err).Error("failed to add faction")
			}

		},
	}

	return addCmd
}
