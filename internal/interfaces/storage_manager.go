package interfaces

// StorageManager is the interface for the storage manager
type StorageManager interface {
	// LoadRepo loads a repo from storage
	LoadRepo(id string, repo Repository) error
	// SaveRepo saves a repo to storage
	SaveRepo(id string, repo Repository) error
	// LoadGameData loads game data from storage
	LoadGameData() (map[string]interface{}, error)
}
