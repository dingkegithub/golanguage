package leetcode

// 给定两个大小为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的中位数。
// 进阶：你能设计一个时间复杂度为 O(log (m+n)) 的算法解决此问题吗
//
//
// 总结：看到 log，很明显，我们只有用到二分的方法才能达到

// 解法1：
// 使用归并的方式，合并两个有序数组，得到一个大的有序数组。
// 大的有序数组的中间位置的元素，即为中位数。
//
// 时间 O(n)
// 空间 O(m + n)
func MidNumLeve1(nums1, nums2 []int) float64 {
	n1Size := len(nums1)
	n2Size := len(nums2)

	size := n1Size + n2Size

	if size == 1 {
		if n1Size == 1 {
			return float64(nums1[0])
		} else {
			return float64(nums2[0])
		}
	}

	i, j, k := 0, 0, 0
	out := make([]int, size)

	// 归并排序放到一个数组
	for k < size {
		if i < n1Size && (j >= n2Size || nums1[i] <= nums2[j]) {
			out[k] = nums1[i]
			i += 1
		} else {
			out[k] = nums2[j]
			j += 1
		}
		k += 1
	}

	if k&1 == 0 {
		e := k / 2
		return float64(out[e]+out[e-1]) / 2.
	} else {
		e := k / 2
		return float64(out[e])
	}
}

// 不需要合并两个有序数组，只要找到中位数的位置即可。由于两个数组的长度已知，
// 因此中位数对应的两个数组的下标之和也是已知的。维护两个指针，初始时分别指
// 向两个数组的下标 00 的位置，每次将指向较小值的指针后移一位（如果一个指针
// 已经到达数组末尾，则只需要移动另一个数组的指针），直到到达中位数的位置
// 奇数: 1 2 3 4 5  -》中位数: 5 / 2 = 2 ===> 中位数下标：len / 2
// 偶数: 1 2 3 4    -》中位数: 4 / 2 = 2 ===> 中位数下标：(len/2) -1、len / 2
// 得出最多遍历 len  / 2 次
//
//
// 时间 O(n)
// 空间 O(1)
func MidNumLeve2(num1, nums2 []int) float64 {
	n1Len, n2Len := len(nums1), len(nums2)
	size := n1Len + n2Len

	// 遍历总次数
	forNums := size / 2

	// 暂存遍历的结果
	left, right := 0, 0

	// k 指针标记最大遍历次数
	// xStart 指针移动，每次较小的那个移动
	k, n1Start, n2Start := 0, 0, 0

	for k <= forNums {

		// 暂存左边的数，偶数时左边的数字是第一个中位数
		left = right

		// 移动较小的那个指针
		if n1Start < n1Len && ((n2Start >= n2Len) || (nums1[n1Start] < nums2[n2Start])) {
			right = nums1[n1Start]
			n1Start += 1
		} else {
			right = nums2[n2Start]
			n2Start += 1
		}
		k += 1
	}

	// 奇数判断
	if size&1 == 1 {
		// 奇数直接去右边的数
		return float64(right)
	} else {
		// 偶数时两个中位数，求平均值
		return float64(left+right) / 2.
	}
}

//
// 1  2  3  4  5  6  -> 6 + 1/ 2 = 3,

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)

	// 奇数还是偶数情况
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0

	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1

		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}

	return 0
}
