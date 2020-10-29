package algo

import "github.com/gogo/protobuf/plugin/compare"

type SortType int

const (
	SortTypeInsert SortType = 1 << iota
	SortTypeBubble
	SortTypeHeap
	SortTypeMerge
	SortTypeQuick
)

func (st SortType) String() string  {
	switch st {
	case SortTypeInsert:
		return "InsertSort"
	case SortTypeBubble:
		return "BubbleSort"
	case SortTypeHeap:
		return "HeapSort"
	case SortTypeMerge:
		return "MergeSort"
	case SortTypeQuick:
		return "QuickSort"
	default:
		return ""
}

type CompareFunc func(interface{}, interface{}) int

type Sort interface {
	Sort([]interface{}, int)
}

type InsertSort struct {
}

func (is *InsertSort) Sort(pokers []int) {

	size := len(pokers)

	for i := 1; i < size; i++ {
		newPoker := pokers[i]

		j := i - 1
		for j >= 0 {
			if pokers[j] < newPoker {
				break
			}

			pokers[j+1] = pokers[j]
			j -= 1
		}
		pokers[j+1] = newPoker
	}
}

func BubbleSort(bubbles []int) {
	size := len(bubbles)

	// 冒泡轮次
	bubbleCircle := size - 1

	for i := 0; i < bubbleCircle; i++ {

		// 冒泡区间
		bubbleRange := size - i - 1

		for j := 0; j < bubbleRange; j++ {
			// 决定下一个继续向上冒的元素
			if bubbles[j] > bubbles[j+1] {
				bubbles[j], bubbles[j+1] = bubbles[j+1], bubbles[j]
			}
		}
	}
}

func heapfy(arr []int, father int, size int) {
	// 左子节点定义 2 * i， 但是序列索引从0开始
	// 因此堆中索引转换到序列时，要 - 1
	lChild := (father << 1) - 1
	rChild := lChild

	biggerIdx := -1

	// 计算左右子节点最大的那个
	if lChild < size && rChild < size {
		if arr[lChild] >= arr[rChild] {
			biggerIdx = lChild
		} else {
			biggerIdx = rChild
		}
	} else if lChild < size && rChild >= size {
		biggerIdx = lChild
	}

	if biggerIdx != -1 && arr[biggerIdx] > arr[father-1] {
		arr[biggerIdx], arr[father-1] = arr[father-1], arr[biggerIdx]

		// biggerIdx 对应元素被调整到了下方，因此再看 biggerIdx 对应的
		// 父节点是否满足大堆，序列索引转换到堆索引，需要 +1
		heapfy(arr, biggerIdx+1, size)
	}
}

func HeapSort(arr []int) {
	size := len(arr)

	// 最后一个父节点的索引 2 * i + 1
	// 堆索引转换成序列索引要 -1
	lastParentIdx := size >> 1

	// 初始堆构建堆: 大堆
	for lastParentIdx > 0 {
		heapfy(arr, lastParentIdx, size)
		lastParentIdx -= 1
	}

	// j 控制参与堆排序的范围
	j := 0
	for j < size {
		// 堆顶和最后一个叶子节点交换
		arr[0], arr[size-j-1] = arr[size-j-1], arr[0]
		// 交换之后，对顶放到了最后一个叶子节点，也就是参与本次
		// 堆排序序列的最后一个位置，下次开始之前要排除这个节点
		// j 表示每次排除最后j个节点
		j += 1
		heapfy(arr, 1, size-j)
	}
}

func merge(arr []int, l, mid, r int) {
	// 额外内存分配
	lArr := make([]int, mid-l+1)
	rArr := make([]int, r-mid)
	copy(lArr, arr[l:mid+1])
	copy(rArr, arr[mid+1:r+1])
	i := 0
	j := 0
	k := l

	// 子序列比较合并
	for i < len(lArr) && j < len(rArr) {
		if lArr[i] < rArr[j] {
			arr[k] = lArr[i]
			i += 1
		} else {
			arr[k] = rArr[j]
			j += 1
		}
		k += 1
	}

	// 比较滞后总有一边剩下几个元素
	// 把剩余的直接放进序列
	println(i, j, k)
	for i < len(lArr) {
		arr[k] = lArr[i]
		k += 1
		i += 1
	}

	for j < len(rArr) {
		arr[k] = rArr[j]
		k += 1
		j += 1
	}
}

func split(arr []int, l, r int) {
	if l < r {
		// 计算拆的中心
		mid := l + ((r - l) >> 1)

		// 递归拆
		split(arr, l, mid)
		split(arr, mid+1, r)

		// 拆完合并
		merge(arr, l, mid, r)
	}
}

func MergeSort(arr []int) {
	split(arr, 0, len(arr)-1)
}

/*
i =  -1  l= 0  r= 4
i =  -1  j= 0  arr[j]= 5  pivot= 6
5, 8, 2, 9, 6,
i =  0  j= 1  arr[j]= 8  pivot= 6
5, 8, 2, 9, 6,
i =  0  j= 2  arr[j]= 2  pivot= 6
5, 8, 2, 9, 6,
i =  1  j= 3  arr[j]= 9  pivot= 6
5, 8, 2, 9, 6,
*/
func partition(arr []int, l, r int) int {
	pivot := arr[r]
	i := l - 1

	for j := l; j < r; j++ {
		if arr[j] < pivot {
			i += 1

			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
		Printf(arr)
	}

	if i+1 < r {
		arr[i+1], arr[r] = arr[r], arr[i+1]
	}

	return i + 1
}

func QuickSort(arr []int, l, r int) {
	if l < r {
		mid := partition(arr, l, r)
		QuickSort(arr, l, mid-1)
		QuickSort(arr, mid+1, r)
	}
}

func Printf(arr []int) {
	for _, v := range arr {
		print(v, ", ")
	}
	println()
}

