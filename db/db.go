package db

import (
	"errors"
	"sync"
	"time"
)

type Record struct {
	Value      string
	Expiration int64
}

type DB struct {
	data map[string]Record
	mu   sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]Record),
	}
}

func (db *DB) Set(key, value string, ttlSeconds int64) {
	db.mu.Lock()
	defer db.mu.Unlock()

	expiration := int64(0)
	if ttlSeconds > 0 {
		expiration = time.Now().Unix() + ttlSeconds
	}

	db.data[key] = Record{
		Value:      value,
		Expiration: expiration,
	}
}

func (db *DB) Get(key string) (string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	value, ok := db.data[key]
	if !ok {
		return "", errors.New("not found")
	}
	if !(value.Expiration == 0) && time.Now().Unix() > value.Expiration {
		return "", errors.New("expired")
	}
	return value.Value, nil
}

func (db *DB) Delete(key string) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.data, key)
}

func (db *DB) Count(value string) int {
	db.mu.RLock()
	defer db.mu.RUnlock()

	count := 0
	for _, v := range db.data {
		if v.Expiration == 0 || v.Expiration > time.Now().Unix() {
			count++
		}
	}
	return count
}
