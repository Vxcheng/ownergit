package creation

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("chicken burger wrapper, coke bottle", func(t *testing.T) {
		m := new(MealBuilder).PrePepsiDrink()
		fmt.Printf("pepsi cost: %.2f\n", m.GetCosts())
		m.ShowItems()
	})
}
