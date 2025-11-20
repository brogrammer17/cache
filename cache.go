package cache

import (
	"sync"
	"time"
)

type Cache interface {
	Set(key string, value any, ttl time.Duration)
	Get(key string) (any, bool)
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

func (m *MemoryCache) Set(key string, value any, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()

	expiration := time.Now().Add(ttl).Unix()
	m.items[key] = Item{value, expiration}
}

func (m *MemoryCache) Get(key string) (any, bool) {
	m.mu.RLock()
	item, exist := m.items[key]
	m.mu.RUnlock()

	if !exist {
		return nil, false
	}

	if item.Expiration > 0 && time.Now().Unix() > item.Expiration {
		m.Delete(key)
		return nil, false
	}

	return item.Value, true
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
