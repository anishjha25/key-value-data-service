package database

import (
	"fmt"
	"sync"
)

var (
	memBD *memoryDB
	once  sync.Once
)

type memoryDB struct {
	store map[string]string
	lock  sync.Mutex
}

func (m *memoryDB) Set(key, value string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.store[key] = value
	return nil
}

func (m *memoryDB) Get(key string) (string, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	value, ok := m.store[key]
	if !ok {
		return "", fmt.Errorf("not found")
	}
	return value, nil
}

func (m *memoryDB) Delete(key string) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.store, key)
	return nil
}

func newMemoryDB() *memoryDB {
	once.Do(func() {
		memBD = &memoryDB{
			store: map[string]string{},
		}
	})

	return memBD
}
