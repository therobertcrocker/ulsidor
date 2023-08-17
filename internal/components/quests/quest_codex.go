package quests

import "github.com/therobertcrocker/ulsidor/internal/data/game"

type QuestCodex struct {
	repo     QuestRepository
	gameData *game.GameData
}

func NewQuestCodex(repo QuestRepository, gd *game.GameData) *QuestCodex {
	return &QuestCodex{
		repo:     repo,
		gameData: gd,
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
