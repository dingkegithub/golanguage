package algo

import "testing"

func TestLinearSearchSearch(t *testing.T) {
	arr := []int{1, 4, 9, 20, 32, 232, 800, 801, 990, 23232}

	i := LinearSearch(arr, 32)
	if i != 4 {
		t.FailNow()
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 4, 9, 20, 32, 232, 800, 801, 990, 23232}

	i := BinarySearch(arr, 32)
	if i != 4 {
		t.FailNow()
	}
}

func TestInterpolationSearch(t *testing.T) {
	arr := []int{1, 4, 9, 20, 32, 232, 800, 801, 990, 23232}

	i := InterpolationSearch(arr, 32)
	if i != 4 {
		t.FailNow()
	}
}

func TestExponentailSearch(t *testing.T) {
	arr := []int{1, 4, 9, 20, 32, 232, 800, 801, 990, 23232}

	i := ExponentailSearch(arr, 32)
	if i != 4 {
		t.FailNow()
	}
}
