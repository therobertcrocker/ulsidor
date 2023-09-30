package factions

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/factions"
	prompts "github.com/therobertcrocker/ulsidor/internal/survey/factions"
)

var (
	nameFlag   string
	sizeFlag   string
	arcaneFlag string
)

func NewCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new faction",
		Run: func(cmd *cobra.Command, args []string) {

			// get the initial input from the flags
			initialInput := &types.CreateFactionInput{
				Name:   nameFlag,
				Size:   sizeFlag,
				Arcane: arcaneFlag,
			}

			// collect the rest of the input from the user
			factionInput, err := prompts.CollectFactionInput(initialInput)
			if err != nil {
				utils.Log.Error(err)
			}

			// confirm the input to the user, allow them to edit it
			err = confirmFactionDetails(factionInput)
			if err != nil {
				utils.Log.Error(err)
			}

			// create the faction
			err = coreInstance.CreateNewFaction(factionInput)
			if err != nil {
				utils.Log.Error(err)
			}

		},
	}

	cmd.Flags().StringVarP(&nameFlag, "name", "n", "", "Name of the faction")
	cmd.Flags().StringVarP(&sizeFlag, "size", "s", "", "Size of the faction")
	cmd.Flags().StringVarP(&arcaneFlag, "arcane", "a", "", "Arcane of the faction")

	return cmd
}

func displayFactionInput(factionInput *types.CreateFactionInput) {
	fmt.Println("Here's the faction you've created:")
	fmt.Println("-------------------------------")
	fmt.Printf("Name: %s\n", factionInput.Name)
	fmt.Printf("Size: %s\n", factionInput.Size)
	fmt.Printf("Arcane: %s\n", factionInput.Arcane)
	fmt.Printf("Attributes: %v\n", factionInput.Attributes)
	fmt.Println("-------------------------------")
}

func confirmFactionDetails(factionInput *types.CreateFactionInput) error {
	var isConfirmed bool
	for !isConfirmed {
		displayFactionInput(factionInput)
		fieldsToEdit, err := prompts.PromptForFieldsToEdit()
		if err != nil {
			return err
		}
		if len(fieldsToEdit) == 0 {
			isConfirmed = true
		}
		prompts.EditSelectedFields(factionInput, fieldsToEdit)
	}

	return nil
}
