package quests

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/types"
	prompts "github.com/therobertcrocker/ulsidor/internal/survey/quests"
)

var (
	titleFlag     string
	questTypeFlag string
	levelFlag     int
)

func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new quest",
		Run: func(cmd *cobra.Command, args []string) {

			// get the initial input from the flags
			initialInput := &types.CreateQuestInput{
				Title:     titleFlag,
				QuestType: questTypeFlag,
				Level:     levelFlag,
			}

			// collect the rest of the input from the user
			questInput, err := prompts.CollectQuestInput(initialInput)
			if err != nil {
				utils.Log.Error(err)
			}

			// confirm the input to the user, allow them to edit it
			err = confirmQuestDetails(questInput)
			if err != nil {
				utils.Log.Error(err)
			}

			// create the quest
			err = coreInstance.CreateNewQuest(questInput)
			if err != nil {
				utils.Log.Error(err)
			}

		},
	}

	cmd.Flags().StringVarP(&titleFlag, "title", "t", "", "Title of the quest")
	cmd.Flags().StringVarP(&questTypeFlag, "type", "y", "", "Type of the quest")
	cmd.Flags().IntVarP(&levelFlag, "level", "l", 0, "Level of the quest")

	return cmd
}

func displayQuestInput(questInput *types.CreateQuestInput) {
	fmt.Println("Here's the quest you've created:")
	fmt.Println("-------------------------------")
	fmt.Printf("Title: %s\n", questInput.Title)
	fmt.Printf("Type: %s\n", questInput.QuestType)
	fmt.Printf("Level: %d\n", questInput.Level)
	fmt.Printf("Description: %s\n", questInput.Description)
	fmt.Printf("Source: %s\n", questInput.Source)
	fmt.Println("-------------------------------")
}

func confirmQuestDetails(questInput *types.CreateQuestInput) error {
	var isConfirmed bool
	for !isConfirmed {
		displayQuestInput(questInput)
		fieldsToEdit, err := prompts.PromptForFieldsToEdit()
		if err != nil {
			return err
		}
		if len(fieldsToEdit) == 0 {
			isConfirmed = true
		}
		prompts.EditSelectedFields(questInput, fieldsToEdit)
	}

	return nil
}
