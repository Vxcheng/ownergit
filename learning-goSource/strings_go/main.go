package main

import (
	"fmt"
	"regexp"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("测试strings")
	testChan()
}
func Fields() {
	str1, str2 := "ONLINE on rac048", "ONLINE"
	fmt.Println("ret1:", strings.Fields(str1)[0])
	fmt.Println("ret2:", strings.Fields(str2)[0])
}

func Split() {
	str, delimet := "a, b", ","
	fmt.Println("Split:", strings.Split(str, delimet))
}

func removeSpacesStr() {
	str := " ONLINE on rac048 "

	reg := regexp.MustCompile("\\s+")
	fmt.Println("removeSpacesStr_", reg.ReplaceAllString(str, ""))
}

func appedAttibute() (string, error) {
	return "hello", nil
}

func testChan() {
	ch := make(chan int)
	for i := 0; i < 4; i++ {
		go func(a int) {
			ch <- a
		}(i)
	}
	for elem := range ch {
		fmt.Println("elem: ", elem)
	}
	fmt.Println("over")
}
