package encryption

import "testing"

func TestBit(t *testing.T) {
	t.Run("", func(t *testing.T) {
		arr := []int{6, 1, 1, 3, 3, 4, 5, 5, 4}
		printOneOddTimesNum(arr)
	})

	t.Run("", func(t *testing.T) {
		arr := []int{6, 1, 3, 3, 4, 5, 5, 4}
		printOddTimesNum(arr)
	})
}
