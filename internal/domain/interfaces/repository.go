package interfaces

// Repository is the interface for a repository

type Repository interface {
	Init() error
	Serialize() (map[string][]byte, error)
	Deserialize(collection map[string][]byte) error
}

// Metadata is the struct for metadata
type Metadata struct {
	Created     string `json:"created"`
	LastUpdated string `json:"last_updated"`
}

// LogEntry is the struct for a log entry
type LogEntry struct {
	Timestamp   int64    `json:"timestamp"` // Unix timestamp
	Action      string   `json:"action"`
	EntityID    string   `json:"entityID"`
	Description string   `json:"description"`
	Changes     []Change `json:"changes"`
}

// Change is the struct for a change
type Change struct {
	Field string `json:"field"`
	Old   string `json:"old"`
	New   string `json:"new"`
}
