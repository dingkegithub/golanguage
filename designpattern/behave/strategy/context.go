package strategy

type StrategyContext struct {
	s Strategy
}

func (sc *StrategyContext) SetStrategy(s Strategy) {
	sc.s = s
}

func (sc *StrategyContext) ExecStrategy(a, b int) {
	sc.s.Execute(a, b)
}
