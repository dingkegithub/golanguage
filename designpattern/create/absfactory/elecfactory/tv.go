package absfactory

type Tv interface {
	Screen(w, h int)

	TurnOn() bool

	TurnOff() bool

	Model() string
}
