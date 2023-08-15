package quests

import "github.com/therobertcrocker/ulsidor/internal/config"

type QuestCodex struct {
	gameData *config.GameData
	repo     QuestRepository
}

func NewQuestCodex(repo QuestRepository, data *config.GameData) *QuestCodex {
	return &QuestCodex{
		repo:     repo,
		gameData: data,
	}
}

// CreateNewQuest creates a new quest.
func (qc *QuestCodex) CreateNewQuest(title, questType, description, source string, level int) (*Quest, error) {
	//builds the quest object and then adds it to the repository

	quest := &Quest{
		Title:       title,
		QuestType:   questType,
		Description: description,
		Source:      source,
		Level:       level,
	}

	quest.CalculateRewards(qc.gameData.PartyLevel)
	err := qc.repo.AddNewQuest(quest)
	if err != nil {
		return nil, err
	}

	return quest, nil

}
