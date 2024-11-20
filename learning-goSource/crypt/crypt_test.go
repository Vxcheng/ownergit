package crypt

import "testing"

func TestXxx(t *testing.T) {
	t.Run("demoAES", func(t *testing.T) {
		demoAES()
	})

	t.Run("demoRSA", func(t *testing.T) {
		demoRSA()
	})
}
