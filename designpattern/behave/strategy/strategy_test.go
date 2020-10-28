package strategy

import "testing"

func TestStrategy(t *testing.T) {
	action := "add"

	ctxStrategy := &StrategyContext{}

	if action == "add" {
		ctxStrategy.SetStrategy(&AddStrategy{})
	} else if action == "sub" {
		ctxStrategy.SetStrategy(&SubStrategy{})
	} else {
		ctxStrategy.SetStrategy(&MultiplyStrategy{})
	}

	res := ctxStrategy.s.Execute(3, 5)

	if res != 8 {
		t.Error("action: add want:8, but get:", res)
		t.FailNow()
	}
}
