package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/vpavlin/simple-p2p-caching-proxy/internal/store"
)

type DB struct {
	db *leveldb.DB
}

func NewDB(path string) (store.IStore, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return nil, err
	}

	return &DB{
		db: db,
	}, nil
}

func (db *DB) Get(key []byte) ([]byte, error) {
	data, err := db.db.Get(key, nil)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *DB) Put(key []byte, value []byte) error {
	return db.db.Put(key, value, nil)
}

func (db *DB) Has(key []byte) (bool, error) {
	return db.db.Has(key, nil)
}
