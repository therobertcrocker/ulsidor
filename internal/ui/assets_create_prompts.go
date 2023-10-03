package ui

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/therobertcrocker/ulsidor/internal/machinations"
	"github.com/therobertcrocker/ulsidor/utils"
)

var (
	traits      = []string{"cunning", "force", "wealth"}
	basePrompts = []*survey.Question{
		{
			Name: "description",
			Prompt: &survey.Multiline{
				Message: "What is the description of the asset?",
			},
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What is the type of the asset?",
				Options: traits,
			},
		},
		{
			Name: "cost",
			Prompt: &survey.Input{
				Message: "What is the cost of the asset?",
			},
			Validate: survey.Required,
		},
		{
			Name: "upkeep",
			Prompt: &survey.Input{
				Message: "What is the upkeep of the asset?",
			},
			Validate: survey.Required,
		},
		{
			Name: "threshold",
			Prompt: &survey.Input{
				Message: "What is the purchase threshold of the asset?",
			},
			Validate: survey.Required,
		},
		{
			Name: "hitPoints",
			Prompt: &survey.Input{
				Message: "How many hit points does the asset have?",
			},
		},
		{
			Name: "arcane",
			Prompt: &survey.Select{
				Message: "What is the arcane level of the asset?",
				Options: []string{"none", "low", "medium", "high"},
			},
		},
		{
			Name: "qualities",
			Prompt: &survey.MultiSelect{
				Message: "What qualities does the asset have?",
				Options: []string{"subtle", "stealth", "special", "action"},
			},
		},
	}
	attackPrompts = []*survey.Question{
		{
			Name: "attackerTrait",
			Prompt: &survey.Select{
				Message: "What trait does the asset use to attack?",
				Options: traits,
			},
		},
		{
			Name: "defenderTrait",
			Prompt: &survey.Select{
				Message: "What trait does the defender use?",
				Options: traits,
			},
		},
	}
	damagePrompts = []*survey.Question{
		{
			Name: "count",
			Prompt: &survey.Input{
				Message: "How many dice does the asset roll for damage?",
			},
		},
		{
			Name: "dice",
			Prompt: &survey.Input{
				Message: "How many sides does the asset's damage dice have?",
			},
		},
		{
			Name: "modifier",
			Prompt: &survey.Input{
				Message: "What is the asset's damage modifier?",
			},
		},
	}

	// keep asking

)

func CollectAssetInput() (*machinations.AssetInput, error) {

	// declare the input structs
	assetInput := machinations.AssetInput{}
	attackInput := machinations.AttackInput{}
	dmgInput := machinations.Damage{}
	counterInput := machinations.Damage{}

	// prompt for the asset name
	var rawName string
	namePrompt := &survey.Input{
		Message: "What is the name of the asset?",
	}
	if err := survey.AskOne(namePrompt, &rawName); err != nil {
		return nil, err
	}

	assetInput.Name = utils.Normalize(rawName)

	// prompt for the base asset information
	if err := survey.Ask(basePrompts, &assetInput); err != nil {
		return nil, err
	}

	// prompt for the asset attack
	attackConfirm := &survey.Confirm{
		Message: "Does the asset have an attack?",
	}
	if err := survey.AskOne(attackConfirm, &assetInput.HasAttack); err != nil {
		return nil, err
	}

	// if the asset has an attack, prompt for the attack information
	if assetInput.HasAttack {
		if err := survey.Ask(attackPrompts, &attackInput); err != nil {
			return nil, err
		}
		if err := survey.Ask(damagePrompts, &dmgInput); err != nil {
			return nil, err
		}
		assetInput.Attack = &machinations.Attack{
			AttackerTrait: attackInput.AttackerTrait,
			DefenderTrait: attackInput.DefenderTrait,
			Damage: machinations.Damage{
				Count:    dmgInput.Count,
				Dice:     dmgInput.Dice,
				Modifier: dmgInput.Modifier,
			},
		}
	}

	// prompt for the asset counter
	counterConfirm := &survey.Confirm{
		Message: "Does the asset have a counter?",
	}
	if err := survey.AskOne(counterConfirm, &assetInput.HasCounter); err != nil {
		return nil, err
	}

	// if the asset has a counter, prompt for the counter information
	if assetInput.HasCounter {
		if err := survey.Ask(damagePrompts, &counterInput); err != nil {
			return nil, err
		}
		assetInput.Counter = &machinations.Damage{
			Count:    counterInput.Count,
			Dice:     counterInput.Dice,
			Modifier: counterInput.Modifier,
		}
	}

	return &assetInput, nil

}
