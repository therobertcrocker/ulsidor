package changelog

import (
	"fmt"
	"strings"

	"github.com/therobertcrocker/ulsidor/internal/domain/types/quests"
)

type NewQuestChange struct {
	Field string       `json:"field"`
	Value quests.Quest `json:"value"`
}

func (nqc *NewQuestChange) DescribeChange() string {
	return fmt.Sprintf("New quest created: Title: %s, Level: %d, Type: %s", nqc.Value.Title, nqc.Value.Level, nqc.Value.QuestType)
}

type ObjectiveChange struct {
	Field    string           `json:"field"`
	OldValue quests.Objective `json:"oldValue"`
	NewValue quests.Objective `json:"newValue"`
}

func (oc *ObjectiveChange) DescribeChange() string {
	var changes []string

	if oc.OldValue.ID != oc.NewValue.ID {
		changes = append(changes, fmt.Sprintf("ID changed from %d to %d", oc.OldValue.ID, oc.NewValue.ID))
	}
	if oc.OldValue.Title != oc.NewValue.Title {
		changes = append(changes, fmt.Sprintf("Title changed from %s to %s", oc.OldValue.Title, oc.NewValue.Title))
	}
	if oc.OldValue.Completed != oc.NewValue.Completed {
		changes = append(changes, fmt.Sprintf("Completed status changed from %v to %v", oc.OldValue.Completed, oc.NewValue.Completed))
	}
	// ... similarly for other fields

	return fmt.Sprintf("Objective changes: %s", strings.Join(changes, ", "))
}

type RewardChange struct {
	Field    string        `json:"field"`
	OldValue quests.Reward `json:"oldValue"`
	NewValue quests.Reward `json:"newValue"`
}

func (rc *RewardChange) DescribeChange() string {
	var changes []string

	if rc.OldValue.Experience != rc.NewValue.Experience {
		changes = append(changes, fmt.Sprintf("Experience changed from %d to %d", rc.OldValue.Experience, rc.NewValue.Experience))
	}
	if rc.OldValue.Gold != rc.NewValue.Gold {
		changes = append(changes, fmt.Sprintf("Gold changed from %d to %d", rc.OldValue.Gold, rc.NewValue.Gold))
	}
	if rc.OldValue.TreasureRating != rc.NewValue.TreasureRating {
		changes = append(changes, fmt.Sprintf("Treasure rating changed from %d to %d", rc.OldValue.TreasureRating, rc.NewValue.TreasureRating))
	}
	if rc.OldValue.Reputation != rc.NewValue.Reputation {
		changes = append(changes, fmt.Sprintf("Reputation changed from %d to %d", rc.OldValue.Reputation, rc.NewValue.Reputation))
	}
	// ... similarly for other fields

	return fmt.Sprintf("Reward changes: %s", strings.Join(changes, ", "))
}
