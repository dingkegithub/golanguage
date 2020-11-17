package slice

import (
	"reflect"
	"sort"
	"testing"
	"unsafe"
)

func TestAdd(t *testing.T) {
	// nil 切片
	var a []int

	// 容量不足，重新分配空间
	a = append(a, 1)                 // 追加1个元素
	a = append(a, 1, 2, 3)           // 追加多个元素, 手写解包方式
	a = append(a, []int{1, 2, 3}...) // 追加一个切片, 切片需要解包
}

func TestRealloc(t *testing.T) {
	// len = cap = 3
	var a = []int{1, 2, 3}

	// 添加 0， 容量不足，重新分配空间和拷贝, O(n)
	a = append([]int{0}, a...) // 在开头添加1个元素

	// 头部添加，挪动元素，复杂度 O(n)
	a = append([]int{-3, -2, -1}, a...) // 在开头添加1个切片
}

func TestAppendLink(t *testing.T) {
	var a = []int{1, 2, 3}

	// 第1个位置添加5
	// 中间 append 会临时创建切片，效率低
	a = append(a[:1], append([]int{5}, a[1:]...)...) // 在第i个位置插入x

	// 第1个位置添加切片
	a = append(a[:1], append([]int{1, 2, 3}, a[1:]...)...) // 在第i个位置插入切片
}

func TestOptiAppendLinkWithCopy(t *testing.T) {
	var a = []int{1, 2, 3}

	// 添加一个元素，先扩展空间，预留位置
	a = append(a, 0) // 切片扩展1个空间

	//
	// 从 i 处向后移动一个位置
	//
	i := 1
	copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置

	// 插入
	a[i] = 9

	x := []int{10, 11, 12}

	// 扩展空间
	a = append(a, x...)

	// i 处的元素向后挪动len(x), 空出前面的位置
	copy(a[i+len(x):], a[i:])

	// 插入切片
	copy(a, x)
}

func TestDel(t *testing.T) {
	a := []int{1, 2, 3}
	a = append(a[:0], a[1:]...) // 删除开头1个元素
	a = append(a[:0], a[2:]...) // 删除开头5个元素
}

func TestDelWithCopy(t *testing.T) {
	a := []int{1, 2, 3}
	// copy 返回的是实际拷贝的个数
	a = a[:copy(a, a[1:])] // 删除开头1个元素
	a = a[:copy(a, a[2:])] // 删除开头N个元素
}

func TestDelWithAppend(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 9, 10, 11, 12, 13}

	i := 2
	n := 3
	a = append(a[:i], a[i+1:]...) // 删除中间1个元素
	a = append(a[:i], a[i+n:]...) // 删除中间N个元素

	a = a[:i+copy(a[i:], a[i+1:])] // 删除中间1个元素
	a = a[:i+copy(a[i:], a[i+n:])] // 删除中间N个元素
}

func TestSliceMem(t *testing.T) {
	var a []*int

	// 若沈略此步，最后一个元素不会释放，直到所有引用都为0
	// 引用带来的是gc不能回收
	a[len(a)-1] = nil // GC回收最后一个元素内存
	a = a[:len(a)-1]  // 从切片删除最后一个元素
}

func TestFloatSort(t *testing.T) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

	// 强制类型转换
	var b []int = ((*[1 << 20]int)(unsafe.Pointer(&a[0])))[:len(a):cap(a)]

	// 以int方式给float64排序
	sort.Ints(b)
}

func TestFloatSortReflect(t *testing.T) {
	var a = []float64{4, 2, 5, 7, 2, 1, 88, 1}

	// 通过 reflect.SliceHeader 更新切片头部信息实现转换
	aHdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))

	// 获取底层指针
	var c []int
	cHdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))

	// c 的底层指向 a 的底层
	*cHdr = *aHdr

	// 以int方式给float64排序
	sort.Ints(c)
}
