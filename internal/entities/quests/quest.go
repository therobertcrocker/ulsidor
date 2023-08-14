package questforge

type Quest struct {
	Title       string `json:"title"`
	QuestClass  string `json:"quest_class"`
	Description string `json:"description"`
	Level       int    `json:"level"`
	Source      string `json:"source"`
	ImagePath   string `json:"image_path"`
	Reward      Reward `json:"reward"`
}

type Reward struct {
	Experience     int `json:"experience"`
	Gold           int `json:"gold"`
	TreasureRating int `json:"treasure_rating"`
	Repuation      int `json:"reputation"`
}

func (q *Quest) CalculateRewards(partyLevel int) {
	q.Reward.Experience = calcExp(q.Level, partyLevel)
	q.Reward.Gold = calcGold(q.Level, partyLevel, q.QuestClass)
}

func calcExp(level, partyLevel int) int {
	return expPerLevel[level-partyLevel]
}

func calcGold(level, partyLevel int, class string) int {
	relativeLevel := level - partyLevel
	return goldbyLevel[partyLevel+relativeLevel+classMultiplier[class]] / missionsPerLevel(relativeLevel)
}
