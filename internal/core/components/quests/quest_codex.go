package quests

import (
	"fmt"
	"time"

	"github.com/therobertcrocker/ulsidor/internal/data"
	"github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/types/changelog"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
)

type QuestCodex struct {
	repo     data.EntityRepository
	gameData *game.GameData
}

func NewQuestCodex(repo data.EntityRepository, gd *game.GameData) *QuestCodex {
	return &QuestCodex{
		repo:     repo,
		gameData: gd,
	}
}

// CreateNewQuest creates a new quest.
func (qc *QuestCodex) CreateNewQuest(questInput *types.CreateQuestInput) (*types.Quest, error) {
	//builds the quest object and then adds it to the repository
	utils.Log.Debugf("Creating new quest %s", questInput.Title)
	timestamp := time.Now().Format(time.RFC3339)
	quest := types.Quest{
		Metadata: types.QuestMetadata{
			NextObjectiveID: 0,
		},
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

	// create log entry
	logEntry := changelog.LogEntry{
		Timestamp:   timestamp,
		EntityID:    quest.ID(),
		Description: fmt.Sprintf("Quest %s created", quest.Title),
		Keywords:    []string{"ADD_QUEST"},
		Changes: []changelog.ChangeDetail{
			{
				ReferenceID: quest.ID(),
				OldValue:    "",
				NewValue:    quest.Title,
			},
		},
	}

	err = qc.repo.AddNewEntity(quest, logEntry)
	if err != nil {
		return nil, fmt.Errorf("failed to add new quest to repository: %w", err)
	}
	utils.Log.Debugf("Quest %s successfully saved to repository", quest.Title)
	return &quest, nil

}
