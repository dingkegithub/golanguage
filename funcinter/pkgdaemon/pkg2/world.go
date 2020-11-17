package pkg2

import (
	"fmt"

	"github.com/dingkegithub/golanguage/funcinter/pkgdaemon/pkg1"
)

const (
	constName = "pkg2-const-world"
)

var (
	varName = "pkg2-var-world"
)

func init() {
	fmt.Println("pkg2 init: ", constName, varName)
}

func PrintName() {
	pkg1.PrintName()
	fmt.Println("pkg2 PrintName: ", constName, varName)
}
