package haier

import "github.com/dingkegithub/golanguage/designpattern/create/absfactory"

type HaierFactory struct {
}

func (h *HaierFactory) CreateFridge() absfactory.Fridge {
	return NewHaierBridge("HaierFridge", 15)
}

func (h *HaierFactory) CreateTv() absfactory.Tv {
	return NewHaierTv(65, 45, "HaierTv")
}
