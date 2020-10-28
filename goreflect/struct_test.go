package main

import (
	"reflect"
	"testing"
)

type G struct {
	M int
	L string
}

type Ab struct {
	A int
	B string
	C []int
	G *G
}

// service=User,cmd=Login,func(cmd MM, a LL)

func TestStruct(t *testing.T) {
	ab := &Ab{}

	t.Log("ab type: ", reflect.TypeOf(ab))
	abType := reflect.TypeOf(ab)

	t.Log("-------------->:", abType.Elem().Kind())

	t.Log(abType.Elem().NumField())
	for i := 0; i < abType.Elem().NumField(); i++ {
		structField := abType.Elem().Field(i)
		t.Log("struct[", i, "]=", structField.Name, structField.Type)
	}
}
