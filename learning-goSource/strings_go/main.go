package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type User struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("测试strings")
	// testChan()
	// Split()
}

func ParseFloat() {
	slot := 10
	fmt.Printf("slot: %s\n", strconv.Itoa(slot))
	str := "0"
	strF64, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatalf("err: %s", err.Error())
	}
	log.Printf("%2.1f\n", strF64)
}

func Fields() {
	str1, str2 := "ONLINE on rac048", "ONLINE"
	fmt.Println("ret1:", strings.Fields(str1)[0])
	fmt.Println("ret2:", strings.Fields(str2)[0])
}

func Split() {
	str, delimet := "aab", "c"
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
