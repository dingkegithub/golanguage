package leetcode

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

//
// 时间复杂度：O(cols * rows)
// 空间复杂度: O(cols * rows)
//
// OUT[R][C] = IN[C][R]
func MartrixTransposeForLoop(mt [][]int) [][]int {

	out := make([][]int, len(mt[0]))

	for cols := 0; cols < len(mt[0]); cols++ {

		out[cols] = make([]int, len(mt))
		for rows := 0; rows < len(mt); rows++ {
			out[cols][rows] = mt[rows][cols]
		}
	}

	return out
}

//
// 方阵直接内部对调
// OUT[R][C] = IN[C][R]
// OUT[C][R] = IN[R][C]
//
func OptimiseSquareMatrix(mt [][]int) [][]int {
	rows := len(mt)
	cols := len(mt[0])

	out := make([][]int, cols)
	for i := 0; i < cols; i++ {
		out[i] = make([]int, rows)
	}

	if rows == cols {
		for c := 0; c < cols; c++ {
			for r := c; r < rows; r++ {
				out[c][r] = mt[r][c]
				out[r][c] = mt[c][r]
			}
		}
	} else {
		for c := 0; c < cols; c++ {
			for r := 0; r < rows; r++ {
				out[c][r] = mt[r][c]
			}
		}
	}

	return out
}
