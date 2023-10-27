package container

import "reflect"

type Factory interface {
	Resolvable(key any) bool
	Resolve(key any) any
	IsShared(key any) bool
}

var DefaultFactoryInstance = DefaultFactory{}

type DefaultFactory struct {
}

func (d DefaultFactory) Resolvable(key any) bool {
	return true
}

func (d DefaultFactory) IsShared(key any) bool {
	return false
}

func (d DefaultFactory) Resolve(key any) any {
	var typ reflect.Type

	if of, ok := key.(reflect.Type); ok {
		typ = of
	} else {
		typ = reflect.TypeOf(key)
	}

	if typ == nil {
		return nil
	}

	if typ.Kind() == reflect.Ptr {
		return reflect.New(typ.Elem()).Interface()
	} else {
		return reflect.New(typ).Elem().Interface()
	}
}
