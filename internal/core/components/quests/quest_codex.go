package quests

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
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
func (qc *QuestCodex) CreateNewQuest(questInput *types.CreateQuestInput) (*types.Quest, error) {
	//builds the quest object and then adds it to the repository
	utils.Log.Debugf("Creating new quest %s", questInput.Title)
	quest := &types.Quest{
		Metadata: types.QuestMetadata{
			NextObjectiveID: 0,
		},
		ID:          questInput.Title,
		Title:       questInput.Title,
		QuestType:   questInput.QuestType,
		Description: questInput.Description,
		Source:      questInput.Source,
		Level:       questInput.Level,
	}

	err := quest.CalculateRewards(qc.gameData.PartyLevel)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate quest rewards: %w", err)
	}
	utils.Log.Debugf("Quest rewards calculated: %v", quest.Reward)
	err = qc.repo.AddNewQuest(quest)
	if err != nil {
		return nil, fmt.Errorf("failed to add new quest to repository: %w", err)
	}
	utils.Log.Debugf("Quest %s successfully saved to repository", quest.Title)
	return quest, nil

}
