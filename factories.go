package service

import "sync"

//FactoriesMap Store Factories service
type FactoriesMap struct {
	items map[string]func(*Manager) interface{}
	sync.RWMutex
}

//Set Factory Service
func (m *FactoriesMap) Set(item string, fn func(*Manager) interface{}) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = fn
}

//Get Factory Service
func (m *FactoriesMap) Get(item string) (fn func(*Manager) interface{}, status bool) {
	m.RLock()
	defer m.RUnlock()
	fn, status = m.items[item]
	return
}

//Remove Factory Service
func (m *FactoriesMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}

//Has Factory Service
func (m *FactoriesMap) Has(item string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.items[item]
	return ok
}
