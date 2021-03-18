package behavioral

import "testing"

func TestStrategy(t *testing.T) {
	t.Run("AddOperate", func(t *testing.T) {
		c := NewStrategyContext(new(AddOperate))
		got := c.Execute(2, 3)
		t.Log(got)
	})

	t.Run("MultiplyOperate", func(t *testing.T) {
		c := NewStrategyContext(new(MultiplyOperate))
		got := c.Execute(2, 3)
		t.Log(got)
	})
}
