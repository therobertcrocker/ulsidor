package storage

import (
	"fmt"

	"github.com/therobertcrocker/ulsidor/internal/data/utils"
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
	collection := make(map[string][]byte)
	err := bsm.db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(id))
		if bucket == nil {
			utils.Log.Debugf("Bucket %s not found, either collection is empty or other error", id)
			return nil
		}
		collection["data"] = bucket.Get([]byte("data"))
		collection["metadata"] = bucket.Get([]byte("metadata"))
		collection["changelog"] = bucket.Get([]byte("changelog"))
		return nil
	})
	if err != nil {
		return err
	}
	return repo.Deserialize(collection)
}

// SaveRepo saves data from a repository to storage.
func (bsm *BBoltStorageManager) SaveRepo(id string, repo interfaces.Repository) error {
	collection, err := repo.Serialize()
	if err != nil {
		return err
	}
	err = bsm.db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(id))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("data"), collection["data"])
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("metadata"), collection["metadata"])
		if err != nil {
			return err
		}
		err = bucket.Put([]byte("changelog"), collection["changelog"])
		if err != nil {
			return err
		}
		return nil
	},
	)
	if err != nil {
		return err
	}
	return nil
}

// Close closes the database.
func (bsm *BBoltStorageManager) Close() error {
	return bsm.db.Close()
}
