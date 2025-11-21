package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value any, ttl time.Duration) error
	Get(key string) (any, error)
	Delete(key string)
	Clear()
}

type MemoryCache struct {
	mu    sync.RWMutex
	items map[string]Item
}

func New() Cache {
	return &MemoryCache{
		items: make(map[string]Item),
	}
}

func (m *MemoryCache) Set(key string, value any, ttl time.Duration) error {
	if ttl < 0 {
		return ErrInvalidTTL
	}

	var exp int64
	if ttl > 0 {
		exp = time.Now().Add(ttl).Unix()
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.items[key] = Item{value, exp}

	return nil
}

func (m *MemoryCache) Get(key string) (any, error) {
	m.mu.RLock()
	item, exist := m.items[key]
	m.mu.RUnlock()

	if !exist {
		return nil, ErrNotFound
	}

	if item.IsExpired() {
		m.Delete(key)
		return nil, ErrExpired
	}

	return item.Value, nil
}

func (m *MemoryCache) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.items, key)
}

func (m *MemoryCache) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.items = make(map[string]Item)
}
