package quests

import (
	"errors"

	"github.com/therobertcrocker/ulsidor/internal/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/storage"
)

// QuestRepository defines the operations that a QuestRepo should support.
type QuestRepository interface {
	interfaces.Repository
	GetQuestByID(id string) (*Quest, error)
	AddNewQuest(quest *Quest) error
	UpdateQuest(id string, quest *Quest) error
	DeleteQuest(id string) error
}

type QuestRepo struct {
	Quests  map[string]Quest
	Storage interfaces.StorageManager
}

// NewQuestRepo returns a new instance of QuestRepo.
func NewQuestRepo() *QuestRepo {
	return &QuestRepo{
		Quests:  make(map[string]Quest),
		Storage: storage.NewJSONStorageManager(),
	}
}

// GetQuestByID fetches a quest by ID.
func (qr *QuestRepo) GetQuestByID(id string) (*Quest, error) {
	if quest, exists := qr.Quests[id]; exists {
		return &quest, nil
	}
	return nil, errors.New("quest not found")
}

// AddNewQuest adds a new quest to the repo.
func (qr *QuestRepo) AddNewQuest(quest *Quest) error {
	if _, exists := qr.Quests[quest.ID]; exists {
		return errors.New("quest already exists")
	}
	qr.Quests[quest.ID] = *quest
	return nil
}

// UpdateQuest updates an existing quest.
func (qr *QuestRepo) UpdateQuest(id string, quest *Quest) error {
	if _, exists := qr.Quests[id]; !exists {
		return errors.New("quest not found")
	}
	qr.Quests[id] = *quest
	return nil
}

// DeleteQuest removes a quest from the repo.
func (qr *QuestRepo) DeleteQuest(id string) error {
	if _, exists := qr.Quests[id]; !exists {
		return errors.New("quest not found")
	}
	delete(qr.Quests, id)
	return nil
}

// LoadFromStorage loads the quests from storage.
func (qr *QuestRepo) LoadFromStorage(id string) error {
	return qr.Storage.LoadRepo(id, qr)
}

// SaveToStorage saves the quests to storage.
func (qr *QuestRepo) SaveToStorage(id string) error {
	return qr.Storage.SaveRepo(id, qr)
}
