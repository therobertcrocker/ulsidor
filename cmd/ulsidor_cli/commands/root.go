package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/cmd/ulsidor_cli/commands/factions"
	"github.com/therobertcrocker/ulsidor/internal/core"
	"github.com/therobertcrocker/ulsidor/utils"
)

var coreInstance *core.Core

var rootCmd = &cobra.Command{
	Use:   "ulsidor",
	Short: "Ulsidor: A GM Tool",
	Long: ` Ulisdor is a tool for managing the TTRPGs I run. 
	   
	   It will eventually feature a number of sub-applications for managing things like Factions, Story Plots, and More.
	   
	   `,

	Run: func(cmd *cobra.Command, args []string) {
		utils.PrintTitle("GM Tools - Root Operation")
	},
}

func Execute(core *core.Core) {
	coreInstance = core
	factionscmd := factions.FactionsCmd(coreInstance)
	rootCmd.AddCommand(factionscmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
