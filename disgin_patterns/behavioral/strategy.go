package behavioral

type strategy interface {
	doOperate(num1, num2 int) (final int)
}

type StrategyContext struct {
	strategy
}

func NewStrategyContext(strategy strategy) *StrategyContext {
	return &StrategyContext{
		strategy,
	}
}

func (s *StrategyContext) Execute(num1, num2 int) (final int) {
	return s.strategy.doOperate(num1, num2)
}

type AddOperate struct{}

func (o *AddOperate) doOperate(num1, num2 int) (final int) {
	final = num1 + num2
	return
}

type MultiplyOperate struct{}

func (o *MultiplyOperate) doOperate(num1, num2 int) (final int) {
	return num1 * num2
}
