package bridge

import "testing"

func TestBridge(t *testing.T) {

	rc := &RemoteControl{}

	rc.TvMode(&TV{})
	rc.VolUp()
	rc.VolDown()

	rc.RadioMode(&Radio{})
	rc.VolUp()
	rc.VolDown()
}
