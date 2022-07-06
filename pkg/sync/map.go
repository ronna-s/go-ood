package sync

import "sync"

type Map[K comparable, V any] struct {
	M  map[K]V
	RW sync.RWMutex
}

func (m *Map[K, V]) Get(key K) V {
	m.RW.RLock()
	defer m.RW.RUnlock()
	return m.M[key]
}

func (m *Map[K, V]) Set(key K, val V) {
	m.RW.Lock()
	defer m.RW.Unlock()
	m.M[key] = val
}
