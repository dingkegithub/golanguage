package slice

import "testing"

func TestSliceDefine(t *testing.T) {
	t.Log("1. go 语言中任何类型定义都将初始化为0值, 切片的0值为nil")
	// nil切片, 和 nil 相等, 一般用来表示一个不存在的切片
	var a []int
	if a != nil {
		t.Error("nil slice != nil")
		t.FailNow()
	}

	t.Log("2. 空切片表示已初始化，不再是0值，其len 和 cap 都为 0")
	// 空切片, 和 nil 不相等, 一般用来表示一个空的集合
	b := []int{}
	if b == nil {
		t.Error("blank slice == nil !!!!!")
		t.FailNow()
	}

	if len(b) != 0 {
		t.Error("blank slice len != 0 !!!!!")
		t.FailNow()
	}

	if cap(b) != 0 {
		t.Error("blank slice cap != 0 !!!!!")
		t.FailNow()
	}

	// 有3个元素的切片, len和cap都为3
	c := []int{1, 2, 3}
	if len(c) != 3 && len(c) != cap(c) {
		t.Error("c = []int{1, 2, 3} slice != cap = 3 !!!!!")
		t.FailNow()
	}

	//
	// 有2个元素的切片, len为2, cap为3
	//
	t.Log("3. 切片在引用时，其大小是实际指定大小，而容量是原切片或底层字节数组的容量减去引用的起始位置")
	d := c[:2]
	if len(d) != 2 && cap(c) != 3 {
		t.Error("d := c[:2] len!=2 or cap != 3 !!!!!")
		t.FailNow()
	}

	//
	// 有2个元素的切片, len为2, cap为3
	//
	t.Log("5. e := c[start:end:cap]")
	e := c[0:2:cap(c)]
	if len(e) != 2 && cap(e) != 3 {
		t.Error("e := c[0:2:cap(c)] len!=2 or cap != 3 !!!!!")
		t.FailNow()
	}

	// 有0个元素的切片, len为0, cap为3
	f := c[:0]
	if len(f) != 0 && cap(f) != 3 {
		t.Error("f := c[:0] len!=0 or cap != 3 !!!!!")
		t.FailNow()
	}

	// 有3个元素的切片, len和cap都为3
	t.Log("5. g := make([]int, n), 初始化一个cap 和 len 为 n 的切片, 其元素值为实际类型的0值")
	g := make([]int, 3)
	if len(g) != 3 && cap(g) != 3 {
		t.Error("g := make([]int, 3) len!=3 or cap != 3 !!!!!")
		t.FailNow()
	}

	// 有2个元素的切片, len为2, cap为3
	t.Log("7. g := make([]int, n, m), 初始化一个cap 为 m 和 len 为 n 的切片, 前n个元素值为实际类型的0值")
	h := make([]int, 2, 3)
	if len(h) != 2 && cap(h) != 3 {
		t.Error("h := make([]int, 2, 3) len!=2 or cap != 3 !!!!!")
		t.FailNow()
	}

	// 有0个元素的切片, len为0, cap为3
	i := make([]int, 0, 3)
	if len(i) != 0 && cap(i) != 3 {
		t.Error("i := make([]int, 0, 3)  len!=0 or cap != 3 !!!!!")
		t.FailNow()
	}
}
