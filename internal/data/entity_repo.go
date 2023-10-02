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

// EntityRepository defines the operations that a EntityRepo should support.
type EntityRepository interface {
	interfaces.Repository
	GetEntityByID(id string) (*interfaces.Entity, error)
	AddNewEntity(entity *interfaces.Entity, log changelog.LogEntry) error
	UpdateEntity(id string, entity *interfaces.Entity) error
	DeleteEntity(id string) error
}

// EntityRepo is the struct for the entity repository.
type EntityRepo struct {
	Metadata  interfaces.Metadata          `json:"metadata"`
	Entities  map[string]interfaces.Entity `json:"entities"`
	ChangeLog []changelog.LogEntry         `json:"changelog"`
	Storage   interfaces.StorageManager
}

// EntityFile is the struct for the entity file.
type EntityFile struct {
	Metadata  interfaces.Metadata          `json:"metadata"`
	Changelog []changelog.LogEntry         `json:"changelog"`
	Data      map[string]interfaces.Entity `json:"data"`
}

// NewEntityRepo returns a new instance of EntityRepo.
func NewEntityRepo(storage interfaces.StorageManager) *EntityRepo {
	return &EntityRepo{
		Entities:  make(map[string]interfaces.Entity),
		Metadata:  interfaces.Metadata{},
		ChangeLog: make([]changelog.LogEntry, 0),
		Storage:   storage,
	}
}

// Init initializes the repo.
func (er *EntityRepo) Init(entityType string) error {
	err := er.Storage.LoadRepo(entityType, er)
	if err != nil {
		return fmt.Errorf("failed to load entity repository: %w", err)
	}
	return nil
}

// LogChange logs a change to the entity repository.
func (er *EntityRepo) LogChange(entry changelog.LogEntry) {
	er.ChangeLog = append(er.ChangeLog, entry)

}

// GetEntityByID returns an entity by ID.
func (er *EntityRepo) GetEntityByID(id string) (*interfaces.Entity, error) {
	entity, ok := er.Entities[id]
	if !ok {
		return nil, errors.New("entity not found")
	}
	return &entity, nil
}

// AddNewEntity adds a new entity to the repository.
func (er *EntityRepo) AddNewEntity(entity interfaces.Entity, log changelog.LogEntry) error {
	// if the entity already exists, return an error
	if _, ok := er.Entities[entity.GetID()]; ok {
		return fmt.Errorf("entity %s already exists", entity.GetID())
	}

	// add the entity to the repository
	er.Entities[entity.GetID()] = entity

	// log the change
	er.LogChange(log)

	return nil
}

// UpdateEntity updates an entity in the repository.
func (er *EntityRepo) UpdateEntity(id string, entity interfaces.Entity) error {
	// if the entity already exists, return an error
	if _, ok := er.Entities[id]; !ok {
		return fmt.Errorf("entity %s does not exist", id)
	}

	// update the entity in the repository
	er.Entities[id] = entity

	return nil
}

// DeleteEntity deletes an entity from the repository.
func (er *EntityRepo) DeleteEntity(id string) error {
	// if the entity already exists, return an error
	if _, ok := er.Entities[id]; !ok {
		return fmt.Errorf("entity %s does not exist", id)
	}

	// delete the entity from the repository
	delete(er.Entities, id)

	return nil
}

// Serialize serializes the repo to JSON.
func (er *EntityRepo) Serialize() (map[string][]byte, error) {
	serializedData := make(map[string][]byte)

	// serialize quest data
	data, err := json.Marshal(er.Entities)
	if err != nil {
		return nil, err
	}
	serializedData["data"] = data

	// serialize quest metadata
	metadata, err := json.Marshal(er.Metadata)
	if err != nil {
		return nil, err
	}
	serializedData["metadata"] = metadata

	// serialize quest changelog
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

	// deserialize quest data
	if err := json.Unmarshal(collection["data"], &entityFile.Data); err != nil {
		return err
	}

	// deserialize quest metadata
	if err := json.Unmarshal(collection["metadata"], &entityFile.Metadata); err != nil {
		return err
	}

	// deserialize quest changelog
	if err := json.Unmarshal(collection["changelog"], &entityFile.Changelog); err != nil {
		return err
	}

	// set quest data
	er.Entities = entityFile.Data
	er.Metadata = entityFile.Metadata
	er.ChangeLog = entityFile.Changelog
	return nil

}
