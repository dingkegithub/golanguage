package main

import (
	"plugin"
	"reflect"
	"testing"
)

func TestRpc(t *testing.T) {
	p, err := plugin.Open("rpc.so")
	if err != nil {
		t.Log("load rpc error", err)
		t.FailNow()
	}

	req, err := p.Lookup("Request")
	if err != nil {
		t.Log("loakup symbol Request", err)
		t.FailNow()
	}

	// 1. 值
	value := reflect.ValueOf(req)
	t.Log("type is: ", value)

	// 2. 值类型
	if value.Type().Kind() != reflect.Func {
		t.Log("type is: func")
		t.FailNow()
	}

	// 取出输入输出参数
	numIn := value.Type().NumIn()
	numOut := value.Type().NumOut()
	t.Log("in: ", numIn, "out: ", numOut)

	urlEncode := map[string]interface{}{
		"Uid":   12345,
		"Token": "aaaaaa",
	}
	print(len(urlEncode))

	// 2. 设置输入参数
	for i := 0; i < numIn; i++ {
		p := value.Type().In(i)

		t.Log("pin", p.Kind(), p.Elem().Kind(), p.Elem().Name())

		for j := 0; j < p.Elem().NumField(); j++ {
			structFields := p.Elem().Field(j)
			t.Log("fields: ", j, " =", structFields.Name)
		}
	}
}
