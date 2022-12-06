package slice

import (
	"fmt"
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
	t.Run("", func(t *testing.T) {
		demo1()
	})
}

func ExampleDemo() {
	fmt.Println("hello")
	// Output: hello
}

func FuzzXxx(f *testing.F) {
	f.Add(1)
	f.Fuzz(func(t *testing.T, sed int) {
		demo1()
	})
}
