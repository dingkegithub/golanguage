package pkg1

import "fmt"

const (
	constName = "const-hello"
)

var (
	varName = "var-hello"
)

func init() {
	fmt.Println("pkg1 init: ", constName, varName)
}

func PrintName() {
	fmt.Println("pkg1 PrintName: ", constName, varName)
}
