package storage

import (
	"log"
	"time"

	badger "github.com/dgraph-io/badger/v3"
)

type LocalDB struct {
	db  *badger.DB
	ttl time.Duration
}

func NewLocalDB(ttl time.Duration) LocalDB {
	db, err := badger.Open(badger.DefaultOptions(".fasi-cache"))
	if err != nil {
		log.Fatal(err)
	}

	return LocalDB{
		db:  db,
		ttl: ttl,
	}
}

func (c *LocalDB) Get(key string) (string, error) {
	var err error
	var data []byte

	err = c.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}

		data, err = item.ValueCopy(nil)
		return err
	})

	return string(data), err
}

func (c *LocalDB) Put(key string, data string) error {
	return c.PutWithTTL(key, data, c.ttl)
}

func (c *LocalDB) PutWithTTL(key string, data string, ttl time.Duration) error {
	return c.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), []byte(data)).WithTTL(ttl)
		return txn.SetEntry(e)
	})
}

func (c *LocalDB) Close() {
	c.db.Close()
}
