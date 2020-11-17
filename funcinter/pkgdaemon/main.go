package main

import (
	"fmt"
	"time"

	"github.com/dingkegithub/golanguage/funcinter/pkgdaemon/pkg2"
	"github.com/dingkegithub/golanguage/funcinter/pkgdaemon/pkg3"
)

var (
	varName = "main-var-main"
)

const (
	constName = "main-const-main"
)

func init() {

	fmt.Println("main init: ", varName, constName)
}

func main() {
	fmt.Println("main running")
	pkg2.PrintName()
	pkg3.PrintName()
	time.Sleep(10)
	fmt.Println("main exit")
}
