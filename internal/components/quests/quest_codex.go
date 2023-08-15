package quests

type QuestCodex struct {
	repo QuestRepository
}

func NewQuestCodex(repo QuestRepository) *QuestCodex {
	return &QuestCodex{
		repo: repo,
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

	return quest, nil

}
