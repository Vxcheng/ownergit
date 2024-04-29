package slice

import (
	"fmt"
	"testing"
)

func TestComp(t *testing.T) {
	t.Run("", func(t *testing.T) {
		Comp()
	})
}

func Comp() {
	fmt.Println(min(1, 2, 3)) //
}
