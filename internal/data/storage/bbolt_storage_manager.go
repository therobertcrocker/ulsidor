package storage

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain/interfaces"
	"go.etcd.io/bbolt"
)

// ensure interface is implemented
var _ interfaces.StorageManager = (*BBoltStorageManager)(nil)

type BBoltStorageManager struct {
	db *bbolt.DB
}

// NewBBoltStorageManager returns a new instance of BBoltStorageManager
func NewBBoltStorageManager(dbPath string) (*BBoltStorageManager, error) {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open bbolt database: %w", err)
	}
	return &BBoltStorageManager{db: db}, nil
}

// LoadRepo loads data from storage into a repository.
func (bsm *BBoltStorageManager) LoadRepo(id string, repo interfaces.Repository) error {
	var data []byte
	err := bsm.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("entities"))
		if bucket == nil {
			return fmt.Errorf("bucket not found")
		}
		data = bucket.Get([]byte(id))
		return nil
	})
	if err != nil {
		return err
	}
	return repo.Deserialize(data)
}

// SaveRepo saves data from a repository to storage.
func (bsm *BBoltStorageManager) SaveRepo(id string, repo interfaces.Repository) error {
	data, err := repo.Serialize()
	if err != nil {
		return err
	}
	return bsm.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("entities"))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(id), data)
	})
}

// Close closes the database.
func (bsm *BBoltStorageManager) Close() error {
	return bsm.db.Close()
}
