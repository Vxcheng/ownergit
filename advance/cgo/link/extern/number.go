package main

import "C"

//$ go build -buildmode=c-archive -o number.a
//$ go build -buildmode=c-shared -o number.so

func main() {}

//export number_add_mod1
func number_add_mod1(a, b, mod C.int) C.int {
	return (a + b) % mod
}
