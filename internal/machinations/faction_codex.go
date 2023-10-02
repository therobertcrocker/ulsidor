package machinations

import (
	"github.com/therobertcrocker/ulsidor/internal/domain"
)

type FactionCodex struct {
	Factions  map[string]*Faction
	Changelog map[string][]*domain.LogEntry
	Storage   domain.StorageManager
}

// NewFactionCodex returns a new instance of FactionCodex.
func NewFactionCodex(storage domain.StorageManager) *FactionCodex {
	return &FactionCodex{
		Factions:  make(map[string]*Faction),
		Changelog: make(map[string][]*domain.LogEntry),
		Storage:   storage,
	}
}

// Init initializes the FactionCodex.
func (fc *FactionCodex) Init() error {
	err := fc.Storage.LoadCodex("factions", fc)
	if err != nil {
		return err
	}
	return nil
}
