package service

import "sync"

//store aliases of other service
type AliasesMap struct {
	items map[string]string
	sync.RWMutex
}

func (m *AliasesMap) Set(item string, target string) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = target
}

func (m *AliasesMap) Get(item string) (target string, status bool) {
	m.RLock()
	defer m.RUnlock()
	target, status = m.items[item]
	return
}

func (m *AliasesMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}

func (m *AliasesMap) Has(item string) bool {
	m.RLock()
	defer m.RUnlock()
	_, ok := m.items[item]
	return ok
}
