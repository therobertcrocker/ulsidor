package ui

import (
	"fmt"
	"strconv"

	"github.com/AlecAivazis/survey/v2"
	"github.com/therobertcrocker/ulsidor/internal/machinations"
)

var (
	traitValues = map[string]map[string][]string{
		"small": {
			"primary":   {"3", "4"},
			"secondary": {"2", "3"},
			"tertiary":  {"1", "2"},
		},
		"medium": {
			"primary":   {"5", "6"},
			"secondary": {"4", "5"},
			"tertiary":  {"2", "3"},
		},
		"large": {
			"primary":   {"7", "8"},
			"secondary": {"6", "7"},
			"tertiary":  {"3", "4"},
		},
	}
)

func CollectFactionInput() (*machinations.FactionInput, error) {
	factionInput := &machinations.FactionInput{}

	// prompt for the faction name
	namePrompt := &survey.Input{
		Message: "What is the name of the faction?",
	}
	if err := survey.AskOne(namePrompt, &factionInput.Name); err != nil {
		return nil, fmt.Errorf("failed to collect faction name: %w", err)
	}

	// prompt for the faction size
	sizePrompt := &survey.Select{
		Message: "What is the size of the faction?",
		Options: []string{"small", "medium", "large"},
	}
	if err := survey.AskOne(sizePrompt, &factionInput.Size); err != nil {
		return nil, fmt.Errorf("failed to collect faction size: %w", err)
	}

	traitLabels := []string{"primary", "secondary", "tertiary"}
	for _, trait := range traitLabels {
		if err := collectTraits(factionInput, trait); err != nil {
			return nil, fmt.Errorf("failed to collect faction traits: %w", err)
		}
	}

	// prompt for the faction arcane
	arcanePrompt := &survey.Select{
		Message: "What is the arcane level of the faction?",
		Options: []string{"none", "low", "medium", "high"},
	}
	if err := survey.AskOne(arcanePrompt, &factionInput.Arcane); err != nil {
		return nil, fmt.Errorf("failed to collect faction arcane: %w", err)
	}

	return factionInput, nil

}

func collectTraits(factionInput *machinations.FactionInput, trait string) error {
	var traitLabel string
	var traitValue string

	// prompt for primary trait
	traitPrompt := &survey.Select{
		Message: fmt.Sprintf("What is the %s trait of the faction?", trait),
		Options: []string{"cunning", "force", "wealth"},
	}
	if err := survey.AskOne(traitPrompt, &traitLabel); err != nil {
		return fmt.Errorf("failed to collect faction %s trait: %w", trait, err)
	}

	// prompt for primary trait value
	valuePrompt := &survey.Select{
		Message: "select a value for this trait",
		Options: traitValues[factionInput.Size][trait],
	}
	if err := survey.AskOne(valuePrompt, &traitValue); err != nil {
		return fmt.Errorf("failed to collect faction %s trait value: %w", trait, err)
	}

	// convert the trait value to an int
	traitValueInt, err := strconv.Atoi(traitValue)
	if err != nil {
		return fmt.Errorf("failed to convert trait value to int: %w", err)
	}

	// set the primary trait
	switch traitLabel {
	case "cunning":
		factionInput.Traits.Cunning = traitValueInt
	case "force":
		factionInput.Traits.Force = traitValueInt
	case "wealth":
		factionInput.Traits.Wealth = traitValueInt
	}

	return nil

}
