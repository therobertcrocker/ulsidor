package storage

import (
	"encoding/json"
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain"
	"go.etcd.io/bbolt"
)

// Ensure implementation of the StorageManager interface.
var _ domain.StorageManager = (*BboltStorageManager)(nil)

// BBoltStorageManager is a storage manager that uses bbolt.
type BboltStorageManager struct {
	db *bbolt.DB
}

// NewBboltStorageManager returns a new instance of BboltStorageManager.
func NewBboltStorageManager(dbPath string) (*BboltStorageManager, error) {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open bbolt database: %w", err)
	}
	return &BboltStorageManager{db: db}, nil
}

// LoadCodex loads a codex from storage.
func (bsm *BboltStorageManager) LoadCodex(id string, codex interface{}) error {
	var data []byte
	err := bsm.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("codices"))
		if b == nil {
			return fmt.Errorf("bucket %q not found", "codices")
		}
		data = b.Get([]byte(id))
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to load data: %w", err)
	}

	err = json.Unmarshal(data, &codex)
	if err != nil {
		return fmt.Errorf("failed to deserialize data: %w", err)
	}

	return nil

}

// SaveCodex saves a codex to storage.
func (bsm *BboltStorageManager) SaveCodex(id string, data interface{}) error {
	// serialize the data
	serialized, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to serialize data: %w", err)
	}
	err = bsm.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("codices"))
		if err != nil {
			return fmt.Errorf("failed to create bucket: %w", err)
		}
		err = b.Put([]byte(id), serialized)
		if err != nil {
			return fmt.Errorf("failed to put data: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to save data: %w", err)
	}
	return nil
}

// Close closes the storage manager.
func (bsm *BboltStorageManager) Close() error {
	return bsm.db.Close()
}
