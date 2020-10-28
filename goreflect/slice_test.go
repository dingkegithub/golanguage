package main

import (
	"log"
	"reflect"
	"testing"
)

type Child struct {
	Name  string
	Grade int
	Nice  bool
}

type Adult struct {
	Name       string
	Occupation string
	Nice       bool
}

// search a slice of structs for Name field that is "Hank" and set its Nice
// field to true.
func nice(i interface{}) {
	// retrieve the underlying value of i.  we know that i is an
	// interface.
	v := reflect.ValueOf(i)

	// we're only interested in slices to let's check what kind of value v is. if
	// it isn't a slice, return immediately.
	if v.Kind() != reflect.Slice {
		return
	}

	// v is a slice.  now let's ensure that it is a slice of structs.  if not,
	// return immediately.
	if e := v.Type().Elem(); e.Kind() != reflect.Struct {
		return
	}

	// determine if our struct has a Name field of type string and a Nice field
	// of type bool
	st := v.Type().Elem()

	if nameField, found := st.FieldByName("Name"); found == false || nameField.Type.Kind() != reflect.String {
		return
	}

	if niceField, found := st.FieldByName("Nice"); found == false || niceField.Type.Kind() != reflect.Bool {
		return
	}

	// Set any Nice fields to true where the Name is "Hank"
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		name := e.FieldByName("Name")
		nice := e.FieldByName("Nice")

		if name.String() == "Hank" {
			nice.SetBool(true)
		}
	}
}

func TestSlice(t *testing.T) {
	children := []Child{
		{Name: "Sue", Grade: 1, Nice: true},
		{Name: "Ava", Grade: 3, Nice: true},
		{Name: "Hank", Grade: 6, Nice: false},
		{Name: "Nancy", Grade: 5, Nice: true},
	}

	adults := []Adult{
		{Name: "Bob", Occupation: "Carpenter", Nice: true},
		{Name: "Steve", Occupation: "Clerk", Nice: true},
		{Name: "Nikki", Occupation: "Rad Tech", Nice: false},
		{Name: "Hank", Occupation: "Go Programmer", Nice: false},
	}

	log.Printf("adults before nice: %v", adults)
	nice(adults)
	log.Printf("adults after nice: %v", adults)

	log.Printf("children before nice: %v", children)
	nice(children)
	log.Printf("children after nice: %v", children)
}
