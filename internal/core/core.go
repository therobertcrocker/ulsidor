package core

import (
	"github.com/therobertcrocker/ulsidor/internal/components/quests"
	"github.com/therobertcrocker/ulsidor/internal/config"
)

type Core struct {
	questCodex *quests.QuestCodex
	questRepo  quests.QuestRepository
	coreConfig *config.Config
}

func NewCore(conf *config.Config) *Core {
	// Load global game data
	return &Core{
		coreConfig: conf,
	}
}

func (c *Core) InitQuestComponents() {
	c.questRepo = quests.NewQuestRepo(c.coreConfig)
	c.questRepo.InitQuestRepo()
	c.questCodex = quests.NewQuestCodex(c.questRepo, c.coreConfig)
}

// CreateNewQuest creates a new quest.
func (c *Core) CreateNewQuest(title, questType, description, source string, level int) (*quests.Quest, error) {
	quest, err := c.questCodex.CreateNewQuest(title, questType, description, source, level)
	if err != nil {
		config.Log.Errorf("Failed to create new quest: %v", err)
		return nil, err
	}
	return quest, nil
}
