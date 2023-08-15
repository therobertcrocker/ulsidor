package core

import (
	"github.com/therobertcrocker/ulsidor/internal/components/quests"
)

type Core struct {
	questCodex *quests.QuestCodex
	questRepo  quests.QuestRepository
}

func NewCore() *Core {
	return &Core{}
}

func (c *Core) InitQuestComponents() {
	c.questRepo = quests.NewQuestRepo()
	c.questRepo.LoadFromStorage("internal/storage/save_data/quests.json")
	c.questCodex = quests.NewQuestCodex(c.questRepo)
}

// CreateNewQuest creates a new quest.
func (c *Core) CreateNewQuest(title, questType, description, source string, level int) (*quests.Quest, error) {
	quest, err := c.questCodex.CreateNewQuest(title, questType, description, source, level)
	if err != nil {
		return nil, err
	}
	return quest, nil
}
