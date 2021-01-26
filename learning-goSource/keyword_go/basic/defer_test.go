package basic

import "testing"

func TestDefer(t *testing.T) {
	t.Run("", func(t *testing.T) {
		deferStu1()
		deferStu2()
		deferStu3()
	})

	t.Run("", func(t *testing.T) {
		r, err := deferStu4()
		if err != nil {
			t.Log(err)
		}
		t.Log(r)
	})

	t.Run("", func(t *testing.T) {
		v := deferInc()
		t.Log(v)
	})

	t.Run("", func(t *testing.T) {
		deferPrint()
		deferPrint1()
		deferPrint2()
	})
}
