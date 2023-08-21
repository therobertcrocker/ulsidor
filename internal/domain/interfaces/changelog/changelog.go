package changelog

import (
	"fmt"
)

type ChangeInterface interface {
	//DescribeChange returns a description of the change
	DescribeChange() string
}

type LogEntry struct {
	Timestamp   int64             `json:"timestamp"` // Unix timestamp
	Action      string            `json:"action"`
	EntityID    string            `json:"entityID"`
	Description string            `json:"description"`
	Changes     []ChangeInterface `json:"changes"`
}

type StringChange struct {
	Field    string `json:"field"`
	OldValue string `json:"oldValue"`
	NewValue string `json:"newValue"`
}

func (sc *StringChange) DescribeChange() string {
	return fmt.Sprintf("%s changed from %s to %s", sc.Field, sc.OldValue, sc.NewValue)
}

type IntChange struct {
	Field    string `json:"field"`
	OldValue int    `json:"oldValue"`
	NewValue int    `json:"newValue"`
}

func (ic *IntChange) DescribeChange() string {
	return fmt.Sprintf("%s changed from %d to %d", ic.Field, ic.OldValue, ic.NewValue)
}
