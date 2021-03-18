package behavioral

import "testing"

func TestTemplate(t *testing.T) {
	t.Run("play basketball", func(t *testing.T) {
		NewBasketballGame("aifusheng").Play()
	})

	t.Run("play football", func(t *testing.T) {
		NewFootballGame("c luo").Play()
	})
}
