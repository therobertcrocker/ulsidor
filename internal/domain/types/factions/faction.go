package factions

type Faction struct {
	ID        int              `json:"id"`
	Name      string           `json:"name"`
	Skills    SkillBlock       `json:"skills"`
	Treasure  int              `json:"treasure"`
	HitPoints int              `json:"hit_points"`
	Assets    map[string]Asset `json:"assets"`
}

type SkillBlock struct {
	Cunning int `json:"cunning"`
	Force   int `json:"force"`
	Wealth  int `json:"wealth"`
	Arcane  int `json:"arcane"`
}

type Asset struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Cost         int      `json:"cost"`
	MinumumScore int      `json:"minimum_score"`
	HitPoints    int      `json:"hit_points"`
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
