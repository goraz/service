package service

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
)

type MyStruct struct {
	Name string
}

type YourStruct struct {
	Name string
}

func TestService(t *testing.T) {
	Convey("get service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("MyStruct", func(sm *Manager) interface{} {
			return &MyStruct{
				Name: "foo",
			}
		})
		service, _ := serviceManager.Get("MyStruct")

		myStruct := service.(*MyStruct)
		So(myStruct.Name, ShouldEqual, "foo")
	})

	Convey("shared service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("MyStruct", func(sm *Manager) interface{} {
			return &MyStruct{
				Name: "foo",
			}
		})

		service, _ := serviceManager.Get("MyStruct")
		myStruct := service.(*MyStruct)

		service2, _ := serviceManager.Get("MyStruct")
		myStruct2 := service2.(*MyStruct)
		myStruct2.Name = "bar"

		So(myStruct.Name, ShouldEqual, "bar")
		So(myStruct2.Name, ShouldEqual, "bar")
	})

	Convey("not shared service", t, func() {
		serviceManager := NewManager(false)
		serviceManager.SetFacgtory("MyStruct", func(sm *Manager) interface{} {
			return &MyStruct{
				Name: "foo",
			}
		})

		service, _ := serviceManager.Get("MyStruct")
		myStruct := service.(*MyStruct)

		service2, _ := serviceManager.Get("MyStruct")
		myStruct2 := service2.(*MyStruct)
		myStruct2.Name = "bar"

		So(myStruct.Name, ShouldEqual, "foo")
		So(myStruct2.Name, ShouldEqual, "bar")
	})

	Convey("not shared service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.ShareByDefault = false
		serviceManager.SetFacgtory("MyStruct", func(sm *Manager) interface{} {
			return &MyStruct{
				Name: "foo",
			}
		})

		service, _ := serviceManager.Get("MyStruct")
		myStruct := service.(*MyStruct)

		service2, _ := serviceManager.Get("MyStruct")
		myStruct2 := service2.(*MyStruct)
		myStruct2.Name = "bar"

		So(myStruct.Name, ShouldEqual, "foo")
		So(myStruct2.Name, ShouldEqual, "bar")
	})

	Convey("Allow overide  service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.AllowOverride = true
		serviceManager.SetFacgtory("service", func(sm *Manager) interface{} {
			return "foo"
		})

		serviceManager.SetFacgtory("service", func(sm *Manager) interface{} {
			return "bar"
		})

		service, _ := serviceManager.Get("service")

		So(service.(string), ShouldEqual, "bar")
	})

	Convey("Allow overide  service with alias", t, func() {
		serviceManager := NewManager(true)
		serviceManager.AllowOverride = true
		serviceManager.SetFacgtory("service", func(sm *Manager) interface{} {
			return "foo"
		})

		serviceManager.SetFacgtory("MyService", func(sm *Manager) interface{} {
			return "bar"
		})

		serviceManager.SetAlias("MyService", "service")

		service, _ := serviceManager.Get("service")

		So(service.(string), ShouldEqual, "foo")
	})

	Convey("check exists service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("service", func(sm *Manager) interface{} {
			return "foo"
		})
		location1, found1 := serviceManager.Has("service")
		location2, found2 := serviceManager.Has("service2")
		So(found1, ShouldEqual, true)
		So(found2, ShouldEqual, false)
		So(location1, ShouldEqual, "factories")
		So(location2, ShouldEqual, "")
	})

	Convey("check already init", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("service", func(sm *Manager) interface{} {
			return "foo"
		})
		location1, found1 := serviceManager.Has("service")
		location2, found2 := serviceManager.Has("service2")
		So(found1, ShouldEqual, true)
		So(found2, ShouldEqual, false)
		So(location1, ShouldEqual, "factories")
		So(location2, ShouldEqual, "")
	})

	Convey("Initializers", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("MyStruct", func(sm *Manager) interface{} {
			return &MyStruct{
				Name: "foo",
			}
		})

		serviceManager.SetFacgtory("YourStruct", func(sm *Manager) interface{} {
			return &YourStruct{
				Name: "foo",
			}
		})

		serviceManager.AddInitializer(func(service interface{}) {
			if reflect.TypeOf(service) == reflect.TypeOf(&YourStruct{}) {
				yourStruct := service.(*YourStruct)
				yourStruct.Name += " World"

			}
		}, 2)

		serviceManager.AddInitializer(func(service interface{}) {
			if reflect.TypeOf(service) == reflect.TypeOf(&YourStruct{}) {
				yourStruct := service.(*YourStruct)
				yourStruct.Name = "Hello"

			}
		}, 1)

		service, _ := serviceManager.Get("MyStruct")
		myStruct := service.(*MyStruct)

		service2, _ := serviceManager.Get("YourStruct")
		yourStruct := service2.(*YourStruct)

		So(myStruct.Name, ShouldEqual, "foo")
		So(yourStruct.Name, ShouldEqual, "Hello World")
	})

	Convey("test alias", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetAlias("MyAlias", "MyStruct")
		serviceManager.SetFacgtory("MyStruct", func(sm *Manager) interface{} {
			return &MyStruct{
				Name: "foo",
			}
		})

		service, _ := serviceManager.Get("MyStruct")
		myStruct := service.(*MyStruct)

		service2, _ := serviceManager.Get("MyAlias")
		myStruct2 := service2.(*MyStruct)
		myStruct2.Name = "bar"

		So(myStruct.Name, ShouldEqual, "bar")
		So(myStruct2.Name, ShouldEqual, "bar")
	})

}

func TestFailService(t *testing.T) {
	Convey("duplicate service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("service1", func(sm *Manager) interface{} {
			return "baz"
		})

		err := serviceManager.SetFacgtory("service1", func(sm *Manager) interface{} {
			return "foo"
		})

		So(err, ShouldNotBeNil)
	})

	Convey("duplicate alias with factory service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetFacgtory("service1", func(sm *Manager) interface{} {
			return "baz"
		})

		err := serviceManager.SetAlias("service1", "AnotherService")
		So(err, ShouldNotBeNil)
	})

	Convey("duplicate factory service with alias", t, func() {
		serviceManager := NewManager(true)

		serviceManager.SetAlias("service1", "AnotherService")

		err := serviceManager.SetFacgtory("service1", func(sm *Manager) interface{} {
			return "baz"
		})

		So(err, ShouldNotBeNil)
	})

	Convey("not found service", t, func() {
		serviceManager := NewManager(true)
		_, err := serviceManager.Get("MyStruct")
		So(err, ShouldNotBeNil)
	})

	Convey("not found target alias service", t, func() {
		serviceManager := NewManager(true)
		serviceManager.SetAlias("MyAlias", "AnotherService")
		_, err := serviceManager.Get("MyAlias")
		So(err, ShouldNotBeNil)
	})

}
