package quests

import (
	"fmt"

	"github.com/spf13/cobra"
)

// This function will return the test_create command
func NewTestCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "test_create",
		Short: "Experimental command to test interactive quest creation",
		RunE: func(cmd *cobra.Command, args []string) error {
			// This is where we'll integrate the survey logic later
			fmt.Println("test_create command executed!")
			return nil
		},
	}

	return cmd
}
