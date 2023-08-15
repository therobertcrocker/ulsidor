package core

import (
	"log"

	"github.com/therobertcrocker/ulsidor/internal/components/quests"
	"github.com/therobertcrocker/ulsidor/internal/config"
)

var (
	gameData *config.GameData
)

type Core struct {
	questCodex *quests.QuestCodex
	questRepo  quests.QuestRepository
}

func NewCore() *Core {
	// Load global game data
	var err error
	gameData, err = config.LoadConfig("../config/game_data.json")
	if err != nil {
		log.Fatalf("Failed to load global game data: %v", err)
	}
	return &Core{}
}

func (c *Core) InitQuestComponents() {
	c.questRepo = quests.NewQuestRepo()
	c.questRepo.LoadFromStorage("internal/storage/save_data/quests.json")
	c.questCodex = quests.NewQuestCodex(c.questRepo, gameData)
}

// CreateNewQuest creates a new quest.
func (c *Core) CreateNewQuest(title, questType, description, source string, level int) (*quests.Quest, error) {
	quest, err := c.questCodex.CreateNewQuest(title, questType, description, source, level)
	if err != nil {
		log.Fatalf("Failed to create new quest: %v", err)
	}
	return quest, nil
}
