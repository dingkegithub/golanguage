package pkg3

import "fmt"

func init() {
	fmt.Println("pkg3's first init")
}

func printStatus() {
	fmt.Println("pkg3 goruntine run")
}

func init() {
	fmt.Println("pkg3's second init")
	go printStatus()
	fmt.Println("pkg3's second init invoked goruntine")
}

func PrintName() {
	fmt.Println("pkg3 PrintName: ok")
}
