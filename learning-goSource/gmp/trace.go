package main

import (
	"os"
	"runtime/trace"
)

func traceG() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	if err = trace.Start(f); err != nil {
		panic(err)
	}

	trace.Stop()
}
