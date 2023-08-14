package quests

type QuestCodex struct {
	repo QuestRepository
}

func NewQuestCodex() *QuestCodex {
	return &QuestCodex{
		repo: NewQuestRepo(),
	}
}

func (qc *QuestCodex) GetQuestByID(id string) (*Quest, error) {
	return qc.repo.GetQuestByID(id)
}
