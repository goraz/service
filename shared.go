package service

import "sync"

//SharedMap store status of services that shared or not
type SharedMap struct {
	items map[string]bool
	sync.RWMutex
}

//Set shared status
func (m *SharedMap) Set(item string, status bool) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = status
}

//Get shared status
func (m *SharedMap) Get(item string) bool {
	m.RLock()
	defer m.RUnlock()
	return m.items[item]
}

//Remove shared status
func (m *SharedMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}
