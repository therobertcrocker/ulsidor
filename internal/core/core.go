package core

import (
	"github.com/therobertcrocker/ulsidor/internal/core/components/quests"
	"github.com/therobertcrocker/ulsidor/internal/data/config"
	dataManager "github.com/therobertcrocker/ulsidor/internal/data/game"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/domain/types"
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

func (c *Core) InitQuestComponents() {
	c.questRepo = quests.NewQuestRepo(c.storage)
	if err := c.questRepo.Init(); err != nil {
		utils.Log.Errorf("Failed to initialize quest repository: %v", err)
	}
	c.questCodex = quests.NewQuestCodex(c.questRepo, c.gameData)
}

// CreateNewQuest creates a new quest.
func (c *Core) CreateNewQuest(questInput *types.CreateQuestInput) error {
	quest, err := c.questCodex.CreateNewQuest(questInput)
	if err != nil {
		utils.Log.Errorf("Failed to create new quest: %v", err)
		return err
	}
	utils.Log.Infof("Created new quest: %v", quest)
	return nil
}
