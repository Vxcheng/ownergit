package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Add(a, b uint64) uint64 {
	return uint64(a + b)
}

func Append() {
	slic := make([]string, 0)
	child := signal()
	for _,_  = range child {

	}

	slic = append(slic, child...)
}

func signal() []string {
	return nil
}

func main() {
	Append()
	fmt.Println(Add(2, 15))
	fmt.Printf("node: %v\n", Storage)
	a := int64(10)
	fmt.Printf("fl64: %f\n", float64(a))
	b := float64(1)
	if b == 1 {
		fmt.Println("===")
	}

	c := int64(2)
	fmt.Printf("int: %d\n", int(c))
	d := []string{"aaa", "0.0"}
	dF, _ := strconv.ParseFloat(d[1], 64)
	fmt.Println(dF)
	out := Formatted("DD/_DROPPED_0002_OCR")
	fmt.Println("out: ", out)
	convertString()
}

func Formatted(key string) string {

	aa := strings.Fields(key)
	return strings.ToLower(strings.Join(aa, "_"))
}

type NodeType int64

const (
	Storage NodeType = iota + 1
	Compute
	Fusion
	Quorum
)

func convertString() {
	str := "hel"
	fmt.Println("index: ", str[2])
	for i, v := range str {
		fmt.Println("i: ", i, ", v:", v)
	}

	runs := []rune(str)
	for i, v := range runs {
		fmt.Println("i: ", i, ", s:", string(v))
	}
}
