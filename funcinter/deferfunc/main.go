package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println("error: ", i)
		}()
	}

	for i := 0; i < 3; i++ {
		defer func(v int) {
			fmt.Println("ok: ", v)
		}(i)
	}
}
