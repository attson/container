package container

import (
	"fmt"
	"golang.org/x/exp/slices"
	"testing"
)

type Test struct {
	Name string
}

func (t Test) Key() string {
	return t.Name
}

type I interface {
	Key() string
}

type II interface {
	Key() string
}

func testSetup() {
	Clear()
}

func TestContainerRegister(t *testing.T) {
	testSetup()
	DefaultContainer.Register((*Test)(nil), func() any {
		return &Test{
			Name: "test_pointer",
		}
	})

	DefaultContainer.Register(Test{}, func() any {
		return Test{
			Name: "test_struct",
		}
	})

	DefaultContainer.Register((*I)(nil), func() any {
		return Test{
			Name: "test_interface",
		}
	})

	v1 := DefaultContainer.Make(Test{}).(Test)
	if v1.Name != "test_struct" {
		t.Error("test fail")
	}

	v2 := DefaultContainer.Make((*Test)(nil)).(*Test)
	if v2.Name != "test_pointer" {
		t.Error("test fail")
	}

	v3 := DefaultContainer.Make((*I)(nil)).(I)
	if v3.Key() != "test_interface" {
		t.Error("test fail")
	}
}

func TestRegisterFuncType(t *testing.T) {
	testSetup()
	Register[Test](func() Test {
		return Test{
			Name: "test_struct",
		}
	})

	Register[*Test](func() *Test {
		return &Test{
			Name: "test_pointer",
		}
	})

	Register[I](func() I {
		return Test{
			Name: "test_interface",
		}
	})

	v1 := Make[Test]()
	if v1.Name != "test_struct" {
		t.Error("test fail")
	}

	v2 := Make[*Test]()
	if v2.Name != "test_pointer" {
		t.Error("test fail")
	}

	v3 := Make[I]()
	if v3.Key() != "test_interface" {
		t.Error("test fail")
	}
}

func TestRegister(t *testing.T) {
	testSetup()
	Register[Test](func() any {
		return Test{
			Name: "test_struct",
		}
	})

	Register[*Test](func() any {
		return &Test{
			Name: "test_pointer",
		}
	})

	Register[I](func() any {
		return Test{
			Name: "test_interface",
		}
	})

	v1 := Make[Test]()
	if v1.Name != "test_struct" {
		t.Error("test fail")
	}

	v2 := Make[*Test]()
	if v2.Name != "test_pointer" {
		t.Error("test fail")
	}

	v3 := Make[I]()
	if v3.Key() != "test_interface" {
		t.Error("test fail")
	}
}

func TestNotRegister(t *testing.T) {
	testSetup()
	v1 := Make[*Test]()
	if v1.Name != "" {
		t.Error("test fail")
	}

	v2 := Make[Test]()
	if v2.Name != "" {
		t.Error("test fail")
	}

	// recover()
	defer func() {
		if err := recover(); err != nil {
			e := fmt.Sprintf("%s", err)
			if e != "interface conversion: interface is nil, not container.I" {
				t.Error("test fail: " + e)
			}
		}
	}()

	_ = Make[I]()
}

func TestSetGet(t *testing.T) {
	testSetup()
	Set[Test](Test{
		Name: "set1",
	})
	Set[*Test](&Test{
		Name: "set2",
	})
	Set[I](Test{
		Name: "set3",
	})

	v1 := Get[Test]()
	if v1.Name != "set1" {
		t.Error("test fail")
	}

	v2 := Get[*Test]()
	if v2.Name != "set2" {
		t.Error("test fail")
	}

	v3 := Get[I]()
	if v3.Key() != "set3" {
		t.Error("test fail")
	}
}

func TestHas(t *testing.T) {
	testSetup()
	Set[Test](Test{
		Name: "set1",
	})
	Set[*Test](&Test{
		Name: "set2",
	})
	Set[I](Test{
		Name: "set3",
	})

	if !Has[Test]() {
		t.Error("test fail")
	}

	if !Has[*Test]() {
		t.Error("test fail")
	}

	if !Has[I]() {
		t.Error("test fail")
	}
}

func TestRegisteredKeys(t *testing.T) {
	testSetup()
	DefaultContainer.Register((*Test)(nil), func() any {
		return &Test{
			Name: "test_pointer",
		}
	})

	DefaultContainer.Register(Test{}, func() any {
		return Test{
			Name: "test_struct",
		}
	})

	DefaultContainer.Register((*I)(nil), func() any {
		return Test{
			Name: "test_interface",
		}
	})
	keys := DefaultContainer.RegisteredKeys()

	for _, s := range []string{"*container.Test", "container.Test", "*container.I"} {
		if !slices.Contains(keys, s) {
			t.Error("keys is not contain: " + s)
		}
	}
}

func TestRegisterWithKey(t *testing.T) {
	testSetup()

	SetK("test1", Test{
		Name: "set1",
	})
	k := GetK[Test]("test1")
	if k.Name != "set1" {
		t.Error("test fail")
	}

	RegisterK("test2", func() any {
		return Test{
			Name: "test_struct",
		}
	})

	v1 := MakeK[Test]("test2")
	if v1.Name != "test_struct" {
		t.Error("test fail")
	}
}

func TestRegisterInterfaces(t *testing.T) {
	Set[I](Test{
		Name: "set1",
	})

	Set[II](Test{
		Name: "set2",
	})

	v1 := Get[I]()
	if v1.Key() != "set1" {
		t.Error("test fail")
	}

	v2 := Get[II]()
	if v2.Key() != "set2" {
		t.Error("test fail")
	}
}

func TestMakeRegistered(t *testing.T) {
	testSetup()
	Set[I](Test{
		Name: "set1",
	})

	v1 := Make[I]()
	if v1.Key() != "set1" {
		t.Error("test fail")
	}
}
