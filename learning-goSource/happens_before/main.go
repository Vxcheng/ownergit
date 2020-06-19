// main
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	f4()
}

func f1() {
	var wg sync.WaitGroup
	var count int
	var ch = make(chan bool, 1)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			ch <- true
			count++
			time.Sleep(time.Millisecond)
			count--
			<-ch
			wg.Done()
		}()
	}
	wg.Wait()
}

func f2() {
	var c = make(chan int, 10)
	var a string
	f := func() {
		a = "hello, world" // (1)
		c <- 0             // (2)
	}
	go f()
	<-c      // (3)
	print(a) // (4)
}

func f3() {
	var c = make(chan int)
	var a string
	f := func() {
		a = "hello, world" // (1)
		<-c                // (2)
	}
	go f()
	c <- 0   // (3)
	print(a) // (4)
}

func f4() {
	var c = make(chan int, 1)
	var a string
	f := func() {
		a = "hello, world" // (1)
		<-c                // (2)
	}
	go f()
	c <- 0   // (3)
	print(a) // (4)
}
