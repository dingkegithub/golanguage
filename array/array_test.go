package array

import (
	"fmt"
	"testing"
)

func TestArrayParam(t *testing.T) {
	a := [...]int{1, 2, 3}

	ArrayAsParam(&a)

	fmt.Print("数组指针for-i: ")
	for i := 0; i < len(a); i++ {
		fmt.Print(a[i], "\t")
	}
	fmt.Println()
}
