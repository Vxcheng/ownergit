package conversion

import (
	"testing"
)

func TestBase(t *testing.T) {
	t.Run("decimalToAny", func(t *testing.T) {
		got := decimalToAny(11, 2)
		if got != "1011" {
			t.Errorf("got is %s", got)
		}
	})

	t.Run("anyToDecimal", func(t *testing.T) {
		got := anyToDecimal("1a", 16)
		if got != 26 {
			t.Errorf("got is %d", got)
		}
	})

}
