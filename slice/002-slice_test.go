package slice

import (
	"testing"
)

func arrayAsFuncParam(a [3]int) {
	a[0] += 1
}

func sliceAsFuncParam(a []int) {
	a[0] += 1
}

func sliceAsFuncParam2(a []int) {
	a = append(a, 10, 9, 8, 7, 6, 5, 4)
}

//
// 1. 数组作为参数是传值而非引用
//
func TestArrayValueType(t *testing.T) {
	a := [3]int{1, 2, 3}
	arrayAsFuncParam(a)

	if a[0] != 1 {
		t.Error("want 1, but: ", a[0])
		t.FailNow()
	}
}

//
// 2. 切片作为参数是引用
//
func TestSliceQuote(t *testing.T) {
	a := []int{1, 2, 3}
	sliceAsFuncParam(a)

	if a[0] == 1 {
		t.Error("want: 2, but: ", a[0])
		t.FailNow()
	}
}

//
// 3. 切片作为参数是引用,但是如果内部进行了
//    越界操作将导致引用发生变化，不再是原
//    数组的引用
//
func TestSliceQuote2(t *testing.T) {
	a := []int{1, 2, 3}
	sliceAsFuncParam2(a)

	t.Log("a[0]=", a[0])

	if a[0] != 1 {
		t.Error("want: 1, but 1")
		t.FailNow()
	}
}
