package service

import "sync"

//AliasesMap store aliases of other service
type AliasesMap struct {
	items map[string]string
	sync.RWMutex
}

//Set alias
func (m *AliasesMap) Set(item string, target string) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = target
}

//Get alias
func (m *AliasesMap) Get(item string) (target string, status bool) {
	m.RLock()
	defer m.RUnlock()
	target, status = m.items[item]
	return
}

//Remove alias
func (m *AliasesMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}

//Has alias
func (m *AliasesMap) Has(item string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.items[item]
	return ok
}
