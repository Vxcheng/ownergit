package main

import (
	"fmt"
	"sync"
)

// Merge different channels in one channel
func Merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup

	out := make(chan int)

	// Start an send goroutine for each input channel in cs. send
	// copies values from c to out until c is closed, then calls wg.Done.
	send := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go send(c)
	}

	// Start a goroutine to close out once all the send goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func fanInMain() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		close(c1)
	}()

	go func() {
		for i := 10; i < 20; i++ {
			c2 <- i
		}
		close(c2)
	}()

	out := Merge(c1, c2)
	for n := range out {
		println(n)
	}
}

// Split a channel into n channels that receive messages in a round-robin fashion.
func Split(ch <-chan int, n int) []chan int {
	cs := make([]chan int, n)
	for i := 0; i < n; i++ {
		cs[i] = make(chan int)
	}

	// 单个 goroutine 顺序读取输入并轮询发送到各输出通道
	go func() {
		defer func() {
			for _, c := range cs {
				close(c)
			}
		}()

		// i := 0
		// for val := range ch {
		// 	cs[i%n] <- val
		// 	i++
		// }

		for {
			for _, c := range cs {
				val, ok := <-ch
				if !ok {
					return
				}

				c <- val

			}
		}
	}()

	return cs
}

func fanOutMain() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 20; i++ {
			ch <- i
		}
		close(ch)
	}()
	worker := 3
	cs := Split(ch, worker)
	// cs 中并发读
	var wg sync.WaitGroup
	wg.Add(worker)
	for i := 0; i < worker; i++ {
		go func(i int) {
			for val := range cs[i] {
				fmt.Println(i, val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

}
