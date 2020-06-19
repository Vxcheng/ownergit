package main

import (
	"fmt"
	"strconv"
)

const (
	Nums = 5
)

var glabelMap map[int]int

func main() {
	fmt.Println("golang基础类型、引用类型")
	printInfo("string")
}

type Style struct {
	style string
}

func (s *Style) stu_string() {
	fmt.Printf("学习基本类型-%s\n", s.style)
	si, _ := strconv.Atoi(s.style)
	fmt.Printf("si: %v \n", si)
	for key, val := range "abcde" {
		fmt.Printf("%v, %v\t", key, val)
	}
	return
}

func printInfo(name string) {
	s := &Style{style: name}
	switch s.style {
	case "string":
		s.stu_string()
	case "rune":
	}
}

func stu_map(_ map[int]int) {
	glabelMap = make(map[int]int)
	for i := 0; i < Nums; i++ {
		glabelMap[i] = i
	}
	fmt.Printf("result map: %v\n", glabelMap)
	return
}

func stu_list(result []int) {
	for i := 0; i < Nums; i++ {
		result[i] = i
	}
	fmt.Printf("result slice: %v\n", result)
}
