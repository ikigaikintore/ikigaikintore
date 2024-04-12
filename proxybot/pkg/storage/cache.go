package storage

import (
	"github.com/ikigaikintore/ikigaikintore/proxybot/pkg/domain"
	"sync"
)

var GlobalCache Cache[domain.Location]

type cache[T any] struct {
	data map[string]any
	mtx  *sync.RWMutex
}

type Cache[T any] interface {
	Get(string) (T, bool)
	Set(string, T)
}

func NewCache[T any]() Cache[T] {
	return &cache[T]{
		data: make(map[string]any),
		mtx:  &sync.RWMutex{},
	}
}

func (c *cache[T]) Get(key string) (T, bool) {
	var zero T
	raw, hit := c.data[key]
	v, ok := raw.(T)
	if !ok {
		return zero, hit
	}
	return v, hit
}

func (c *cache[T]) Set(key string, value T) {
	c.mtx.Lock()
	c.data[key] = value
	c.mtx.Unlock()
}
