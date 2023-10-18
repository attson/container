# 基于 golang 泛型实现的简单容器

- 支持单例注册
- 支持通过类型构建对象

## 使用示例

```go
package main

import "github.com/attson/container"

type Test struct {
	Name string
}

func (t Test) Key() string {
	return t.Name
}

type I interface {
	Key() string
}
```

### 为结构体注册一个构建函数

```
// 为结构体注册一个构建函数
container.Register[Test](func() any {
    return Test{
        Name: "test",
    }
})
// 通过容器构建实例
v1 := container.Make[Test]()
println(v1.Name) // test
```

### 为结构体指针注册一个构建函数

```
// 为结构体指针注册一个构建函数
container.Register[*Test](func() any {
    return &Test{
        Name: "test_pointer",
    }
})
// 通过容器构建实例
v2 := container.Make[*Test]()
println(v2.Name) // test_pointer
```

### 为接口注册一个构建函数

```
// 为接口注册一个构建函数
container.Register[I](func() any {
    return Test{
        Name: "test_interface",
    }
})
// 通过容器构建实例
v3 := container.Make[I]()
println(v3.Key()) // test_interface
```
### 在容器中设置一个实例(单例)
```
// 在容器中设置一个实例(单例)
container.Set[Test](Test{
    Name: "test_set",
})
// 通过容器获取实例
v4 := container.Get[Test]()
println(v4.Name) // test_set
	
```

### 在容器中设置一个指针实例
```
// 在容器中设置一个指针实例
container.Set[*Test](&Test{
    Name: "test_pointer_set",
})
// 通过容器获取实例
v5 := container.Get[*Test]()
println(v5.Name) // test_pointer_set
```

### 在容器中设置一个接口实例

```
// 在容器中设置一个接口实例
container.Set[I](Test{
    Name: "test_interface_set",
})
// 通过容器获取实例
v6 := container.Get[I]()
println(v6.Key()) // test_interface_set
```

