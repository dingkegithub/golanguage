package strategy

type AddStrategy struct {
}

func (as *AddStrategy) Execute(a, b int) int {
	return a + b
}

type SubStrategy struct {
}

func (as *SubStrategy) Execute(a, b int) int {
	return a - b
}

type MultiplyStrategy struct {
}

func (as *MultiplyStrategy) Execute(a, b int) int {
	return a * b
}
