package main

import (
	"fmt"
	"math"
)

func demo1() {
	m := make(map[string]int)
	b := m["b"]
	_ = b
	_, ok := m["a"]
	_ = ok
}

func demo2() {
	m := make(map[string]int)
	m["b"] = 1
	for k, v := range m {
		_, _ = k, v
	}
}

func demo3() {
	m := make(map[float64]int)
	m[1.4] = 1
	m[2.4] = 2
	m[math.NaN()] = 3
	m[math.NaN()] = 3

	for k, v := range m {
		fmt.Printf("[%v, %d] ", k, v)
	}

	fmt.Printf("\nk: %v, v: %d\n", math.NaN(), m[math.NaN()])
	fmt.Printf("k: %v, v: %d\n", 2.400000000001, m[2.400000000001])
	fmt.Printf("k: %v, v: %d\n", 2.4000000000000000000000001, m[2.4000000000000000000000001])

	fmt.Println(math.NaN() == math.NaN())
}
