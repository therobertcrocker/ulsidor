package quests

import (
	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/domain/types"
)

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
func (qc *QuestCodex) CreateNewQuest(questInput *types.CreateQuestInput) (*Quest, error) {
	//builds the quest object and then adds it to the repository

	quest := &Quest{
		ID:          questInput.Title,
		Title:       questInput.Title,
		QuestType:   questInput.QuestType,
		Description: questInput.Description,
		Source:      questInput.Source,
		Level:       questInput.Level,
	}

	quest.CalculateRewards(qc.gameData.PartyLevel)
	err := qc.repo.AddNewQuest(quest)
	if err != nil {
		return nil, err
	}

	return quest, nil

}
