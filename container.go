package container

import (
	"reflect"
)

type Container struct {
	instances map[reflect.Type]any
	registers map[reflect.Type]any
}

var DefaultContainer = NewContainer()

func NewContainer() *Container {
	return &Container{
		instances: make(map[reflect.Type]any),
		registers: make(map[reflect.Type]any),
	}
}

func (c *Container) Register(key any, value any) {
	c.registers[bindKey(key)] = value
}

func (c *Container) Make(key any) any {
	if a, ok := c.registers[bindKey(key)]; ok {
		if b, ok := a.(func() any); ok {
			return b()
		}
	} else {
		if of, ok := key.(reflect.Type); ok {
			if of.Kind() == reflect.Ptr {
				return reflect.New(of.Elem()).Interface()
			} else {
				return reflect.New(of).Elem().Interface()
			}
		} else {
			return reflect.New(reflect.TypeOf(key)).Elem().Interface()
		}
	}

	return nil
}

func (c *Container) Set(key any, ins any) {
	c.instances[bindKey(key)] = ins
}

func (c *Container) Get(key any) any {
	return c.instances[bindKey(key)]
}

func (c *Container) Clear() {
	c.instances = make(map[reflect.Type]any)
	c.registers = make(map[reflect.Type]any)
}

func bindKey(key any) reflect.Type {
	var of reflect.Type

	// if key is reflect.Type
	if t, ok := key.(reflect.Type); ok {
		of = t
	} else {
		of = reflect.TypeOf(key)
	}

	return of
}

func Set[T any](ins T) {
	var t T

	DefaultContainer.Set(reflect.TypeOf(t), ins)
}

func Get[T any]() T {
	var t T
	return DefaultContainer.Get(reflect.TypeOf(t)).(T)
}

func Make[T any]() T {
	var t T

	return DefaultContainer.Make(reflect.TypeOf(t)).(T)
}

func Register[T any](callback any) {
	var t T
	DefaultContainer.Register(reflect.TypeOf(t), callback)
}

func Clear() {
	DefaultContainer.Clear()
}
