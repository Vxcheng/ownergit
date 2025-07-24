package main

import "fmt"

func main() {
	r := increase()
	// fmt.Println("r: ", increase()())
	fmt.Println("r: ", r())
	fmt.Println("r: ", r())

	fmt.Println("r: ", increase()())
	fmt.Println("r: ", increase()())
}

func increase() func() int {
	n := 0
	return func() int {
		n++
		return n
	}
}
