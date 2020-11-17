


> 字符串特性

- 字符串内部元素不可修改

```
a := "hello"

b := "world"

// 会提示 cannot assign to a[0]
a[0] = b[0]
```

- 字符串长度固定，但长度不是字符串类型的一部分

- go原因字符串 utf8编码，通常被解释为utf8 编码的 Unicode 码点 rune 序列

- 字符串可以包含任何数据，包括 byte 值0

- 字符串的底层结构, 字符串底层来说就是一个结构体

字符串赋值操作是 reflect.StringHeader 结构体的复制过程，但并不会涉及

底层字节数组的复制

```
type StringHeader struct {
    Data uintptr  // 底层字节数组
    Len int       // 字符串的字节长度
}

// hello world 字符串底层数组其实和下面字节数组一致
var data = [...]byte{
    'h', 'e', 'l', 'l', 'o', ',', ' ', 'w', 'o', 'r', 'l', 'd',
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

```
s := "hello, world"
hello := s[:5]
world := s[7:]

s1 := "hello, world"[:5]
s2 := "hello, world"[7:]
```

- 因为字符串是只读的，相同的字符串面值常量通常是对应同一个字符串常量

- 字符串可以存放任意二进制字节序列，遇到坏的编码，将生产unicode字符 \uFFFD, 显示白色问好�

```
// \xe4\x00\x00 是破坏后的编码
fmt.Println("\xe4\x00\x00\xe7\x95\x8cabc") // �界abc

// 遍历时，坏的会被分开拆成三部分，合法的遍历成一个
for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
    fmt.Println(i, c)
}
// 0 65533  // \uFFFD, 对应 �
// 1 0      // 空字符
// 2 0      // 空字符
// 3 30028  // 界
// 6 97     // a
// 7 98     // b
// 8 99     // c
```

- rune 其实是int32类型，用于表示unicode 码点，目前只使用21个bit位


- string， byte， rune 相互转换

```
// to rune, byte
s := "aaaa"
bs := []byte(s)
rs := []rune(s)

// to string
sfb := string(bs)
sfr := string(rs)
```

由于rune 是int32，底层数据结构不一致，加之字符串只读，所以涉及空间重新

开辟，然后拷贝，其时间复杂度为 O(n), 而对于 []byte, 虽然字符串底层也是

[]byte 数组，但是一旦涉及修改操作可能也会重分配空间拷贝，时间复杂度仍然

是 O(n)
