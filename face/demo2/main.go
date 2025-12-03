package main

import (
	"fmt"
	"sync"
)

// 三个输出函数，dog、fish、cat，每个函数调用一个go协程，输出10个，要求输出顺序是：dog、fish、cat
const N = 10

func main() {
	// demo1()
	demo2()
}

// 编排顺序
func demo1() {
	fmt.Println("demo1")
	dogC, fishC, catC := make(chan string, 1), make(chan string, 1), make(chan string, 1)

	// 产生数据
	go dog(dogC)
	go fish(fishC)
	go cat(catC)

	// 编排
	for i := 0; i < N; i++ {
		fmt.Println(i)
		fmt.Println(<-dogC)
		fmt.Println(<-fishC)
		fmt.Println(<-catC)
	}
}

func dog(dogC chan string) {
	for i := 0; i < N; i++ {
		dogC <- "dog"
	}
}

func fish(fishC chan string) {
	for i := 0; i < N; i++ {
		fishC <- "fish"
	}
}

func cat(catC chan string) {
	for i := 0; i < N; i++ {
		catC <- "cat"
	}
}

// 通讯传递控制顺序
func demo2() {
	dogC, fishC, catC := make(chan struct{}), make(chan struct{}), make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < N; i++ {
			<-dogC
			fmt.Println(i)
			fmt.Println("dog")
			fishC <- struct{}{}
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 0; i < N; i++ {
			<-fishC
			fmt.Println("fish")
			catC <- struct{}{}
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for i := 0; i < N; i++ {
			<-catC
			fmt.Println("cat")
			if i < N-1 {
				dogC <- struct{}{}
			}

		}
	}(&wg)

	dogC <- struct{}{}
	wg.Wait()

	fmt.Println("over")
}
