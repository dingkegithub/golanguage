package tree

import (
	"fmt"
	"testing"
)

func compare(a, b interface{}) int {
	intA := a.(int)
	intB := b.(int)

	if intA == intB {
		return 0
	} else if intA > intB {
		return 1
	} else {
		return -1
	}
}

func printSlice(name string, s []interface{}) {
	fmt.Println()
	fmt.Println(name, "-----------------------------------------")
	for _, v := range s {
		fmt.Print(v, "\t")
	}
	fmt.Println()
}

func equalTraverse(a, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}

	for i, av := range a {
		bv := b[i]

		if av.(int) != bv.(int) {
			return false
		}
	}

	return true
}

func TestBinary(t *testing.T) {
	binary := &Binary{
		Compare: compare,
	}

	pre := []interface{}{20, 10, 8, 16, 15, 29, 24, 25, 50}
	mid := []interface{}{8, 10, 15, 16, 20, 24, 25, 29, 50}
	post := []interface{}{8, 15, 16, 10, 25, 24, 50, 29, 20}
	level := []interface{}{20, 10, 29, 8, 16, 24, 50, 15, 25}

	binary.BuildWithPreMidOrder(pre, mid)
	if binary.Empty() {
		t.Error("failed build tree")
		t.FailNow()
	}

	var preTraverse, midTraverse, postTraverse, levelTraverse []interface{}
	binary.PreTraverse(&preTraverse)
	if !equalTraverse(preTraverse, pre) {
		t.Log("pre traverse failed exception")
		printSlice("want   pre: ", pre)
		printSlice("actual pre: ", preTraverse)
	}

	binary.MidTraverse(&midTraverse)
	if !equalTraverse(midTraverse, mid) {
		t.Log("mid traverse failed exception")
		printSlice("want   pre: ", mid)
		printSlice("actual pre: ", midTraverse)
	}

	binary.PostTraverse(&postTraverse)
	if !equalTraverse(postTraverse, post) {
		t.Log("post traverse failed exception")
		printSlice("want   pre: ", post)
		printSlice("actual pre: ", postTraverse)
	}

	binary.LevelTraverse(&levelTraverse)
	if !equalTraverse(levelTraverse, level) {
		t.Log("level traverse failed exception")
		printSlice("want   pre: ", level)
		printSlice("actual pre: ", levelTraverse)
	}
}

//
// golang 中slice 作为函数参数是值传递，非引用
// 若想引用传递，必须指针
//
//func addS(s []interface{}) {
func addS(s *[]interface{}) {
	for i := 0; i < 10; i++ {
		*s = append(*s, i)
	}
}

func TestSlice(t *testing.T) {
	var s []interface{}
	addS(&s)
	printSlice("slice", s)
}
