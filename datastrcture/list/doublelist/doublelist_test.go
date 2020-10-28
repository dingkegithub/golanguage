package doublelist

import (
	"testing"

	"github.com/dingkegithub/golanguage/datastrcture/list"
)

func TestNewDoubleList(t *testing.T) {
	doubleList := NewDoubleList(list.WithIntCompare())

	want := []int{1, 4, 3, 2, 0, 5, 7}

	for _, v := range want {
		doubleList.Insert(v)
	}

	for i := 0; i < len(want); i++ {
		node := doubleList.Index(i)
		if node == nil {
			t.Log("error index ", i)
			t.FailNow()
		}

		nodeData, ok := node.(int)
		if !ok {
			t.Log("type not match: ", i)
			t.FailNow()
		}

		if nodeData != want[i] {
			t.Log("index:", i, "want:", want[i], "actual:", nodeData)
			t.FailNow()
		}
	}

	want2 := []int{1, 3, 0, 7}
	doubleList.Delete(4)
	doubleList.Delete(2)
	doubleList.Delete(5)

	if doubleList.Size() != 4 {
		t.Log("wantSize: ", 4, "actualSize: ", doubleList.Size())
		t.FailNow()
	}

	for i := 0; i < len(want2); i++ {
		node := doubleList.Index(i)
		if node == nil {
			t.Log("error index ", i)
			t.FailNow()
		}

		nodeData, ok := node.(int)
		if !ok {
			t.Log("type not match: ", i)
			t.FailNow()
		}

		if nodeData != want2[i] {
			t.Log("index:", i, "want:", want2[i], "actual:", nodeData)
			t.FailNow()
		}
	}

	doubleList.Insert(9)
	idx, err := doubleList.Query(9)
	if err != nil {
		t.Log("query error: ", err)
		t.FailNow()
	}

	if idx != 4 {
		t.Log("insert position error want 4, but", idx)
		t.FailNow()
	}
}
