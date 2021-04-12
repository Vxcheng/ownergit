package main

import "fmt"

func main() {
	debugG()
}

func debugG() {
	for i := 0; i < 4; i++ {
		fmt.Println("hello: ", i)
	}
}
