package main

import (
	"fmt"

	. "ownergit/disgin_patterns/active"
)

func main() {
	fmt.Println("-------二十三种常用设计模式-------")
	selectPattern("Active Object")
	fmt.Println("-----------finish!!!--------")
}

func selectPattern(patternName string) {
	switch patternName {
	case "Active Object":
		Stu_ActiveObject()
	default:
		fmt.Printf("unknown patternName: %s", patternName)
	}
}
