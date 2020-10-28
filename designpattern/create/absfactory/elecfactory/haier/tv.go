package haier

import "fmt"

type HaierTv struct {
	Width     int
	Height    int
	ModelName string
}

func NewHaierTv(h, w int, name string) *HaierTv {
	return &HaierTv{
		Width:     w,
		Height:    h,
		ModelName: name,
	}
}

func (h *HaierTv) Screen(width, height int) {
	h.Width = width
	h.Height = height
}

func (h *HaierTv) TurnOn() bool {
	fmt.Println(h.ModelName, ": Welcome")
	return true
}

func (h *HaierTv) TurnOff() bool {
	fmt.Println(h.ModelName, ": Goodbye")
	return true
}

func (h *HaierTv) Model() string {
	return h.ModelName
}
