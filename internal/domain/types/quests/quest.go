package quests

import (
	"errors"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
)

type Quest struct {
	Metadata      QuestMetadata     `json:"metadata"`
	ID            string            `json:"id"`
	Title         string            `json:"title"`
	QuestType     string            `json:"quest_type"`
	Status        string            `json:"status"`
	QuestCategory string            `json:"quest_category"`
	Description   string            `json:"description"`
	Level         int               `json:"level"`
	Source        string            `json:"source"`
	Reward        Reward            `json:"reward"`
	Objectives    map[int]Objective `json:"objectives"`
}

type QuestMetadata struct {
	NextObjectiveID int `json:"next_objective_id"`
}

type Reward struct {
	Experience     int `json:"experience"`
	Gold           int `json:"gold"`
	TreasureRating int `json:"treasure_rating"`
	Reputation     int `json:"reputation"`
}

type Objective struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Completed   bool   `json:"completed"`
	Description string `json:"description"`
}

// GetID returns the Quest ID.
func (q *Quest) GetID() string {
	return q.ID
}

func (q *Quest) AddNewObjective(objective Objective) {
	objective.ID = q.Metadata.NextObjectiveID
	q.Objectives[objective.ID] = objective
	q.Metadata.NextObjectiveID++
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
		return ExpPerLevel[-2]
	case relativeLevel >= -2 && relativeLevel <= 3:
		return ExpPerLevel[relativeLevel]
	case relativeLevel > 3:
		return ExpPerLevel[3]
	}

	return 0

}

func calcGold(level, partyLevel int, class string) int {
	relativeLevel := level - partyLevel
	gold := GoldByLevel[partyLevel+relativeLevel+ClassMultiplier[class]]
	missions := MissionsPerLevel(relativeLevel)

	if missions == 0 {
		return 0
	}

	return gold / missions
}
