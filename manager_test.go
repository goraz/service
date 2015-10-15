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
		serviceManager := NewManager()
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
		serviceManager := NewManager()
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
		serviceManager := NewManager()
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

	Convey("not shared service", t, func() {
		serviceManager := NewManager()
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
		serviceManager := NewManager()
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

	Convey("check exists service", t, func() {
		serviceManager := NewManager()
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
		serviceManager := NewManager()
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
		serviceManager := NewManager()
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

		serviceManager.addInitializer(func(service interface{}) {
			if reflect.TypeOf(service) == reflect.TypeOf(&YourStruct{}) {
				yourStruct := service.(*YourStruct)
				yourStruct.Name += " World"

			}
		}, 2)

		serviceManager.addInitializer(func(service interface{}) {
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

}

func TestFailService(t *testing.T) {
	Convey("duplicate service", t, func() {
		serviceManager := NewManager()
		serviceManager.SetFacgtory("service1", func(sm *Manager) interface{} {
			return "baz"
		})

		err := serviceManager.SetFacgtory("service1", func(sm *Manager) interface{} {
			return "foo"
		})

		So(err, ShouldNotBeNil)
	})

	Convey("not found service", t, func() {
		serviceManager := NewManager()
		_, err := serviceManager.Get("MyStruct")
		So(err, ShouldNotBeNil)
	})

	Convey("not found service", t, func() {
		serviceManager := NewManager()
		_, err := serviceManager.Get("MyStruct")
		So(err, ShouldNotBeNil)
	})

}
