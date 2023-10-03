package core

import (
	"github.com/therobertcrocker/ulsidor/internal/domain"
	"github.com/therobertcrocker/ulsidor/internal/machinations"
)

// Core handles the initialization of the app's core components.
type Core struct {
	Factions FactionConfig
	Storage  domain.StorageManager
}

type FactionConfig struct {
	Codex  *machinations.FactionCodex
	Assets *machinations.AssetCodex
}

// NewCore returns a new instance of Core.
func NewCore(storage domain.StorageManager) *Core {
	return &Core{
		Factions: FactionConfig{
			Codex:  machinations.NewFactionCodex(storage),
			Assets: machinations.NewAssetCodex(storage),
		},
		Storage: storage,
	}
}

// Init initializes the core components.
func (c *Core) Init() error {
	if err := c.Factions.Codex.Init(); err != nil {
		return err
	}
	if err := c.Factions.Assets.Init(); err != nil {
		return err
	}
	return nil
}
