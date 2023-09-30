package factions

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/therobertcrocker/ulsidor/internal/domain/types/factions"
)

// CollectFactionInput collects the input for a faction from the user
func CollectFactionInput(initialInput *factions.CreateFactionInput) (*factions.CreateFactionInput, error) {
	factionInput := initialInput
	var err error

	// if the name is empty, prompt for it
	if factionInput.Name == "" {
		factionInput.Name, err = promptForFactionName()
		if err != nil {
			return nil, err
		}
	}

	// if the size is empty, prompt for it
	if factionInput.Size == "" {
		factionInput.Size, err = promptForFactionSize()
		if err != nil {
			return nil, err
		}
	}

	// if the arcane is empty, prompt for it
	if factionInput.Arcane == "" {
		factionInput.Arcane, err = promptForFactionArcane()
		if err != nil {
			return nil, err
		}
	}

	// prompt for each attribute
	factionInput.Attributes, err = promptForFactionAttributes()

	return factionInput, err

}

// promptForFactionName prompts the user for the name of the faction
func promptForFactionName() (string, error) {
	var name string
	prompt := &survey.Input{
		Message: "What is the name of the faction?",
	}
	err := survey.AskOne(prompt, &name)

	return name, err
}

// promptForFactionSize prompts the user for the size of the faction
func promptForFactionSize() (string, error) {
	var size string
	prompt := &survey.Select{
		Message: "What is the size of the faction?",
		Options: []string{"Small", "Medium", "Large"},
	}
	err := survey.AskOne(prompt, &size)

	return size, err

}

// promptForFactionArcane prompts the user for the arcane of the faction
func promptForFactionArcane() (string, error) {
	var arcane string
	prompt := &survey.Select{
		Message: "What is the arcane of the faction?",
		Options: []string{"None", "Low", "Medium", "High"},
	}
	err := survey.AskOne(prompt, &arcane)

	return arcane, err

}

// promptForFactionAttributes prompts the user for the attributes of the faction
func promptForFactionAttributes() ([]string, error) {
	var attributes []string
	var primary, secondary, tertiary string

	// prompt for the primary attribute
	primary, err := promptForAttribute("primary")
	if err != nil {
		return nil, err
	}

	// prompt for the secondary attribute
	secondary, err = promptForAttribute("secondary")
	if err != nil {
		return nil, err
	}

	// prompt for the tertiary attribute
	tertiary, err = promptForAttribute("tertiary")
	if err != nil {
		return nil, err
	}

	// add the attributes to the slice
	attributes = append(attributes, primary, secondary, tertiary)

	return attributes, nil

}

func promptForAttribute(tier string) (string, error) {
	var attribute string
	message := fmt.Sprintf("What is the %s attribute of the faction?", tier)
	prompt := &survey.Select{
		Message: message,
		Options: []string{"Cunning", "Force", "Wealth"},
	}
	err := survey.AskOne(prompt, &attribute)

	return attribute, err
}

// PromptForFieldsToEdit prompts the user for the fields to edit
func PromptForFieldsToEdit() ([]string, error) {
	var fieldsToEdit []string
	prompt := &survey.MultiSelect{
		Message: "What fields would you like to edit?",
		Options: []string{"Name", "Size", "Arcane", "Attributes"},
	}
	err := survey.AskOne(prompt, &fieldsToEdit)

	return fieldsToEdit, err
}

// EditSelectedFields edits the selected fields
func EditSelectedFields(initialInput *factions.CreateFactionInput, fieldsToEdit []string) (*factions.CreateFactionInput, error) {
	factionInput := initialInput
	var err error
	for _, field := range fieldsToEdit {
		switch field {
		case "Name":
			factionInput.Name, err = promptForFactionName()
		case "Size":
			factionInput.Size, err = promptForFactionSize()
		case "Arcane":
			factionInput.Arcane, err = promptForFactionArcane()
		case "Attributes":
			factionInput.Attributes, err = promptForFactionAttributes()
		}
		if err != nil {
			return nil, err
		}
	}

	return factionInput, nil
}
