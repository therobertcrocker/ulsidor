package machinations

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain"
	"github.com/therobertcrocker/ulsidor/utils"
)

type AssetCodex struct {
	Assets    *AssetPool
	Changelog []*domain.LogEntry
	Storage   domain.StorageManager
}

type AssetPool struct {
	Cunning map[string]*Asset
	Force   map[string]*Asset
	Wealth  map[string]*Asset
}

// NewAssetCodex returns a new instance of AssetCodex.
func NewAssetCodex(storage domain.StorageManager) *AssetCodex {
	return &AssetCodex{
		Assets: &AssetPool{
			Cunning: make(map[string]*Asset),
			Force:   make(map[string]*Asset),
			Wealth:  make(map[string]*Asset),
		},
		Changelog: make([]*domain.LogEntry, 0),
		Storage:   storage,
	}
}

// Init initializes the AssetCodex.
func (ac *AssetCodex) Init() error {
	utils.Log.Debug("initializing asset codex")
	err := ac.Storage.LoadCodex("assets", ac)
	if err != nil {
		return err
	}
	return nil
}

// AddAsset adds a new asset to the AssetCodex.
func (ac *AssetCodex) AddAsset(assetInput *AssetInput) error {
	asset := Asset{
		Name:        assetInput.Name,
		Description: assetInput.Description,
		Type:        assetInput.Type,
		Cost:        assetInput.Cost,
		Upkeep:      assetInput.Upkeep,
		Threshold:   assetInput.Threshold,
		HitPoints:   assetInput.HitPoints,
		Arcane:      assetInput.Arcane,
		Attack:      assetInput.Attack,
		Counter:     assetInput.Counter,
		Qualities:   assetInput.Qualities,
	}

	switch assetInput.Type {
	case "cunning":
		ac.Assets.Cunning[assetInput.Name] = &asset
	case "force":
		ac.Assets.Force[assetInput.Name] = &asset
	case "wealth":
		ac.Assets.Wealth[assetInput.Name] = &asset
	}

	utils.Log.Debug("added asset to codex")
	// Log the addition
	logEntry := domain.LogEntry{
		Timestamp:   utils.GetTimestamp(),
		EntityID:    asset.ID(),
		Description: fmt.Sprintf("Added %s to the codex", asset.Name),
		Keywords:    []string{"codex", "asset", "add"},
	}

	logEntry.AddChangeDetail("create", nil, asset)

	ac.Changelog = append(ac.Changelog, &logEntry)

	// Save the codex
	err := ac.Storage.SaveCodex("assets", ac)
	if err != nil {
		utils.Log.WithError(err).Error("failed to save asset codex")
		return fmt.Errorf("failed to save asset codex: %w", err)
	}

	utils.Log.Debug("assets saved to storage")
	return nil

}

// LoadAssets loads a slice of assets into the AssetCodex.
func (ac *AssetCodex) LoadAssets(assets []*Asset) error {
	utils.Log.Debug("loading assets into codex")

	for _, asset := range assets {
		switch asset.Type {
		case "cunning":
			_, ok := ac.Assets.Cunning[asset.ID()]
			if ok {
				utils.Log.WithField("asset", asset.ID()).Error("asset already exists in codex")
				return fmt.Errorf("asset already exists in codex")
			}
			ac.Assets.Cunning[asset.ID()] = asset
		case "force":
			_, ok := ac.Assets.Force[asset.ID()]
			if ok {
				utils.Log.WithField("asset", asset.ID()).Error("asset already exists in codex")
				return fmt.Errorf("asset already exists in codex")
			}
			ac.Assets.Force[asset.ID()] = asset
		case "wealth":
			_, ok := ac.Assets.Wealth[asset.ID()]
			if ok {
				utils.Log.WithField("asset", asset.ID()).Error("asset already exists in codex")
				return fmt.Errorf("asset already exists in codex")
			}
			ac.Assets.Wealth[asset.ID()] = asset
		}
	}

	utils.Log.WithField("count", len(assets)).Debug("assets loaded into codex")

	// Save the codex
	err := ac.Storage.SaveCodex("assets", ac)
	if err != nil {
		return fmt.Errorf("failed to save asset codex: %w", err)
	}

	utils.Log.Debug("assets saved to storage")

	// Log the addition
	logEntry := domain.LogEntry{
		Timestamp:   utils.GetTimestamp(),
		EntityID:    "codex",
		Description: fmt.Sprintf("Loaded %d assets into the codex", len(assets)),
		Keywords:    []string{"codex", "asset", "load"},
	}
	logEntry.AddChangeDetail("load", nil, assets)
	ac.Changelog = append(ac.Changelog, &logEntry)

	return nil
}

// GetAsset returns an asset from the AssetCodex.
func (ac *AssetCodex) GetAsset(id string) (*Asset, error) {
	utils.Log.WithField("asset", id).Debug("getting asset from codex")

	assetInfo := parseID(id)
	switch assetInfo.AssetType {
	case "cunning":
		asset, ok := ac.Assets.Cunning[id]
		if !ok {
			utils.Log.WithFields(map[string]interface{}{
				"asset": id,
				"type":  assetInfo.AssetType,
			}).Error("asset not found in codex")
			return nil, fmt.Errorf("asset not found in codex")
		}
		return asset, nil
	case "force":
		asset, ok := ac.Assets.Force[id]
		if !ok {
			utils.Log.WithFields(map[string]interface{}{
				"asset": id,
				"type":  assetInfo.AssetType,
			}).Error("asset not found in codex")
			return nil, fmt.Errorf("asset not found in codex")
		}
		return asset, nil
	case "wealth":
		asset, ok := ac.Assets.Wealth[id]
		if !ok {
			utils.Log.WithFields(map[string]interface{}{
				"asset": id,
				"type":  assetInfo.AssetType,
			}).Error("asset not found in codex")
			return nil, fmt.Errorf("asset not found in codex")
		}
		return asset, nil
	default:
		return nil, fmt.Errorf("invalid asset type")
	}

}

type AssetInfo struct {
	AssetType string
	AssetName string
}

func parseID(assetID string) AssetInfo {
	decodedID := utils.DecodeID(assetID)

	return AssetInfo{
		AssetType: decodedID[1],
		AssetName: decodedID[2],
	}

}

// Clear resets the AssetCodex.
func (ac *AssetCodex) Clear() error {
	utils.Log.Debug("resetting asset codex")

	ac.Assets = &AssetPool{
		Cunning: make(map[string]*Asset),
		Force:   make(map[string]*Asset),
		Wealth:  make(map[string]*Asset),
	}

	// Save the codex
	err := ac.Storage.SaveCodex("assets", ac)
	if err != nil {
		return fmt.Errorf("failed to save asset codex: %w", err)
	}

	utils.Log.Debug("assets saved to storage")

	// Log the addition
	logEntry := domain.LogEntry{
		Timestamp:   utils.GetTimestamp(),
		EntityID:    "codex",
		Description: "Reset the asset codex",
		Keywords:    []string{"codex", "asset", "reset"},
	}
	logEntry.AddChangeDetail("reset", nil, nil)
	ac.Changelog = append(ac.Changelog, &logEntry)

	return nil
}
