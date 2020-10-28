package absfactory

type Fridge interface {
	TempUpper() int

	TempDower() int

	Model() string
}
