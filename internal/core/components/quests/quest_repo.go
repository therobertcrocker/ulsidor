package quests

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces/changelog"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
)

// QuestRepository defines the operations that a QuestRepo should support.
type QuestRepository interface {
	interfaces.Repository
	GetQuestByID(id string) (*types.Quest, error)
	AddNewQuest(quest *types.Quest) error
	UpdateQuest(id string, quest *types.Quest) error
	DeleteQuest(id string) error
}

type QuestRepo struct {
	Metadata  interfaces.Metadata    `json:"metadata"`
	Quests    map[string]types.Quest `json:"quests"`
	ChangeLog []changelog.LogEntry   `json:"changelog"`
	Storage   interfaces.StorageManager
}

type QuestFile struct {
	Metadata  interfaces.Metadata    `json:"metadata"`
	Changelog []changelog.LogEntry   `json:"changelog"`
	Data      map[string]types.Quest `json:"data"`
}

// NewQuestRepo returns a new instance of QuestRepo.
func NewQuestRepo(storage interfaces.StorageManager) *QuestRepo {
	return &QuestRepo{
		Quests:    make(map[string]types.Quest),
		Metadata:  interfaces.Metadata{},
		ChangeLog: make([]changelog.LogEntry, 0),
		Storage:   storage,
	}
}

// Init initializes the repo.
func (qr *QuestRepo) Init() error {
	err := qr.Storage.LoadRepo("quests", qr)
	if err != nil {
		return fmt.Errorf("failed to load quest repository: %w", err)
	}
	return nil
}

// GetQuestByID fetches a quest by ID.
func (qr *QuestRepo) GetQuestByID(id string) (*types.Quest, error) {
	if quest, exists := qr.Quests[id]; exists {
		return &quest, nil
	}
	return nil, errors.New("quest not found with id " + id)
}

// AddNewQuest adds a new quest to the repo.
func (qr *QuestRepo) AddNewQuest(quest *types.Quest) error {
	utils.Log.Debugf("Adding new quest %s to repository", quest.Title)

	// Check if quest already exists
	utils.Log.Debugf("Checking if quest already exists with id %s", quest.ID)
	if _, exists := qr.Quests[quest.ID]; exists {
		return errors.New("quest already exists with id " + quest.ID + "already exists")
	}

	// Add quest to repo
	qr.Quests[quest.ID] = *quest

	// Update metadata
	qr.Metadata.LastUpdated = time.Now().Format(time.RFC3339)

	// Save repo
	utils.Log.Debugf("Saving quest %s to repository", quest.Title)
	qr.Storage.SaveRepo("quests", qr)
	return nil
}

// UpdateQuest updates an existing quest.
func (qr *QuestRepo) UpdateQuest(id string, updatedQuest *types.Quest) error {
	_, exists := qr.Quests[id]
	if !exists {
		return errors.New("quest not found with id " + id)
	}

	qr.Quests[id] = *updatedQuest
	return nil
}

func detectChanges(oldQuest *types.Quest, newQuest *types.Quest) []changelog.ChangeInterface {
	var changes []changelog.ChangeInterface

	// check for changes to the quest title
	if oldQuest.Title != newQuest.Title {
		changes = append(changes, &changelog.StringChange{
			Field:    "title",
			OldValue: oldQuest.Title,
			NewValue: newQuest.Title,
		})
	}

	// check for changes to the quest type
	if oldQuest.QuestType != newQuest.QuestType {
		changes = append(changes, &changelog.StringChange{
			Field:    "quest_type",
			OldValue: oldQuest.QuestType,
			NewValue: newQuest.QuestType,
		})
	}

	// check for changes to the quest description
	if oldQuest.Description != newQuest.Description {
		changes = append(changes, &changelog.StringChange{
			Field:    "description",
			OldValue: oldQuest.Description,
			NewValue: newQuest.Description,
		})
	}

	// check for changes to the quest source
	if oldQuest.Source != newQuest.Source {
		changes = append(changes, &changelog.StringChange{
			Field:    "source",
			OldValue: oldQuest.Source,
			NewValue: newQuest.Source,
		})
	}

	// check for changes to the quest level
	if oldQuest.Level != newQuest.Level {
		changes = append(changes, &changelog.IntChange{
			Field:    "level",
			OldValue: oldQuest.Level,
			NewValue: newQuest.Level,
		})
	}

	return changes
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
func (qr *QuestRepo) Serialize() (map[string][]byte, error) {
	serielizedData := make(map[string][]byte)

	// serialize quest data
	data, err := json.Marshal(qr.Quests)
	if err != nil {
		return nil, err
	}
	serielizedData["data"] = data

	// serialize quest metadata
	metadata, err := json.Marshal(qr.Metadata)
	if err != nil {
		return nil, err
	}
	serielizedData["metadata"] = metadata

	// serialize quest changelog
	changelog, err := json.Marshal(qr.ChangeLog)
	if err != nil {
		return nil, err
	}
	serielizedData["changelog"] = changelog

	return serielizedData, nil
}

// Deserialize deserializes the repo from JSON.
func (qr *QuestRepo) Deserialize(collection map[string][]byte) error {
	var questsFile QuestFile

	if collection == nil || len(collection) == 0 {

		// new collection, populate initial metadata
		utils.Log.Debugf("New collection, populating initial Metadata")

		timeStamp := time.Now().Format(time.RFC3339)
		qr.Metadata.Created = timeStamp
		qr.Metadata.LastUpdated = timeStamp

		return nil
	}

	// deserialize quest data
	if err := json.Unmarshal(collection["data"], &questsFile.Data); err != nil {
		return err
	}

	// deserialize quest metadata
	if err := json.Unmarshal(collection["metadata"], &questsFile.Metadata); err != nil {
		return err
	}

	// deserialize quest changelog
	if err := json.Unmarshal(collection["changelog"], &questsFile.Changelog); err != nil {
		return err
	}

	// set quest data
	qr.Quests = questsFile.Data
	qr.Metadata = questsFile.Metadata
	qr.ChangeLog = questsFile.Changelog
	return nil

}
