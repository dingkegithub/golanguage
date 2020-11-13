package main

import (
	"fmt"
	"reflect"
	"time"
)

// https://colobu.com/2018/02/27/go-addressable/

func main() {
	checkCanAddr()
}

type S struct {
	X int
	Y string
	z int
}

func M() int {
	return 100
}

var x0 = 0

func checkCanAddr() {
	// 1. 包级别的变量: 不可寻址
	// x0: 0 	can be addressable and set: false, false
	v := reflect.ValueOf(x0)
	fmt.Printf("x0: %v \tcan be addressable and set: %t, %t\n", x0, v.CanAddr(), v.CanSet()) //false,false

	// 2. 函数级别的变量: 不可寻址
	// x1: 1 	can be addressable and set: false, false
	var x1 = 1
	v = reflect.Indirect(reflect.ValueOf(x1))
	fmt.Printf("x1: %v \tcan be addressable and set: %t, %t\n", x1, v.CanAddr(), v.CanSet()) //false,false

	// 3. 函数级别的宾亮引用：可寻址
	// x2: 0xc0000160c0 	can be addressable and set: true, true
	var x2 = &x1
	v = reflect.Indirect(reflect.ValueOf(x2))
	fmt.Printf("x2: %v \tcan be addressable and set: %t, %t\n", x2, v.CanAddr(), v.CanSet()) //true,true

	// 4. 函数返回的结构体: 可寻址
	// x3: 2020-11-04 20:36:22.202639 +0800 CST m=+0.000155475 	can be addressable and set: false, false
	var x3 = time.Now()
	v = reflect.Indirect(reflect.ValueOf(x3))
	fmt.Printf("x3: %v \tcan be addressable and set: %t, %t\n", x3, v.CanAddr(), v.CanSet()) //false,false

	// 5. 指针变量指向的值时可寻址的：指针指向的是地址，当然可寻址
	// x4: 2020-11-04 20:36:22.202639 +0800 CST m=+0.000155475 	can be addressable and set: true, true
	var x4 = &x3
	v = reflect.Indirect(reflect.ValueOf(x4))
	fmt.Printf("x4: %v \tcan be addressable and set: %t, %t\n", x4, v.CanAddr(), v.CanSet()) // true,true

	// 6. 切片不可寻址
	// x5: [1 2 3] 	can be addressable and set: false, false
	var x5 = []int{1, 2, 3}
	v = reflect.ValueOf(x5)
	fmt.Printf("x5: %v \tcan be addressable and set: %t, %t\n", x5, v.CanAddr(), v.CanSet()) // false,false

	// 7. 切片元素值不可寻址
	// x6: 1 	can be addressable and set: false, false
	var x6 = []int{1, 2, 3}
	v = reflect.ValueOf(x6[0])
	fmt.Printf("x6: %v \tcan be addressable and set: %t, %t\n", x6[0], v.CanAddr(), v.CanSet()) //false,false

	// 8.
	// x7: 1 	can be addressable and set: true, true
	var x7 = []int{1, 2, 3}
	v = reflect.ValueOf(x7).Index(0)
	fmt.Printf("x7: %v \tcan be addressable and set: %t, %t\n", x7[0], v.CanAddr(), v.CanSet()) //true,true

	// 9.
	// x7.1: 2 	can be addressable and set: false, false
	v = reflect.ValueOf(&x7[1])
	fmt.Printf("x7.1: %v \tcan be addressable and set: %t, %t\n", x7[1], v.CanAddr(), v.CanSet()) //true,true

	// 10
	// x8: 1 	can be addressable and set: false, false
	var x8 = [3]int{1, 2, 3}
	v = reflect.ValueOf(x8[0])
	fmt.Printf("x8: %v \tcan be addressable and set: %t, %t\n", x8[0], v.CanAddr(), v.CanSet()) //false,false
	// https://groups.google.com/forum/#!topic/golang-nuts/RF9zsX82MWw
	// 11.
	// x9: 1 	can be addressable and set: false, false
	var x9 = [3]int{1, 2, 3}
	v = reflect.Indirect(reflect.ValueOf(x9).Index(0))
	fmt.Printf("x9: %v \tcan be addressable and set: %t, %t\n", x9[0], v.CanAddr(), v.CanSet()) //false,false

	//
	// x9: 1 	can be addressable and set: true, true
	var x10 = [3]int{1, 2, 3}
	v = reflect.Indirect(reflect.ValueOf(&x10)).Index(0)
	fmt.Printf("x9: %v \tcan be addressable and set: %t, %t\n", x10[0], v.CanAddr(), v.CanSet()) //true,true

	// 12. 结构体不可寻址
	// x11: {0  0} 	can be addressable and set: false, false
	var x11 = S{}
	v = reflect.ValueOf(x11)
	fmt.Printf("x11: %v \tcan be addressable and set: %t, %t\n", x11, v.CanAddr(), v.CanSet()) //false,false

	// 13. 结构体引用可寻址
	//x12: {0  0} 	can be addressable and set: true, true
	var x12 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x12))
	fmt.Printf("x12: %v \tcan be addressable and set: %t, %t\n", x12, v.CanAddr(), v.CanSet()) //true,true

	// 14.
	// x13: {0  0} 	can be addressable and set: false, false
	var x13 = S{}
	v = reflect.ValueOf(x13).FieldByName("X")
	fmt.Printf("x13: %v \tcan be addressable and set: %t, %t\n", x13, v.CanAddr(), v.CanSet()) //false,false

	// 15.
	// x14: {0  0} 	can be addressable and set: true, true
	var x14 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x14)).FieldByName("X")
	fmt.Printf("x14: %v \tcan be addressable and set: %t, %t\n", x14, v.CanAddr(), v.CanSet()) //true,true

	// 16.
	// x15: {0  0} 	can be addressable and set: true, false
	var x15 = S{}
	v = reflect.Indirect(reflect.ValueOf(&x15)).FieldByName("z")
	fmt.Printf("x15: %v \tcan be addressable and set: %t, %t\n", x15, v.CanAddr(), v.CanSet()) //true,false

	// 17.
	// x15.1: &{0  0} 	can be addressable and set: true, true
	v = reflect.Indirect(reflect.ValueOf(&S{}))
	fmt.Printf("x15.1: %v \tcan be addressable and set: %t, %t\n", &S{}, v.CanAddr(), v.CanSet()) //true,true

	/*
	 */

	// 17
	// x16: 0x10a8940 	can be addressable and set: false, false
	var x16 = M
	v = reflect.ValueOf(x16)
	fmt.Printf("x16: %p \tcan be addressable and set: %t, %t\n", x16, v.CanAddr(), v.CanSet()) //false,false

	// 18
	// x17: 0x10a8940 	can be addressable and set: true, true
	var x17 = M
	v = reflect.Indirect(reflect.ValueOf(&x17))
	fmt.Printf("x17: %p \tcan be addressable and set: %t, %t\n", x17, v.CanAddr(), v.CanSet()) //true,true

	// 19.
	// x18: &{0  0} 	can be addressable and set: false, false
	var x18 interface{} = &x11
	v = reflect.ValueOf(x18)
	fmt.Printf("x18: %v \tcan be addressable and set: %t, %t\n", x18, v.CanAddr(), v.CanSet()) //false,false

	// 20.
	// x19: &{0  0} 	can be addressable and set: true, true
	var x19 interface{} = &x11
	v = reflect.ValueOf(x19).Elem()
	fmt.Printf("x19: %v \tcan be addressable and set: %t, %t\n", x19, v.CanAddr(), v.CanSet()) //true,true

	// 21
	// x20: [1 2 3] 	can be addressable and set: false, false
	var x20 = [...]int{1, 2, 3}
	v = reflect.ValueOf([...]int{1, 2, 3})
	fmt.Printf("x20: %v \tcan be addressable and set: %t, %t\n", x20, v.CanAddr(), v.CanSet()) //false,false
}
