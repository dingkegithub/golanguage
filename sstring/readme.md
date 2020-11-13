


> 字符串特性

- 字符串内部元素不可修改

```
a := "hello"

b := "world"

// 会提示 cannot assign to a[0]
a[0] = b[0]
```

- 字符串长度固定，但长度不是字符串类型的一部分

- 字符串的底层结构, 字符串底层来说就是一个结构体

字符串赋值操作是 reflect.StringHeader 结构体的复制过程，但并不会涉及

底层字节数组的复制

```
type StringHeader struct {
    Data uintptr  // 底层字节数组
    Len int       // 字符串的字节长度
}
```

- 求解字符串长度两种方法

```

s = "hello"

// len 求解
l1 := len(s)

// 直接访问底层数据结构
l2 := (*reflect.StringHeader)(unsafe.Pointer(&s)).Len
```

- string 支持切片操作,切片访问同一块内存数据

- 因为字符串是只读的，相同的字符串面值常量通常是对应同一个字符串常量

