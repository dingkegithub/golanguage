package guomei

import "github.com/dingkegithub/golanguage/designpattern/create/absfactory"

type GmFactory struct {
}

func (h *GmFactory) CreateFridge() absfactory.Fridge {
	return NewGmBridge("GmFridge", 15)
}

func (h *GmFactory) CreateTv() absfactory.Tv {
	return NewGmTv(65, 45, "GmTv")
}
