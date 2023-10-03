package machinations

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain"
	"github.com/therobertcrocker/ulsidor/utils"
)

type FactionCodex struct {
	Factions  map[string]*Faction
	Changelog []*domain.LogEntry
	Storage   domain.StorageManager
}

// NewFactionCodex returns a new instance of FactionCodex.
func NewFactionCodex(storage domain.StorageManager) *FactionCodex {
	return &FactionCodex{
		Factions:  make(map[string]*Faction),
		Changelog: make([]*domain.LogEntry, 0),
		Storage:   storage,
	}
}

// Init initializes the FactionCodex.
func (fc *FactionCodex) Init() error {
	utils.Log.Debug("initializing faction codex")
	err := fc.Storage.LoadCodex("factions", fc)
	if err != nil {
		return err
	}
	return nil
}

// AddFaction adds a new faction to the FactionCodex.
func (fc *FactionCodex) AddFaction(factionInput *FactionInput) error {
	return nil
}

// Clear clears the FactionCodex.
func (fc *FactionCodex) Clear() error {
	fc.Factions = make(map[string]*Faction)

	//save the change
	if err := fc.Storage.SaveCodex("factions", fc); err != nil {
		return fmt.Errorf("failed to save faction codex: %w", err)
	}

	//log the change
	log := domain.LogEntry{
		Timestamp:   utils.GetTimestamp(),
		EntityID:    "factions",
		Description: "FactionCodex cleared",
		Keywords:    []string{"codex", "faction", "clear"},
	}
	log.AddChangeDetail("reset", nil, nil)
	fc.Changelog = append(fc.Changelog, &log)

	return nil

}
