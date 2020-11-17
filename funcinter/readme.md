


<h2 id='1'> 1. 包初始化流程 </h2>

```
        mian               pkg1               pkg2

->  import pkg1 -----> import pkg2 --------> const
         |                  |                  |
         | <------+         | <--------+       |
        \|/       |        \|/         |      \|/
       const      |       const        |      var
         |        |         |          |       |
         |        |         |          |       |
        \|/       |        \|/         |      \|/
        var       |        var         |     init()
         |        |         |          |       |
         |        |         |          +-------+
        \|/       |        \|/
       init()     |       init()  它先执行
         |        |       init()  后执行
         |        |         |
         |        |         |
        \|/       +---------+ 
       main()

```

> 特性

- 调用关系：main--->pkg1--->pkg2, 初始化次序应该是 pkg2-->pkg1---->main

- 一个包内的init 将顺序调用

- 在main函数之前执行的所有代码都在同样一个 goroutine, 称为zhugoroutine

- 其它 goroutine 只有进入 main 之后才可能被执行

> 详见 pkgdaemon


<h2 id='2'> 2. 函数 </h2>

<h3 id='2.1'> 2.1 函数分类 </h3>

> 有无函数名分类

- 具名函数

```
func Add(a, b int) int {
    return a+b
}
```

- 匿名函数

```
var add = func(a, b int) int {
    return a+b
}
```

> 参数分类

- 可变参数: 可变参数必须放置最后

```
func Add(a int, b ...int) int {
    for _, v := range b {
        a += v
    }

    return a
}
```

- 不可变参数

```
func Add(a, b int) int {
    return a+b
}
```

> 返回值分类

- 非命名

```
func Add(a, b int) int {
    return a+b
}
```

- 命名

直接操作命名返回值或defer中判断一些条件操作命名

```
func Add(a, b int) (sum int) {
    sum = a + b
}
```

或者

```
func Add(a, b int) (sum int) {
    defer func() {
        sum = a+b
    }()
}
```

<h3 id='2.2'> 2.2 函数特性 </h3>

- go函数的栈动态分配，没有栈深度的限制，栈大小可以达到GB级别

- 参数和返回值都是以 **`传值`** 方式和调用者进行交换数据

传入切片好像修改值，改变了参数，似乎是引用，看代码

下面操作修改的时候似乎引用，但是一旦使用Add 之后返现并没有

添加进入


>
> 本质上说传值是传入的是本身的值，比如切片传入的是底层结构体的值

> 包含数据指针和长度，但不包含指针指向的内容，所以底层修改了len

> 和cap，指针等时不会反应到外层
>

```
// 修改操作
func Update(arr []int) {
	if len(arr) <= 0 {
		return
	}
	arr[0] = 9999
}

// 添加操作
func Add(arr []int) {
	items := []int{100, 101, 102}
	arr = append(arr, items...)
}


arr := []int{1, 2, 3}
-----------------------------------------
update: : 9999 92 93 9

add: : 9999 92 93 9
```

<h3 id='2.3'> 2.3 defer 匿名陷阱 </h3>

defer 是延迟执行, 所以下面的代码将最后输出全部为 3

```
func main() {
    for i:=0; i<3; i++ {
        defer func(){
            fmt.Printn("error: ", i)
        }()
    }
}

----------------------------------
error:  3
error:  3
error:  3
```

修正

```
func main() {
    for i:=0; i<3; i++ {
        // 传入参数defer语句马上对参数求知，保存到函数栈中
        defer func(v int){
            fmt.Printn("ok: ", v)
        }(i)
    }
}

----------------------------------
ok:  2
ok:  1
ok:  0
```

<h2 id='3'> 3. 方法 </h2>

go 语言不支持继承，但是可以通过组和匿名对象来继承，从而访问父类的方法

```
type Pool struct {
	// 组合匿名实现继承
	sync.Mutex
	total int
	idle  int
}

func (p *Pool) GetTotal() int {
	// 访问父类的方法
    // 和访问自己的方法没有区别
	p.Lock()
	defer p.Unlock()
	return p.total
}
```


<h2 id='4'> 4. 接口 </h2>

实现纯虚的关键是动态延迟绑定接口道实际类型

- 鸭子模型

接口满足鸭子模型(走起路来像鸭子,就是说实现了接口方法就行), 将实现接口的类传给接口实现多态

> Go语言的接口类型是延迟绑定，可以实现类似虚函数的多态功能

fmt.Printf 可以打印任何满足 fmt.Stringer 接口的对象

- 纯虚继承

通过嵌入匿名接口或嵌入匿名指针对象来实现继承的做法其实是一种纯虚继承，我们继承的只是接口指定的规范，真正的实现在运行的时候才被注入
