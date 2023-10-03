package machinations

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain"
	"github.com/therobertcrocker/ulsidor/utils"
)

// ================= Faction Types =================

type Faction struct {
	Name      string           `json:"name"`
	Size      string           `json:"size"`
	HomeBase  string           `json:"home_base"`
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

// ================= Asset Types =================

type Asset struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"`
	Cost        int       `json:"cost"`
	Upkeep      int       `json:"upkeep"`
	Threshold   int       `json:"purchase_threshold"`
	HitPoints   int       `json:"hit_points"`
	Arcane      string    `json:"arcane"`
	HasAttack   bool      `json:"has_attack"`
	Attack      *Attack   `json:"attack"`
	HasCounter  bool      `json:"has_counter"`
	Counter     *Damage   `json:"counter"`
	Qualities   []string  `json:"qualities"`
	Data        AssetData `json:"data"`
}

type AssetData struct {
	Owner      string `json:"owner"`
	LocationID string `json:"location_id"`
	CurrentHP  int    `json:"current_hp"`
}

type Attack struct {
	AttackerTrait string `json:"attacker_trait"`
	DefenderTrait string `json:"defender_trait"`
	Damage        Damage `json:"damage"`
}

type Damage struct {
	Count    int `json:"count"`
	Dice     int `json:"dice"`
	Modifier int `json:"modifier"`
}

type AssetInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Type        string   `json:"type"`
	Cost        int      `json:"cost"`
	Upkeep      int      `json:"upkeep"`
	Threshold   int      `json:"purchase_threshold"`
	HitPoints   int      `json:"hit_points"`
	Arcane      string   `json:"arcane"`
	HasAttack   bool     `json:"has_attack"`
	Attack      *Attack  `json:"attack"`
	HasCounter  bool     `json:"has_counter"`
	Counter     *Damage  `json:"counter"`
	Qualities   []string `json:"qualities"`
}

type AttackInput struct {
	AttackerTrait string `json:"attacker_trait"`
	DefenderTrait string `json:"defender_trait"`
}

var _ domain.Entity = new(Asset)

func (a *Asset) ID() string {
	return fmt.Sprintf("asset_%s_%s", a.Type, utils.Normalize(a.Name))
}

// Display returns a string representation of the asset.
func (a *Asset) Display() string {
	return fmt.Sprintf("%s - %s | cost: %d | hp: %d | arcane: %s | attack: %s | dmg: %s | counter: %s | qualities: %s",
		a.Name, a.Type, a.Cost, a.HitPoints, a.Arcane, a.AttackString(), a.DamageString(), a.CounterString(), a.Qualities)
}

func (a *Asset) AttackString() string {
	if a.HasAttack {
		return fmt.Sprintf("%s vs %s", a.Attack.AttackerTrait, a.Attack.DefenderTrait)
	}
	return "none"
}

func (a *Asset) DamageString() string {
	if a.HasAttack {
		return fmt.Sprintf("%dd%d+%d", a.Attack.Damage.Count, a.Attack.Damage.Dice, a.Attack.Damage.Modifier)
	}

	return "none"
}

func (a *Asset) CounterString() string {
	if a.HasCounter {
		return fmt.Sprintf("%dd%d+%d", a.Counter.Count, a.Counter.Dice, a.Counter.Modifier)
	}

	return "none"
}
