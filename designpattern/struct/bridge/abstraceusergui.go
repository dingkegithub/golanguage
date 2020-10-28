package bridge

type RemoteControl struct {
	device Device
	light  string
}

func (rc *RemoteControl) VolUp() {
	rc.device.SetVolume(rc.device.GetVolume() + 1)
}

func (rc *RemoteControl) VolDown() {
	rc.device.SetVolume(rc.device.GetVolume() - 1)
}

func (rc *RemoteControl) TvMode(device Device) {
	rc.light = "tv灯亮"
	rc.device = device
}

func (rc *RemoteControl) RadioMode(device Device) {
	rc.light = "radio 灯亮"
	rc.device = device
}
