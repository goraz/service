package service

import (
	"errors"
	"sort"
)

// Manager for Service
type Manager struct {
	AllowOverride  bool
	ShareByDefault bool //Whether or not to share by default

	factories    FactoriesMap
	services     ServicesMap
	shared       SharedMap
	aliases      AliasesMap
	initializers Initializers
}

//Has Check for exists service
func (sm *Manager) Has(name string) (string, bool) {

	if sm.factories.Has(name) {
		return "factories", true
	}

	if sm.aliases.Has(name) {
		return "aliases", true
	}

	return "", false
}

//UnRegister service
func (sm *Manager) UnRegister(name string) {
	sm.factories.Remove(name)
	sm.aliases.Remove(name)
	sm.shared.Remove(name)
	sm.services.Remove(name)
}

//SetFacgtory Register service as factory
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

//SetAlias Register new alias to another service
func (sm *Manager) SetAlias(name string, target string) error {

	if _, find := sm.Has(name); find {

		if sm.AllowOverride == false {
			return errors.New("A service by the name " + name + " already exists and cannot be overridden, please use an alternate name")
		}
		sm.UnRegister(name)
	}

	sm.aliases.Set(name, target)

	return nil
}

//AddInitializer add new initializer
func (sm *Manager) AddInitializer(fn func(interface{}), order float32) {
	sm.initializers = append(sm.initializers, Initializer{
		fn:    fn,
		order: order,
	})

	sort.Sort(sm.initializers)
}

//Get service
func (sm *Manager) Get(name string) (service interface{}, err error) {

	if se, found := sm.services.Get(name); found {

		service = se
		return
	}

	if factory, found := sm.factories.Get(name); found {
		service = factory(sm)
	} else if alias, found := sm.aliases.Get(name); found {
		if service, err = sm.Get(alias); err != nil {
			return
		}

	} else {
		err = errors.New("unable to fetch or create an instance for " + name)
		return
	}
	// apply initializers
	for _, init := range sm.initializers {
		init.fn(service)
	}

	if sm.shared.Get(name) == true {

		sm.services.Set(name, service)
	}

	return
}

//NewManager Create New Manager Struct
func NewManager() *Manager {
	return &Manager{
		ShareByDefault: true,
		shared:         SharedMap{items: make(map[string]bool)},
		aliases:        AliasesMap{items: make(map[string]string)},
		factories:      FactoriesMap{items: make(map[string]func(*Manager) interface{})},
		services:       ServicesMap{items: make(map[string]interface{})},
		initializers:   Initializers{},
	}
}
