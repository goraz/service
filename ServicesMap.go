package service

import "sync"

//Store service
type ServicesMap struct {
	items map[string]interface{}
	sync.RWMutex
}

func (m *ServicesMap) Set(item string, service interface{}) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = service
}

func (m *ServicesMap) Get(item string) (service interface{}, status bool) {
	m.RLock()
	defer m.RUnlock()
	service, status = m.items[item]
	return
}

func (m *ServicesMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}
