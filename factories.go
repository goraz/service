package service

import "sync"

type FactoriesMap struct {
	items map[string]func(*Manager) interface{}
	sync.RWMutex
}

func (m *FactoriesMap) Set(item string, fn func(*Manager) interface{}) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = fn
}

func (m *FactoriesMap) Get(item string) (fn func(*Manager) interface{}, status bool) {
	m.RLock()
	defer m.RUnlock()
	fn, status = m.items[item]
	return
}

func (m *FactoriesMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}

func (m *FactoriesMap) Has(item string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.items[item]
	return ok
}
