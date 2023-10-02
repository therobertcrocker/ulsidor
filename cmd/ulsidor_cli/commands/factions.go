package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewFactionsCmd() *cobra.Command {
	factionsCmd := &cobra.Command{
		Use:   "factions",
		Short: "Commands for managing factions",
	}

	createCmd := NewCreateFactionCmd()
	factionsCmd.AddCommand(createCmd)

	return factionsCmd

}

func NewCreateFactionCmd() *cobra.Command {
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new faction",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Creating a new faction")
		},
	}

	return createCmd
}
