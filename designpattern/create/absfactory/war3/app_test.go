package war3

import "testing"

func TestNewWar3(t *testing.T) {

	war3 := NewWar3(ORC)
	war3.Start()

	war3 = NewWar3(HUM)
	war3.Start()
}
