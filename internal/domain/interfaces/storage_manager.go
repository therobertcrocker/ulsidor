package interfaces

// StorageManager is the interface for the storage manager
//
//go:generate mockgen -source=storage_manager.go -destination=mocks/mock_storage_manager.go -package=mocks
type StorageManager interface {
	// LoadRepo loads a repo from storage
	LoadRepo(id string, repo Repository) error
	// SaveRepo saves a repo to storage
	SaveRepo(id string, repo Repository) error
}
