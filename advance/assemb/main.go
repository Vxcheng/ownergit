// main.go
package main

import "fmt"

func main() {
	fmt.Println("Calling assembly function...")
	callAssemblyFunction()
}

//go:linkname callAssemblyFunction internal/assemblyFunction
func callAssemblyFunction()
