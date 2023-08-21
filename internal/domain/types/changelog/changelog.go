package changelog

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
