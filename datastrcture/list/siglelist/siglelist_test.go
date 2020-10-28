package singlelist

import "testing"

func TestNewSigleList(t *testing.T) {
	sigleList := NewSingleList(WithIntCompare())

	want := []int{1, 4, 3, 2, 0, 5, 7}

	for _, v := range want {
		sigleList.Insert(v)
	}

	for i := 0; i < len(want); i++ {
		node := sigleList.Index(i)
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
	sigleList.Delete(4)
	sigleList.Delete(2)
	sigleList.Delete(5)

	if sigleList.Size() != 4 {
		t.Log("wantSize: ", 4, "actualSize: ", sigleList.Size())
		t.FailNow()
	}

	for i := 0; i < len(want2); i++ {
		node := sigleList.Index(i)
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

	sigleList.Insert(9)
	idx, err := sigleList.Query(9)
	if err != nil {
		t.Log("query error: ", err)
		t.FailNow()
	}

	if idx != 4 {
		t.Log("insert position error want 4, but", idx)
		t.FailNow()
	}
}
