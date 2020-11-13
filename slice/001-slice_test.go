package slice

import (
	"bytes"
	"testing"
)

// 总结
//
// 切片是对底层数组的引用，切片在不越界底层数组cap的操作下，
// 所有切片的操作都将影响底层数据，一旦 append 等操作越界
// cap 都将导致slice被重新分配底层数组地址，此时slice的操作
// 都将原来的数组无关，此时slice是完全新的一个数组
//
// append: a = append(a, ...) 之所以这样定义是因为一旦越过原始cap，将引用新的地址数组, 否则在原数组上操作
//

/*
      path     size=14, cap = 15
   --------------------------------------------
   0  1  2  3  4  5  6  7  8  9  10  11  12  13
   --------------------------------------------
   A  A  A  A  /  B  B  B  B  B  B   B   B   B
   +---------     ----------------------------
       dir1                  dir2

假定数组 path 的 size 14, cap 15
dir1 和 dir 引用了原始数组

   1. dir1 的 size 和 cap 分别是多少?

       dir1: AAAA
       size: 4
	   cap:  15

   2. dir2 的 size 和 cap 分别是多少?

       dir2: BBBBBBBBB
       size: 9
	   cap:  15 - 5 = 10

    3. dir1 中添加元素，假设添加元素不超过 path 的cap 15

       dir1 = append(dir1, "suffix"...), 请问此时 path 是
	   设么？dir2 又是什么？

	   dir1: AAAAsuffix
	   dir2: uffixBBBB
	   path: AAAAsuffix/uffixBBBB

    4. dir1 中添加元素，假设添加元素超过 path 的cap 15

       dir1 = append(dir1, "suffixsuffixsuffix"...), 请问此时 path 是
	   设么？dir2 又是什么？

	   dir1: AAAAsuffixsuffixsuffix
	   dir2: BBBBBBBBB
	   path: AAAA/BBBBBBBBB
*/
func TestByteSlice(t *testing.T) {
	// size: 14
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	// want 4
	if sepIndex != 4 {
		t.Log("want: 4, but: ", sepIndex)
		t.FailNow()
	}

	// dir1: AAAA
	// size: 4
	// cap:  15
	dir1 := path[:sepIndex]
	if cap(dir1) != cap(path) || string(dir1) != "AAAA" || len(dir1) != 4 {
		t.Error("cap want: AAAA, 4, 15;  but: ", string(dir1), len(dir1), cap(dir1))
		t.FailNow()
	}

	//  dir2: BBBBBBBBB
	//  size: 9
	//  cap:  15 - 5 = 10
	dir2 := path[sepIndex+1:]
	t.Log("dir2: ", string(dir2))
	want := cap(path) - sepIndex + 1
	if cap(dir2) != want || string(dir2) != "BBBBBBBBB" || len(dir2) != 9 {
		t.Error("cap want: BBBBBBBBB, 9, ", want, " but: ", string(dir1), len(dir1), cap(dir1))
		t.FailNow()
	}

	//   dir1 = append(dir1, "suffix"...), 请问此时 path 是
	//   设么？dir2 又是什么？
	//   dir1: AAAAsuffix
	//   dir2: uffixBBBB
	//   path: AAAAsuffix/uffixBBBB
	dir1 = append(dir1, "suffix"...)
	if string(dir1) != "AAAAsuffix" || string(dir2) != "uffixBBBB" || string(path) != "AAAAsuffix/uffixBBBB" {
		t.Error("Want: AAAAsuffix, uffixBBBB, AAAAsuffix/uffixBBBB; but: ", string(dir1), string(dir2), string(path))
		t.FailNow()
	}
}

func TestByteSliceAppend(t *testing.T) {
	// size: 14
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	// want 4
	if sepIndex != 4 {
		t.Log("want: 4, but: ", sepIndex)
		t.FailNow()
	}

	// dir1: AAAA
	// size: 4
	// cap:  15
	dir1 := path[:sepIndex]
	if cap(dir1) != cap(path) || string(dir1) != "AAAA" || len(dir1) != 4 {
		t.Error("cap want: AAAA, 4, 15;  but: ", string(dir1), len(dir1), cap(dir1))
		t.FailNow()
	}

	//  dir2: BBBBBBBBB
	//  size: 9
	//  cap:  15 - 5 = 10
	dir2 := path[sepIndex+1:]
	t.Log("dir2: ", string(dir2))
	want := cap(path) - sepIndex + 1
	if cap(dir2) != want || string(dir2) != "BBBBBBBBB" || len(dir2) != 9 {
		t.Error("cap want: BBBBBBBBB, 9, ", want, " but: ", string(dir1), len(dir1), cap(dir1))
		t.FailNow()
	}

	// 4. dir1 中添加元素，假设添加元素超过 path 的cap 15
	// dir1 = append(dir1, "suffixsuffixsuffix"...), 请问此时 path 是
	// 设么？dir2 又是什么？
	//   dir1: AAAAsuffixsuffixsuffix
	//   dir2: BBBBBBBBB
	//   path: AAAA/BBBBBBBBB
	dir1 = append(dir1, "suffixsuffixsuffix"...)
	if string(dir1) != "AAAAsuffixsuffixsuffix" || string(dir2) != "BBBBBBBBB" || string(path) != "AAAA/BBBBBBBBB" {
		t.Error("Want: AAAAsuffixsuffixsuffix, BBBBBBBBB, AAAA/BBBBBBBBB; but: ", string(dir1), string(dir2), string(path))
		t.FailNow()
	}
}
