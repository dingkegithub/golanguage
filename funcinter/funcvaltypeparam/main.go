package main

import "fmt"

func Update(arr []int) {
	if len(arr) <= 0 {
		return
	}
	arr[0] = 9999
}

func Add(arr []int) {
	items := []int{100, 101, 102}
	arr = append(arr, items...)
}

func PrintArr(name string, a []int) {
	fmt.Print(name, ": ")
	for _, v := range a {
		fmt.Print(v, '\t')
	}
	fmt.Println()
}

func main() {
	// len, cap  is 3
	arr := []int{1, 2, 3}

	Update(arr)
	PrintArr("update: ", arr)

	fmt.Println()
	Add(arr)
	PrintArr("add: ", arr)
}
