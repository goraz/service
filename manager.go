package service

import "errors"

// Service Manager
type Manager struct {
	AllowOverride  bool
	ShareByDefault bool

	factories FactoriesMap
	services  ServicesMap
	shared    SharedMap
}

//Check for exists service
func (sm *Manager) Has(name string) (string, bool) {

	if sm.factories.Has(name) {
		return "factories", true
	}

	return "", false
}

//Un Register service
func (sm *Manager) UnRegister(name string) {
	sm.factories.Remove(name)
	sm.shared.Remove(name)
	sm.services.Remove(name)
}

// Register service as factory
func (sm *Manager) SetFacgtory(name string, fn func(*Manager) interface{}) error {

	if _, find := sm.Has(name); find {

		if sm.AllowOverride == false {
			return errors.New("A service by the name " + name + " already exists and cannot be overridden, please use an alternate name")
		}
		sm.UnRegister(name)
	}

	sm.factories.Set(name, fn)
	sm.shared.Set(name, sm.ShareByDefault)

	return nil

}

//Retrieve a registered instance
func (sm *Manager) Get(name string) (service interface{}, err error) {

	if se, found := sm.services.Get(name); found {

		service = se
		return
	}

	if factory, found := sm.factories.Get(name); found {
		service = factory(sm)
	} else {
		err = errors.New("unable to fetch or create an instance for " + name)
		return
	}

	if sm.shared.Get(name) == true {

		sm.services.Set(name, service)
	}

	return
}

//Create New Manager Struct
func NewManager() *Manager {
	return &Manager{
		ShareByDefault: true,
		shared:         SharedMap{items: make(map[string]bool)},
		factories:      FactoriesMap{items: make(map[string]func(*Manager) interface{})},
		services:       ServicesMap{items: make(map[string]interface{})},
	}
}
