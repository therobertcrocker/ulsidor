package storage

import (
	"encoding/json"
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/domain"
	"github.com/therobertcrocker/ulsidor/utils"
	"go.etcd.io/bbolt"
)

var (
	// ensure that BboltStorageManager implements StorageManager
	_ domain.StorageManager = (*BboltStorageManager)(nil)
	// top level buckets
	buckets = []string{"codices"}
)

// BBoltStorageManager is a storage manager that uses bbolt.
type BboltStorageManager struct {
	db *bbolt.DB
}

// NewBboltStorageManager returns a new instance of BboltStorageManager.
func NewBboltStorageManager(dbPath string) (*BboltStorageManager, error) {

	// open the database
	utils.Log.Debug("creating bbolt storage manager")
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to open bbolt database: %w", err)
	}

	// create top level bucket
	utils.Log.Debug("creating top level buckets")
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, bucket := range buckets {
			_, err := tx.CreateBucketIfNotExists([]byte(bucket))
			if err != nil {
				return fmt.Errorf("failed to create bucket: %w", err)
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to create top level buckets: %w", err)
	}

	return &BboltStorageManager{db: db}, nil
}

// LoadCodex loads a codex from storage.
func (bsm *BboltStorageManager) LoadCodex(id string, codex interface{}) error {
	var data []byte
	//load codeices bucket
	utils.Log.Debug("loading codices bucket")
	err := bsm.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("codices"))
		if b == nil {
			return fmt.Errorf("bucket not found, it should have been created")
		}
		data = b.Get([]byte(id))
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to load data: %w", err)
	}

	if data == nil {
		return nil
	}

	utils.Log.Debug("deserializing data")

	err = json.Unmarshal(data, &codex)
	if err != nil {
		return fmt.Errorf("failed to deserialize data: %w", err)
	}

	utils.Log.WithField("codex", id).Debug("codex data loaded and deserialized")

	return nil

}

// SaveCodex saves a codex to storage.
func (bsm *BboltStorageManager) SaveCodex(id string, data interface{}) error {
	// serialize the data
	utils.Log.WithField("codex", id).Debug("serializing data")

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

	utils.Log.WithField("codex", id).Info("codex data saved and serialized")
	return nil
}

// Close closes the storage manager.
func (bsm *BboltStorageManager) Close() error {
	utils.Log.Debug("closing bbolt storage manager")
	return bsm.db.Close()
}
