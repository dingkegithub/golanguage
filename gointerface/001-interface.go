package gointerface

import "testing"

func TestInterface(t testing.T) {
	var data *byte

	var target interface{}

	if data != nil {
		t.Error("pointer init value want nil; but: ", data)
		t.FailNow()
	}

	if target != nil {
		t.Error("interface init want nil; but: ", target)
		t.FailNow()
	}

	if target != data {
		t.Error("interface nil != pointer nil")
		t.FailNow()
	}

	target = data
	if target == data {
		t.Error("inited interface == pointer nil")
		t.FailNow()
	}
}
