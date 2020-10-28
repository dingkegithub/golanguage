package absfactory

type ElecFactory interface {
	CreateFridge() Fridge
	CreateTv() Tv
}
