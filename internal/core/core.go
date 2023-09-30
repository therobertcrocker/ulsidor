package core

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/core/components/quests"
	"github.com/therobertcrocker/ulsidor/internal/data/config"
	dataManager "github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
)

type Core struct {
	questCodex *quests.QuestCodex
	questRepo  quests.QuestRepository
	storage    interfaces.StorageManager
	config     *config.Config
	gameData   *dataManager.GameData
}

func NewCore(conf *config.Config, storage interfaces.StorageManager, gd *dataManager.GameData) *Core {
	// Load global game data
	return &Core{
		config:   conf,
		storage:  storage,
		gameData: gd,
	}
}

func (c *Core) InitQuestComponents() error {
	c.questRepo = quests.NewQuestRepo(c.storage)
	if err := c.questRepo.Init("quests"); err != nil {
		return fmt.Errorf("failed to initialize quest repository: %w", err)
	}
	c.questCodex = quests.NewQuestCodex(c.questRepo, c.gameData)
	return nil
}

// CreateNewQuest creates a new quest.
func (c *Core) CreateNewQuest(questInput *types.CreateQuestInput) error {
	quest, err := c.questCodex.CreateNewQuest(questInput)
	if err != nil {
		return fmt.Errorf("quest creation failed: %w", err)
	}
	utils.Log.Infof("Created new quest %s", quest.Title)
	return nil
}
