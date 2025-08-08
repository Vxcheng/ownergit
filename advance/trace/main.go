package main

import (
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("trace.out")

	trace.Start(f)
	defer trace.Stop()

	for i := 0; i < 10000; i++ {
		_ = make([]byte, 1<<20)
	}
}

