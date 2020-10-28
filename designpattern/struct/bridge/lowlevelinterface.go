package bridge

// 实现部分接口声明上层抽象需要使用的
// 通用方法，这部分操作可以是原语操作
type Device interface {
	GetVolume() int
	SetVolume(int)
}

// 实际底层系统具体实现
type TV struct {
	vol int
}

func (tv TV) GetVolume() int {
	return tv.vol
}

func (tv *TV) SetVolume(vol int) {
	tv.vol = vol
}

// 实际底层系统具体实现
type Radio struct {
	vol int
}

func (radio Radio) GetVolume() int {
	return radio.vol
}

func (radio *Radio) SetVolume(vol int) {
	radio.vol = vol
}
