package service

import "sync"

//store status of services that shared or not
type SharedMap struct {
	items map[string]bool
	sync.RWMutex
}

func (m *SharedMap) Set(item string, status bool) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = status
}

func (m *SharedMap) Get(item string) bool {
	m.RLock()
	defer m.RUnlock()
	return m.items[item]
}

func (m *SharedMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}
