package container

import (
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

func TestContainerRegister(t *testing.T) {
	DefaultContainer.Register((*Test)(nil), func() any {
		return &Test{
			Name: "test",
		}
	})

	DefaultContainer.Register(Test{}, func() any {
		return Test{
			Name: "test",
		}
	})

	DefaultContainer.Register((*I)(nil), func() any {
		return Test{
			Name: "test",
		}
	})

	v1 := DefaultContainer.Make(Test{}).(Test)
	if v1.Name != "test" {
		t.Error("test fail")
	}

	v2 := DefaultContainer.Make((*Test)(nil)).(*Test)
	if v2.Name != "test" {
		t.Error("test fail")
	}

	v3 := DefaultContainer.Make((*I)(nil)).(I)
	if v3.Key() != "test" {
		t.Error("test fail")
	}
}

func TestRegister(t *testing.T) {
	Register[Test](func() any {
		return Test{
			Name: "test",
		}
	})

	Register[*Test](func() any {
		return &Test{
			Name: "test",
		}
	})

	Register[I](func() any {
		return Test{
			Name: "test",
		}
	})

	v1 := Make[Test]()
	if v1.Name != "test" {
		t.Error("test fail")
	}

	v2 := Make[*Test]()
	if v2.Name != "test" {
		t.Error("test fail")
	}

	v3 := Make[I]()
	if v3.Key() != "test" {
		t.Error("test fail")
	}
}

func TestNotRegister(t *testing.T) {
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
			t.Log(err)
		}
	}()

	v3 := Make[I]()
	if v3.Key() != "" {
		t.Error("test fail")
	}
}

func TestSetGet(t *testing.T) {
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
