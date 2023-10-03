package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/core"
)

var coreInstance *core.Core

var rootCmd = &cobra.Command{
	Use:   "ulsidor",
	Short: "Ulsidor: A GM Tool",
	Long: ` Ulisdor is a tool for managing the TTRPGs I run. 
	   
	   It will eventually feature a number of sub-applications for managing things like Factions, Story Plots, and More.
	   
	   `,

	Run: func(cmd *cobra.Command, args []string) {
		printTitle("GM Tools - Root Operation")
	},
}

func printTitle(text string) {
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("                 %s\n", text)
	fmt.Println("---------------------------------------------------------------------")
}

func Execute(core *core.Core) {
	coreInstance = core
	factionscmd := NewFactionsCmd()
	rootCmd.AddCommand(factionscmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}
