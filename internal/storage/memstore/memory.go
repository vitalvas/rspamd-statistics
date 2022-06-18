package memstore

import "sync"

type MemoryStorage struct {
	sync.RWMutex
	items map[uint64]bool
}

func NewStorage() *MemoryStorage {
	return &MemoryStorage{
		items: make(map[uint64]bool),
	}
}

func (ms *MemoryStorage) Find(items []uint64) (found []uint64, err error) {
	ms.RLock()
	defer ms.RUnlock()

	for _, item := range items {
		if ok := ms.items[item]; ok {
			found = append(found, item)
		}
	}

	return
}

func (ms *MemoryStorage) Add(items []uint64) error {
	ms.Lock()
	defer ms.Unlock()

	for _, item := range items {
		ms.items[item] = true
	}

	return nil
}

func (ms *MemoryStorage) Delete(items []uint64) error {
	ms.Lock()
	defer ms.Unlock()

	for _, item := range items {
		delete(ms.items, item)
	}

	return nil
}
