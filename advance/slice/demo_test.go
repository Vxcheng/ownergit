package slice

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func BenchmarkDemo(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		demo1()
	}
}

func TestDemo(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		demo1()
		Comp()
	})

	t.Run("demo2", func(t *testing.T) {
		demo2()
	})
}

func FuzzXxx(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, sed int) {
		demo1()
	})
}
