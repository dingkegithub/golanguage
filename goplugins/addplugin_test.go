package main

import (
	"fmt"
	"plugin"
	"reflect"
	"testing"
)

func TestAddPlugin(t *testing.T) {
	p, err := plugin.Open("add.so")
	if err != nil {
		t.Error("load add plugin so err:", err)
		t.FailNow()
	}

	a, err := p.Lookup("A")
	if err != nil {
		t.Error("load A err:", err)
		t.FailNow()
	}

	b, err := p.Lookup("B")
	if err != nil {
		t.Error("load B err:", err)
		t.FailNow()
	}

	add, err := p.Lookup("Add")
	if err != nil {
		t.Error("load Add err:", err)
		t.FailNow()
	}

	mm, err := p.Lookup("Mm")
	if err != nil {
		t.Error("load mm err:", err)
		//		t.FailNow()
	}

	newM, err := p.Lookup("NewMm")
	if err != nil {
		t.Error("load newM err:", err)
		t.FailNow()
	}

	fmt.Println("-------------->:", reflect.TypeOf(a))
	fmt.Println("-------------->:", reflect.TypeOf(b))
	fmt.Println("-------------->:", reflect.TypeOf(add))
	fmt.Println("-------------->:", reflect.TypeOf(mm))
	fmt.Println("-------------->:", reflect.TypeOf(newM))

	*a.(*int) = 5
	*b.(*int) = 6

	add.(func())()
}
