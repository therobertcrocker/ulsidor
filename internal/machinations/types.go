package machinations

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain"
)

type Faction struct {
	Name      string           `json:"name"`
	Size      string           `json:"size"`
	HomeBase  string           `json:"HomeBase"`
	Traits    TraitBlock       `json:"skills"`
	Arcane    string           `json:"arcane"`
	Treasure  int              `json:"treasure"`
	HitPoints int              `json:"hit_points"`
	CurrentHP int              `json:"current_hp"`
	XP        int              `json:"xp"`
	Income    int              `json:"income"`
	Upkeep    int              `json:"upkeep"`
	Assets    map[string]Asset `json:"assets"`
}

type TraitBlock struct {
	Cunning int `json:"cunning"`
	Force   int `json:"force"`
	Wealth  int `json:"wealth"`
}

type Asset struct {
	Name         string   `json:"name"`
	Owner        string   `json:"owner"`
	Type         string   `json:"type"`
	Cost         int      `json:"cost"`
	LocationID   string   `json:"location_id"`
	MinumumScore int      `json:"minimum_score"`
	HitPoints    int      `json:"hit_points"`
	CurrentHP    int      `json:"current_hp"`
	Upkeep       int      `json:"upkeep"`
	Arcane       int      `json:"arcane"`
	Attack       Attack   `json:"attack"`
	Counter      Damage   `json:"counter"`
	Qualites     []string `json:"qualities"`
}
type Attack struct {
	Trait  string `json:"trait"`
	Damage Damage `json:"damage"`
}

type Damage struct {
	Count    int `json:"count"`
	Dice     int `json:"dice"`
	Modifier int `json:"modifier"`
}

type FactionInput struct {
	Name   string     `json:"name"`
	Size   string     `json:"size"`
	Traits TraitBlock `json:"skills"`
	Arcane string     `json:"arcane"`
}

var _ domain.Entity = new(Faction)

func (f *Faction) ID() string {
	return fmt.Sprintf("faction_%s", f.Name)
}
