package guomei

import "fmt"

type GmTv struct {
	Width     int
	Height    int
	ModelName string
}

func NewGmTv(h, w int, name string) *GmTv {
	return &GmTv{
		Width:     w,
		Height:    h,
		ModelName: name,
	}
}

func (h *GmTv) Screen(w, h int) {
	h.Width = w
	h.Height = h
}

func (h *GmTv) TurnOn() bool {
	fmt.Println(h.ModelName, ": Welcome")
	return true
}

func (h *GmTv) TurnOff() bool {
	fmt.Println(h.ModelName, ": Goodbye")
	return true
}

func (h *GmTv) Model() string {
	return h.ModelName
}
