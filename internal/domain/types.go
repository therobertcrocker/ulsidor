package domain

type LogEntry struct {
	Timestamp   string         `json:"timestamp"`
	EntityID    string         `json:"entityID"`
	Description string         `json:"description"`
	Keywords    []string       `json:"keywords"`
	Changes     []ChangeDetail `json:"changes"`
}

type ChangeDetail struct {
	ReferenceID string      `json:"referenceID"`
	OldValue    interface{} `json:"oldValue"`
	NewValue    interface{} `json:"newValue"`
}

func (l *LogEntry) AddChangeDetail(referenceID string, oldValue interface{}, newValue interface{}) {
	l.Changes = append(l.Changes, ChangeDetail{
		ReferenceID: referenceID,
		OldValue:    oldValue,
		NewValue:    newValue,
	})
}
