package quests

import (
	"encoding/json"
	"errors"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
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
	Metadata interfaces.Metadata `json:"metadata"`
	Quests   map[string]Quest    `json:"quests"`
	Storage  interfaces.StorageManager
}

type QuestFile struct {
	Metadata interfaces.Metadata `json:"metadata"`
	Data     map[string]Quest    `json:"data"`
}

// NewQuestRepo returns a new instance of QuestRepo.
func NewQuestRepo(storage interfaces.StorageManager) *QuestRepo {
	return &QuestRepo{
		Quests:  make(map[string]Quest),
		Storage: storage,
	}
}

// Init initializes the repo.
func (qr *QuestRepo) Init() error {
	return qr.Storage.LoadRepo("quests", qr)
}

// GetQuestByID fetches a quest by ID.
func (qr *QuestRepo) GetQuestByID(id string) (*Quest, error) {
	if quest, exists := qr.Quests[id]; exists {
		return &quest, nil
	}
	return nil, errors.New("quest not found with id " + id)
}

// AddNewQuest adds a new quest to the repo.
func (qr *QuestRepo) AddNewQuest(quest *Quest) error {
	utils.Log.Debugf("Adding new quest %s to repository", quest.Title)

	// Check if quest already exists
	utils.Log.Debugf("Checking if quest already exists with id %s", quest.ID)
	if _, exists := qr.Quests[quest.ID]; exists {
		return errors.New("quest already exists with id " + quest.ID + "already exists")
	}

	// Add quest to repo
	qr.Quests[quest.ID] = *quest

	// Save repo
	utils.Log.Debugf("Saving quest %s to repository", quest.Title)
	qr.Storage.SaveRepo("quests", qr)
	return nil
}

// UpdateQuest updates an existing quest.
func (qr *QuestRepo) UpdateQuest(id string, quest *Quest) error {
	if _, exists := qr.Quests[id]; !exists {
		return errors.New("quest not found with id " + id)
	}
	qr.Quests[id] = *quest
	return nil
}

// DeleteQuest removes a quest from the repo.
func (qr *QuestRepo) DeleteQuest(id string) error {
	if _, exists := qr.Quests[id]; !exists {
		return errors.New("quest not found with id " + id)
	}
	delete(qr.Quests, id)
	return nil
}

// Serialize serializes the repo to JSON.
func (qr *QuestRepo) Serialize() ([]byte, error) {
	questsFile := QuestFile{
		Metadata: qr.Metadata,
		Data:     qr.Quests,
	}
	return json.Marshal(questsFile)
}

// Deserialize deserializes the repo from JSON.
func (qr *QuestRepo) Deserialize(data []byte) error {
	var questsFile QuestFile
	if err := json.Unmarshal(data, &questsFile); err != nil {
		return err
	}
	qr.Metadata = questsFile.Metadata
	qr.Quests = questsFile.Data

	return nil
}
