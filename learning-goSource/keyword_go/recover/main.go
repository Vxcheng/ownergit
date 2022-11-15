package main

import (
	"fmt"
	"time"
)

func main() {
	demo5()

	time.Sleep(time.Second)
}

func demo1() {
	defer recover()

	func() {
		panic(4041)
	}()
	time.Sleep(time.Second)
}

func myRecover() {
	if err := recover(); err != nil {
		fmt.Println(err, 1)
	}
}

func demo2() {
	defer func() {
		myRecover()
	}()
	func() {
		panic(4041)
	}()
	time.Sleep(time.Second)
}

func demo3() {
	defer myRecover()
	go func() {
		panic(4041)
	}()
	time.Sleep(time.Second)
}

func demo4() {
	defer myRecover()
	panic(4)
	time.Sleep(time.Second)
}

func demo5() {
	go func() {
		panic(5)
	}()
}
