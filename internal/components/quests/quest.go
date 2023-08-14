package quests

import (
	"github.com/therobertcrocker/ulsidor/internal/components/quests/data"
)

type Quest struct {
	ID          string      `json:"id"`
	Title       string      `json:"title"`
	QuestType   string      `json:"quest_type"`
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Level       int         `json:"level"`
	Source      string      `json:"source"`
	Reward      Reward      `json:"reward"`
	Objectives  []Objective `json:"objectives"`
}

type Reward struct {
	Experience     int `json:"experience"`
	Gold           int `json:"gold"`
	TreasureRating int `json:"treasure_rating"`
	Repuation      int `json:"reputation"`
}

type Objective struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

// GetID returns the Quest ID.
func (q *Quest) GetID() string {
	return q.ID
}

func (q *Quest) CalculateRewards(partyLevel int) {
	q.Reward.Experience = calcExp(q.Level, partyLevel)
	q.Reward.Gold = calcGold(q.Level, partyLevel, q.QuestType)
}

func calcExp(level, partyLevel int) int {
	return data.ExpPerLevel[level-partyLevel]
}

func calcGold(level, partyLevel int, class string) int {
	relativeLevel := level - partyLevel
	return data.GoldByLevel[partyLevel+relativeLevel+data.ClassMultiplier[class]] / data.MissionsPerLevel(relativeLevel)
}
