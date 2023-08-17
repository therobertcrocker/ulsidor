package quests

import (
	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/config"
)

func NewCreateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new quest",
		Run: func(cmd *cobra.Command, args []string) {
			// Extract necessary arguments from args or flags
			title := "test"
			questType := "test"
			description := "test"
			source := "test"
			level := 1

			_, err := coreInstance.CreateNewQuest(title, questType, description, source, level)
			if err != nil {
				// Handle error
				config.Log.Errorf("Failed to create new quest: %v", err)
			}

		},
	}
}
