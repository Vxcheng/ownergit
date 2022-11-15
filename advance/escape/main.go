package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

func demo1() {
	x := foo()
	fmt.Println(*x)
}

// go build -gcflags '-m -l' main.go
func main() {
	// demo1()
	demo2()
}

type S struct{}

func demo2() {
	var x S
	_ = identity(x)
}

func identity(x S) S {
	return x
}
