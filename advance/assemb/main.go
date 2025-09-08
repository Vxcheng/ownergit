// main.go
package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	println("Calling assembly function...")
	if 1 > 2 {
		println("no")
	} else {
		println("yes")
	}

	for i := 0; i < 2; i++ {
		println(i)
	}

	println(GetGoid())
}

func GetGoid() int64 {
	var (
		buf [64]byte
		n   = runtime.Stack(buf[:], false)
		stk = strings.TrimPrefix(string(buf[:n]), "goroutine")
	)

	idField := strings.Fields(stk)[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Errorf("can not get goroutine id: %v", err))
	}

	return int64(id)
}
