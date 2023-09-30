package factions

import (
	"encoding/json"
	"time"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/domain/types/changelog"
	types "github.com/therobertcrocker/ulsidor/internal/domain/types/factions"
)

// FactionRepository is the interface for the FactionRepository
type FactionRepository interface {
	interfaces.Repository
	AddNewFaction(faction *types.Faction) error
}

// FactionRepo is the struct for the FactionRepo
type FactionRepo struct {
	Metadata  interfaces.Metadata
	ChangeLog []changelog.LogEntry
	Factions  map[string]types.Faction
	Storage   interfaces.StorageManager
}

// FactionFile is the struct for the FactionFile
type FactionFile struct {
	Metadata  interfaces.Metadata
	Changelog []changelog.LogEntry
	Data      map[string]types.Faction
}

// NewFactionRepo returns a new instance of FactionRepo

func NewFactionRepo(storage interfaces.StorageManager) *FactionRepo {
	return &FactionRepo{
		Factions: make(map[string]types.Faction),
		Metadata: interfaces.Metadata{},
		Storage:  storage,
	}
}

// Init initializes the repo
func (fr *FactionRepo) Init() error {
	err := fr.Storage.LoadRepo("factions", fr)
	if err != nil {
		return err
	}
	return nil
}

// LogChange logs a change to the repo
func (fr *FactionRepo) LogChange(entry changelog.LogEntry) {
	fr.ChangeLog = append(fr.ChangeLog, entry)

}

// AddNewFaction adds a new faction to the repo
func (fr *FactionRepo) AddNewFaction(faction *types.Faction) error {
	fr.Factions[faction.GetID()] = *faction
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
