package quests

import (
	"errors"

	"github.com/therobertcrocker/ulsidor/internal/core/components/quests/data"
	"github.com/therobertcrocker/ulsidor/internal/data/utils"
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

func (q *Quest) CalculateRewards(partyLevel int) error {
	utils.Log.Debugf("Calculating rewards for quest %s", q.Title)
	exp := calcExp(q.Level, partyLevel)
	if exp == 0 {
		return errors.New("unable to calculate experience")
	}
	gold := calcGold(q.Level, partyLevel, q.QuestType)
	if gold == 0 {
		return errors.New("unable to calculate gold")
	}
	q.Reward.Experience = exp
	utils.Log.Debugf("Experience calculated: %d", exp)
	q.Reward.Gold = gold
	utils.Log.Debugf("Gold calculated: %d", gold)
	return nil
}

func calcExp(level, partyLevel int) int {
	relativeLevel := level - partyLevel

	switch {
	case relativeLevel < -2:
		return data.ExpPerLevel[-2]
	case relativeLevel >= -2 && relativeLevel <= 3:
		return data.ExpPerLevel[relativeLevel]
	case relativeLevel > 3:
		return data.ExpPerLevel[3]
	}

	return 0

}

func calcGold(level, partyLevel int, class string) int {
	relativeLevel := level - partyLevel
	gold := data.GoldByLevel[partyLevel+relativeLevel+data.ClassMultiplier[class]]
	missions := data.MissionsPerLevel(relativeLevel)

	if missions == 0 {
		return 0
	}

	return gold / missions
}
