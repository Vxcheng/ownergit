package structural

import "testing"

func TestBridge(t *testing.T) {
	t.Run("circle red wooden", func(t *testing.T) {
		c := NewCircleDrawing(NewRedColor(), NewWooden())
		got := c.DoDraw("xiaoming")
		t.Log(got)
	})
}
