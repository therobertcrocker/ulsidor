package factions

import "fmt"

type Faction struct {
	Name      string           `json:"name"`
	Size      string           `json:"size"`
	HomeBase  string           `json:"HomeBase"`
	Skills    SkillBlock       `json:"skills"`
	Arcane    string           `json:"arcane"`
	Treasure  int              `json:"treasure"`
	HitPoints int              `json:"hit_points"`
	CurrentHP int              `json:"current_hp"`
	XP        int              `json:"xp"`
	Income    int              `json:"income"`
	Upkeep    int              `json:"upkeep"`
	Assets    map[string]Asset `json:"assets"`
}

type SkillBlock struct {
	Cunning int `json:"cunning"`
	Force   int `json:"force"`
	Wealth  int `json:"wealth"`
}

func (f *Faction) GetID() string {
	return fmt.Sprintf("faction_%s", f.Name)
}
