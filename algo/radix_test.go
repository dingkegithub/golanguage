package algo

import (
	"fmt"
	"testing"
)

var srcArr = []int{121, 432, 564, 23, 1, 45, 788}

func TestRadixSort(b *testing.T) {
	arr := make([]int, len(srcArr))
	copy(arr, srcArr)

	rs := &RadixSort{
		arr: arr,
	}

	rs.radixSort()

	for _, v := range rs.arr {
		fmt.Println(v, "\t")
	}
	fmt.Println()
}
