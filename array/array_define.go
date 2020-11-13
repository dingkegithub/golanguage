package array

import "fmt"

func Define() {

}

//
// go 里面数组不同于 c 语言数组
// 指向首地址，而代表的是整个
// 数组值，当赋值时，将会整个
// 拷贝
func ArrayRefernce() {
	a := [...]int{1, 2, 3}

	// b 将完全拷贝a
	b := a

	b[0] = 5

	if b[0] == a[0] {
		fmt.Println("数组是引用类型")
	} else {
		fmt.Println("数组是值类型类型")
	}

	// 为了引用数组必须用地址
	c := &a
	c[0] = 3
	if c[0] == a[0] {
		fmt.Println("指针引用数组")
	}
}

//
// 即便指针引用数组，也不是很灵活，限定死了
// 大小和类型，但是指针好处在于引用，减少
// 整个数组拷贝，此时作为参数内部修改反馈到
// 外边
//
//
// 指针对数组的引用，使用时和直接使用数组
// 没什么区别, 可以for-range, 也可以len()
// 然后for-i 形式访问
//
func ArrayAsParam(p *[3]int) {
	fmt.Print("数组指针for-range: ")
	for _, v := range p {
		fmt.Print(v, "\t")
	}
	fmt.Println()

	fmt.Print("数组指针for-i: ")
	for i := 0; i < len(p); i++ {
		fmt.Print(p[i], "\t")
	}
	fmt.Println()

	// 内部修改反馈到外面
	p[0] = 3
}
