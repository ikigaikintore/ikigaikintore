package storage

import "sync"

var GlobalCache cache

type cache struct {
	data map[string]any
	mtx  *sync.RWMutex
}

type Cache[T any] interface {
	Get(string) (T, bool)
	Set(string, T)
}

func NewCache() Cache[any] {
	return &cache{
		data: make(map[string]any),
		mtx:  new(sync.RWMutex),
	}
}

func (c *cache) Get(key string) (any, bool) {
	if v, hit := c.data[key]; hit {
		return v, true
	}
	return nil, false
}

func (c *cache) Set(key string, value any) {
	c.mtx.Lock()
	c.data[key] = value
	c.mtx.Unlock()
}
