package quests

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/therobertcrocker/ulsidor/internal/domain/types"
)

// CollectQuestInput collects the input for a quest from the user
func CollectQuestInput(initialInput *types.CreateQuestInput) (*types.CreateQuestInput, error) {
	questInput := initialInput
	var err error

	// if the title is empty, prompt for it
	if questInput.Title == "" {
		questInput.Title, err = promptForQuestTitle()
		if err != nil {
			return nil, err
		}
	}

	// if the quest type is empty, prompt for it
	if questInput.QuestType == "" {
		questInput.QuestType, err = promptForQuestType()
		if err != nil {
			return nil, err
		}
	}

	// if the level is empty, prompt for it
	if questInput.Level == 0 {
		questInput.Level, err = promptForQuestLevel()
		if err != nil {
			return nil, err
		}
	}

	// prompt for the quest description
	questInput.Description, err = promptForQuestDescription()
	if err != nil {
		return nil, err
	}

	//prompt for the quest source
	questInput.Source, err = promptForQuestSource()
	if err != nil {
		return nil, err
	}

	return questInput, nil
}

// promptForQuestTitle prompts the user for the title of the quest
func promptForQuestTitle() (string, error) {
	var title string
	prompt := &survey.Input{
		Message: "What is the title of the quest?",
	}
	err := survey.AskOne(prompt, &title)

	return title, err
}

// promptForQuestType prompts the user for the type of the quest
func promptForQuestType() (string, error) {
	var questType string
	prompt := &survey.Select{
		Message: "What type of quest is it?",
		Options: []string{"Hunt", "Acquisition", "Whisper", "Knowledge"},
	}
	err := survey.AskOne(prompt, &questType)

	return questType, err
}

// promptForQuestLevel prompts the user for the level of the quest
func promptForQuestLevel() (int, error) {
	var level int
	var err error
	for level < 1 || level > 20 {
		prompt := &survey.Input{
			Message: "What level is the quest?",
		}
		err = survey.AskOne(prompt, &level)
	}

	return level, err
}

// promptForQuestDescription prompts the user for the description of the quest
func promptForQuestDescription() (string, error) {
	var description string
	prompt := &survey.Multiline{
		Message: "What is the description of the quest?",
	}
	err := survey.AskOne(prompt, &description)

	return description, err
}

// promptForQuestSource prompts the user for the source of the quest
func promptForQuestSource() (string, error) {
	var source string
	prompt := &survey.Input{
		Message: "Who is the source of the quest?",
	}
	err := survey.AskOne(prompt, &source)

	return source, err
}

// PromptForFieldsToEdit prompts the user for the fields to edit
func PromptForFieldsToEdit() ([]string, error) {
	var fieldsToEdit []string
	prompt := &survey.MultiSelect{
		Message: "What fields would you like to edit?",
		Options: []string{"Title", "Type", "Level", "Description", "Source"},
	}
	err := survey.AskOne(prompt, &fieldsToEdit)

	return fieldsToEdit, err
}

// EditSelectedFields edits the selected fields
func EditSelectedFields(initialInput *types.CreateQuestInput, fieldsToEdit []string) (*types.CreateQuestInput, error) {
	questInput := initialInput
	var err error
	for _, field := range fieldsToEdit {
		switch field {
		case "Title":
			questInput.Title, err = promptForQuestTitle()
		case "Type":
			questInput.QuestType, err = promptForQuestType()
		case "Level":
			questInput.Level, err = promptForQuestLevel()
		case "Description":
			questInput.Description, err = promptForQuestDescription()
		case "Source":
			questInput.Source, err = promptForQuestSource()
		}
		if err != nil {
			return nil, err
		}
	}

	return questInput, nil

}
