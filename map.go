package gsync

import "sync"

// sync.Map but with generics.
type Map[K any, V any] struct {
	m      sync.Map
	initV  sync.Once
	emptyV V
}

func (m *Map[K, V]) Delete(key K) {
	m.m.Delete(key)
}

func (m *Map[K, V]) Load(key K) (value V, ok bool) {
	v, ok := m.m.Load(key)
	if !ok {
		m.initV.Do(func() {
			m.emptyV = *new(V)
		})
		return m.emptyV, ok
	}
	return v.(V), ok
}

func (m *Map[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.m.LoadAndDelete(key)
	if !loaded {
		m.initV.Do(func() {
			m.emptyV = *new(V)
		})
		return m.emptyV, loaded
	}
	return v.(V), loaded
}

func (m *Map[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.m.LoadOrStore(key, value)
	return v.(V), loaded
}

func (m *Map[K, V]) Range(f func(key K, value V) bool) {
	m.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func (m *Map[K, V]) Store(key K, value V) {
	m.m.Store(key, value)
}

// Map but with the CompareAnd* functions.
type ComparableMap[K any, V comparable] struct {
	Map[K, V]
}

func (c *ComparableMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return c.m.CompareAndDelete(key, old)
}

func (c *ComparableMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return c.m.CompareAndSwap(key, old, new)
}
