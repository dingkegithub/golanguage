package strategy

type Strategy interface {
	Execute(int, int) int
}
