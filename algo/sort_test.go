package algo

import "testing"

func TestInsertSort(t *testing.T) {
	src := []int{5, 8, 2, 9, 6}
	dst := []int{2, 5, 6, 8, 9}

	InsertSort(src)

	Printf(src)

	for i, _ := range src {
		if src[i] != dst[i] {
			t.FailNow()
		}
	}
}

func TestBubbleSort(t *testing.T) {
	src := []int{5, 8, 2, 9, 6}
	dst := []int{2, 5, 6, 8, 9}

	BubbleSort(src)

	Printf(src)

	for i, _ := range src {
		if src[i] != dst[i] {
			t.FailNow()
		}
	}
}

func TestHeapSort(t *testing.T) {
	src := []int{5, 8, 2, 9, 6}
	dst := []int{2, 5, 6, 8, 9}

	HeapSort(src)

	Printf(src)

	for i, _ := range src {
		if src[i] != dst[i] {
			t.FailNow()
		}
	}
}

func TestMergeSort(t *testing.T) {
	src := []int{5, 8, 2, 9, 6}
	dst := []int{2, 5, 6, 8, 9}

	MergeSort(src)

	Printf(src)

	for i, _ := range src {
		if src[i] != dst[i] {
			t.FailNow()
		}
	}
}
func TestSlice(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	subArr := arr[1:]

	subArr[2] = 9
	Printf(arr)
	Printf(subArr)
}

func TestQuickSort(t *testing.T) {
	src := []int{5, 8, 2, 9, 6}
	dst := []int{2, 5, 6, 8, 9}

	QuickSort(src, 0, 4)

	Printf(src)

	for i, _ := range src {
		if src[i] != dst[i] {
			t.FailNow()
		}
	}
}
