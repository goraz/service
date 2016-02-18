package service

import "sync"

//ServicesMap Store service
type ServicesMap struct {
	items map[string]interface{}
	sync.RWMutex
}

//Set Service
func (m *ServicesMap) Set(item string, service interface{}) {
	m.Lock()
	defer m.Unlock()
	m.items[item] = service
}

//Get Service
func (m *ServicesMap) Get(item string) (service interface{}, status bool) {
	m.RLock()
	defer m.RUnlock()
	service, status = m.items[item]
	return
}

//Remove Service
func (m *ServicesMap) Remove(item string) {
	m.Lock()
	defer m.Unlock()
	delete(m.items, item)
}
