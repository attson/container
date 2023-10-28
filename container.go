package container

import (
	"reflect"
)

type Container struct {
	instances map[any]any
	registers map[any]any
	factories []Factory
}

var DefaultContainer = NewContainer()

func NewContainer() *Container {
	return &Container{
		instances: make(map[any]any),
		registers: make(map[any]any),
	}
}

func (c *Container) AddFactory(factory Factory) {
	c.factories = append(c.factories, factory)
}

func (c *Container) Register(key any, value any) {
	c.registers[key] = value
}

func (c *Container) Make(key any) any {
	if a, ok := c.registers[key]; ok {
		of := reflect.ValueOf(a)
		if of.Kind() == reflect.Func {
			return of.Call(nil)[0].Interface()
		}
	} else {
		for _, factory := range c.factories {
			if factory.Resolvable(key) {
				v := factory.Resolve(key)
				if factory.IsShared(key) {
					c.instances[key] = v
				}

				return v
			}
		}

		return DefaultFactoryInstance.Resolve(key)
	}

	return nil
}

func (c *Container) Set(key any, ins any) {
	c.instances[key] = ins
}

func (c *Container) Get(key any) any {
	return c.instances[key]
}

func (c *Container) Has(key any) bool {
	_, ok := c.instances[key]

	return ok
}

func (c *Container) RegisteredKeys() []string {
	var keys []string

	for k, _ := range c.registers {
		keys = append(keys, reflect.TypeOf(k).String())
	}

	return keys
}

func (c *Container) Clear() {
	c.instances = make(map[any]any)
	c.registers = make(map[any]any)
}

func Set[T any](ins T) {
	var t *T

	DefaultContainer.Set(reflect.TypeOf(t).Elem(), ins)
}

func Get[T any]() T {
	var t *T

	return DefaultContainer.Get(reflect.TypeOf(t).Elem()).(T)
}

func Has[T any]() bool {
	var t *T

	return DefaultContainer.Has(reflect.TypeOf(t).Elem())
}

func Make[T any]() T {
	var t *T

	return DefaultContainer.Make(reflect.TypeOf(t).Elem()).(T)
}

func Register[T any](callback any) {
	var t *T

	DefaultContainer.Register(reflect.TypeOf(t).Elem(), callback)
}

func RegisteredKeys() []string {
	return DefaultContainer.RegisteredKeys()
}

func RegisterK(key string, callback any) {
	DefaultContainer.Register(key, callback)
}

func SetK(key string, ins any) {
	DefaultContainer.Set(key, ins)
}

func MakeK[T any](key string) T {
	return DefaultContainer.Make(key).(T)
}

func GetK[T any](key string) T {
	return DefaultContainer.Get(key).(T)
}

func Clear() {
	DefaultContainer.Clear()
}

func AddFactory(factory Factory) {
	DefaultContainer.AddFactory(factory)
}
