package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	"github.com/therobertcrocker/ulsidor/internal/domain/types/changelog"
)

// EntityRepository is the interface for the EntityRepository
type EntityRepository interface {
	interfaces.Repository
	GetEntityByID(id string) (*interfaces.Entity, error)
	AddNewEntity(quest interfaces.Entity, log changelog.LogEntry) error
	UpdateEntity(id string, quest interfaces.Entity) error
	DeleteEntity(id string) error
}

type EntityRepo struct {
	Metadata  interfaces.Metadata          `json:"metadata"`
	Entities  map[string]interfaces.Entity `json:"entities"`
	ChangeLog []changelog.LogEntry         `json:"changelog"`
	Storage   interfaces.StorageManager
}

type EntityFile struct {
	Metadata  interfaces.Metadata          `json:"metadata"`
	Changelog []changelog.LogEntry         `json:"changelog"`
	Data      map[string]interfaces.Entity `json:"data"`
}

// NewEntityRepo returns a new instance of EntityRepo
func NewEntityRepo(storage interfaces.StorageManager) *EntityRepo {
	return &EntityRepo{
		Entities:  make(map[string]interfaces.Entity),
		Metadata:  interfaces.Metadata{},
		ChangeLog: make([]changelog.LogEntry, 0),
		Storage:   storage,
	}
}

// Init initializes the repo
func (er *EntityRepo) Init(entityType string) error {
	err := er.Storage.LoadRepo(entityType, er)
	if err != nil {
		return fmt.Errorf("failed to load entity repository: %w", err)
	}
	return nil
}

// LogChange logs a change to the repo
func (er *EntityRepo) LogChange(log changelog.LogEntry) {
	er.ChangeLog = append(er.ChangeLog, log)

}

// AddNewEntity adds a new entity to the repo
func (er *EntityRepo) AddNewEntity(entity interfaces.Entity, log changelog.LogEntry) error {
	er.Entities[entity.ID()] = entity
	return nil
}

// GetEntityByID fetches a entity by ID
func (er *EntityRepo) GetEntityByID(id string) (*interfaces.Entity, error) {
	if entity, exists := er.Entities[id]; exists {
		return &entity, nil
	}
	return nil, errors.New("entity not found with id " + id)
}

// UpdateEntity updates an entity in the repo
func (er *EntityRepo) UpdateEntity(id string, entity interfaces.Entity) error {
	_, exists := er.Entities[id]
	if !exists {
		return errors.New("entity not found with id " + id)
	}
	er.Entities[id] = entity
	return nil
}

// DeleteEntity deletes an entity from the repo
func (er *EntityRepo) DeleteEntity(id string) error {
	_, exists := er.Entities[id]
	if !exists {
		return errors.New("entity not found with id " + id)
	}
	delete(er.Entities, id)
	return nil
}

// Serialize serializes the repo to JSON.
func (er *EntityRepo) Serialize() (map[string][]byte, error) {
	serializedData := make(map[string][]byte)

	// serialize entity data
	data, err := json.Marshal(er.Entities)
	if err != nil {
		return nil, err
	}
	serializedData["data"] = data

	// serialize entity metadata
	metadata, err := json.Marshal(er.Metadata)
	if err != nil {
		return nil, err
	}
	serializedData["metadata"] = metadata

	// serialize entity changelog
	changelog, err := json.Marshal(er.ChangeLog)
	if err != nil {
		return nil, err
	}
	serializedData["changelog"] = changelog

	return serializedData, nil
}

// Deserialize deserializes the repo from JSON.
func (er *EntityRepo) Deserialize(collection map[string][]byte) error {
	var entityFile EntityFile

	if len(collection) == 0 {

		// new collection, populate initial metadata
		utils.Log.Debugf("New collection, populating initial Metadata")

		timeStamp := time.Now().Format(time.RFC3339)
		er.Metadata.Created = timeStamp
		er.Metadata.LastUpdated = timeStamp

		return nil
	}

	// deserialize entity data
	if err := json.Unmarshal(collection["data"], &entityFile.Data); err != nil {
		return err
	}

	// deserialize entity metadata
	if err := json.Unmarshal(collection["metadata"], &entityFile.Metadata); err != nil {
		return err
	}

	// deserialize entity changelog
	if err := json.Unmarshal(collection["changelog"], &entityFile.Changelog); err != nil {
		return err
	}

	er.Entities = entityFile.Data
	er.Metadata = entityFile.Metadata
	er.ChangeLog = entityFile.Changelog

	return nil
}
