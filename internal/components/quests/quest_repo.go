package quests

import (
	"encoding/json"
	"errors"

	"github.com/therobertcrocker/ulsidor/internal/config"
	"github.com/therobertcrocker/ulsidor/internal/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/storage"
)

// QuestRepository defines the operations that a QuestRepo should support.
type QuestRepository interface {
	interfaces.Repository
	GetQuestByID(id string) (*Quest, error)
	InitQuestRepo() error
	AddNewQuest(quest *Quest) error
	UpdateQuest(id string, quest *Quest) error
	DeleteQuest(id string) error
}

type QuestRepo struct {
	Metadata     interfaces.Metadata `json:"metadata"`
	Quests       map[string]Quest    `json:"quests"`
	Storage      interfaces.StorageManager
	questsConfig *config.Config
}

type QuestFile struct {
	Metadata interfaces.Metadata `json:"metadata"`
	Data     map[string]Quest    `json:"data"`
}

// NewQuestRepo returns a new instance of QuestRepo.
func NewQuestRepo(conf *config.Config) *QuestRepo {
	return &QuestRepo{
		Quests:       make(map[string]Quest),
		Storage:      storage.NewJSONStorageManager(conf),
		questsConfig: conf,
	}
}

// InitQuestRepo initializes the quest repository
func (qr *QuestRepo) InitQuestRepo() error {
	err := qr.Storage.LoadRepo("quests", qr)
	if err != nil {
		config.Log.Errorf("Failed to load quest repository: %v", err)
		return err
	}
	return nil
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
	qr.Storage.SaveRepo("quests", qr)
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
