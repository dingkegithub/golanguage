package leetcode

import "testing"

// 给定一个矩阵 A， 返回 A 的转置矩阵。
// 矩阵的转置是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引。
//
// 案例1：
// 输入：[[1,2,3],[4,5,6],[7,8,9]]
// 输出：[[1,4,7],[2,5,8],[3,6,9]]
//
// 案例2：
// 输入：[[1,2,3],[4,5,6]]
// 输出：[[1,4],[2,5],[3,6]]
func TestMartrixTransposeForLoop(t *testing.T) {

	in := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	want := [][]int{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	out := MartrixTransposeForLoop(in)
	for i := range out {
		for j := range out[i] {
			if want[i][j] != out[i][j] {
				t.Log("want != in")
				t.FailNow()
			}
		}
	}
}
