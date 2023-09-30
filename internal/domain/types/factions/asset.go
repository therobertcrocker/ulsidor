package factions

import "fmt"

type Asset struct {
	Name         string   `json:"name"`
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

func (a *Asset) GetID() string {
	return fmt.Sprintf("asset_%s", a.Name)
}
