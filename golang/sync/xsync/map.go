package xsync

// a kv is a key-value pair.
type kv[K comparable, V any] struct {
	key   K
	value V
}

// a k is a key.
type k[K comparable] struct {
	key K
}

// a vOk is a value with ok.
type vOk[V any] struct {
	value V
	ok    bool
}

// Map is like a Go map but is safe for concurrent use by multiple goroutines,
// without additional locking or coordination.
//
// The zero Map is empty and ready for use. A Map must not be copied after first use.
type Map[K comparable, V any] struct {
	m            map[K]V
	addChan      chan kv[K, V]
	addRespChan  chan struct{}
	loadChan     chan k[K]
	loadRespChan chan vOk[V]
	delChan      chan k[K]
	delRespChan  chan struct{}
}

func NewMap[K comparable, V any]() *Map[K, V] {
	m := &Map[K, V]{
		m:            make(map[K]V),
		addChan:      make(chan kv[K, V]),
		addRespChan:  make(chan struct{}),
		loadChan:     make(chan k[K]),
		loadRespChan: make(chan vOk[V]),
		delChan:      make(chan k[K]),
		delRespChan:  make(chan struct{}),
	}

	go func() {
		for {
			select {
			case req := <-m.addChan:
				m.m[req.key] = req.value
				m.addRespChan <- struct{}{}
			case req := <-m.loadChan:
				value, ok := m.m[req.key]
				m.loadRespChan <- vOk[V]{value: value, ok: ok}
			case req := <-m.delChan:
				delete(m.m, req.key)
				m.delRespChan <- struct{}{}
			}
		}
	}()

	return m
}

// Add sets the value for key.
func (m *Map[K, V]) Add(key K, value V) {
	m.addChan <- kv[K, V]{
		key:   key,
		value: value,
	}
	<-m.addRespChan
}

// Load2 returns the value for key, ok indicates whether key was found.
func (m *Map[K, V]) Load2(key K) (value V, ok bool) {
	m.loadChan <- k[K]{key: key}
	loadResp := <-m.loadRespChan
	return loadResp.value, loadResp.ok
}

// Load returns the value for key.
func (m *Map[K, V]) Load(key K) (value V) {
	value, _ = m.Load2(key)
	return
}

// Delete deletes the value for key.
func (m *Map[K, V]) Delete(key K) {
	m.delChan <- k[K]{key: key}
	<-m.delRespChan
}
