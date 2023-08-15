package quests

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/config"
)

type QuestCodex struct {
	questConfig *config.Config
	repo        QuestRepository
}

func NewQuestCodex(repo QuestRepository, conf *config.Config) *QuestCodex {
	return &QuestCodex{
		repo:        repo,
		questConfig: conf,
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
	fmt.Println(qc.questConfig.GameData.PartyLevel)
	quest.CalculateRewards(qc.questConfig.GameData.PartyLevel)
	err := qc.repo.AddNewQuest(quest)
	if err != nil {
		return nil, err
	}

	return quest, nil

}
