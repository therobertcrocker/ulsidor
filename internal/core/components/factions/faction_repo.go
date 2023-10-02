package factions

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/domain/types/changelog"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/factions"
)

// use the same structure as quest_repo.go
type FactionRepository interface {
	interfaces.Repository
	GetFactionByID(id string) (*types.Faction, error)
	AddNewFaction(faction *types.Faction, log changelog.LogEntry) error
	UpdateFaction(id string, faction *types.Faction) error
	DeleteFaction(id string) error
}

// use the same structure as quest_repo.go
type FactionRepo struct {
	Metadata  interfaces.Metadata      `json:"metadata"`
	Factions  map[string]types.Faction `json:"factions"`
	ChangeLog []changelog.LogEntry     `json:"changelog"`
	Storage   interfaces.StorageManager
}

// use the same structure as quest_repo.go
type FactionFile struct {
	Metadata  interfaces.Metadata      `json:"metadata"`
	Changelog []changelog.LogEntry     `json:"changelog"`
	Data      map[string]types.Faction `json:"data"`
}

// use the same structure as quest_repo.go
func NewFactionRepo(storage interfaces.StorageManager) *FactionRepo {
	return &FactionRepo{
		Factions:  make(map[string]types.Faction),
		Metadata:  interfaces.Metadata{},
		ChangeLog: make([]changelog.LogEntry, 0),
		Storage:   storage,
	}
}

// use the same structure as quest_repo.go
func (fr *FactionRepo) Init(EntityType string) error {
	err := fr.Storage.LoadRepo("factions", fr)
	if err != nil {
		return fmt.Errorf("failed to load faction repository: %w", err)
	}
	return nil
}

// use the same structure as quest_repo.go
func (fr *FactionRepo) LogChange(entry changelog.LogEntry) {
	fr.ChangeLog = append(fr.ChangeLog, entry)
}

// AddNewFaction adds a new faction to the repository.
func (fr *FactionRepo) AddNewFaction(faction *types.Faction, log changelog.LogEntry) error {
	// if the faction already exists, return an error
	if _, ok := fr.Factions[faction.GetID()]; ok {
		return fmt.Errorf("faction %s already exists", faction.GetID())
	}

	// add the faction to the repository
	fr.Factions[faction.GetID()] = *faction

	// log the change
	fr.LogChange(log)

	return nil
}

// GetFactionByID fetches a faction by ID.
func (qr *FactionRepo) GetFactionByID(id string) (*types.Faction, error) {
	if faction, exists := qr.Factions[id]; exists {
		return &faction, nil
	}
	return nil, fmt.Errorf("faction not found with id %s", id)
}

// UpdateFaction updates a faction in the repository.
func (qr *FactionRepo) UpdateFaction(id string, faction *types.Faction) error {
	// if the faction doesn't exist, return an error
	if _, ok := qr.Factions[id]; !ok {
		return fmt.Errorf("faction %s does not exist", id)
	}

	// update the faction in the repository
	qr.Factions[id] = *faction

	// log the change

	return nil
}

// DeleteFaction removes a faction from the repository.
func (qr *FactionRepo) DeleteFaction(id string) error {
	// if the faction doesn't exist, return an error
	if _, ok := qr.Factions[id]; !ok {
		return fmt.Errorf("faction %s does not exist", id)
	}

	// delete the faction from the repository
	delete(qr.Factions, id)

	// log the change

	return nil
}

// Serialize serializes the repo to JSON.
func (fr *FactionRepo) Serialize() (map[string][]byte, error) {
	serializedData := make(map[string][]byte)

	// serialize quest data
	data, err := json.Marshal(fr.Factions)
	if err != nil {
		return nil, err
	}
	serializedData["data"] = data

	// serialize quest metadata
	metadata, err := json.Marshal(fr.Metadata)
	if err != nil {
		return nil, err
	}
	serializedData["metadata"] = metadata

	// serialize quest changelog
	changelog, err := json.Marshal(fr.ChangeLog)
	if err != nil {
		return nil, err
	}
	serializedData["changelog"] = changelog

	return serializedData, nil
}

// Deserialize deserializes the repo from JSON.
func (fr *FactionRepo) Deserialize(collection map[string][]byte) error {
	var factionFile FactionFile

	if len(collection) == 0 {

		// new collection, populate initial metadata
		utils.Log.Debugf("New collection, populating initial Metadata")

		timeStamp := time.Now().Format(time.RFC3339)
		fr.Metadata.Created = timeStamp
		fr.Metadata.LastUpdated = timeStamp

		return nil
	}

	// deserialize quest data
	if err := json.Unmarshal(collection["data"], &factionFile.Data); err != nil {
		return err
	}

	// deserialize quest metadata
	if err := json.Unmarshal(collection["metadata"], &factionFile.Metadata); err != nil {
		return err
	}

	// deserialize quest changelog
	if err := json.Unmarshal(collection["changelog"], &factionFile.Changelog); err != nil {
		return err
	}

	// set quest data
	fr.Factions = factionFile.Data
	fr.Metadata = factionFile.Metadata
	fr.ChangeLog = factionFile.Changelog
	return nil

}
