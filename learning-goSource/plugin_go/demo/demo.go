package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("main.so")
	if err != nil {
		return
	}

	s, err := p.Lookup("Speek")
	if err != nil {
		return
	}

	c, err := p.Lookup("P")
	if err != nil {
		return
	}
	fmt.Printf("s: %v, c: %v\n", s, c)

}
