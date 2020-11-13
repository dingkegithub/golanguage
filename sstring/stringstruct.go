package sstring

import (
	"reflect"
	"unsafe"
)

func StringSize(s string) bool {

	l1 := len(s)

	l2 := (*reflect.StringHeader)(unsafe.Pointer(&s)).Len

	return l1 == l2
}
