package parallel

import (
	"fmt"
	"log"
	"sync"
	"time"
)

/*
chan分类

	无缓冲chan
	有缓冲chan
	单向chan
	关闭和nil
*/
func Stu_chan() {
	log.Println("学习chan")
	// stu1_chan()
	chan_stu4()
	time.Sleep(time.Second)
	log.Println("finished")

}

func stu2_chan() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover: %v", r)
		}
	}()
	{
		c := make(chan int)
		go func() {
			c <- 1
		}()
		log.Printf("c1: %v", <-c)
	}

	{
		c := make(chan int, 1)
		go func() {
			c <- 1
		}()
		log.Printf("c2: %v", <-c)
	}

	{
		c := make(chan int)
		go func() {
			log.Printf("c3: %v", <-c)
		}()
		c <- 1
	}

	{
		c := make(chan int, 1)
		go func() {
			log.Printf("c4: %v", <-c)
		}()
		c <- 1
	}
	{
		c := make(chan int, 1)
		c <- 1
		log.Printf("c5: %v", <-c)
	}

}

func stu1_chan() {
	ch1 := make(chan int)
	writeToChan(ch1)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go readFromChan(ch1, &wg)
	wg.Wait()
}

func writeToChan(ch chan int) {
	for i := 0; i < 4; i++ {
		go func(a int) {
			ch <- a
		}(i)
	}
}

func readFromChan(ch chan int, wg *sync.WaitGroup) {
	// log.Println("msg: ", <-ch)
	defer wg.Done()

	for {
		select {
		case msg := <-ch:
			log.Println("msg: ", msg)
		case <-time.After(time.Second * 2):
			return
		}
	}
}

func chan_stu3() {
	aC, bC := make(chan int), make(chan interface{})
	a, b := 1, "hi"
	go func() {
		for {
			aC <- a
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for {
			bC <- b
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-aC:
			fmt.Printf("%d\n", v)
		case v := <-bC:
			fmt.Printf("%v\n", v)
		case <-time.After(time.Second):
			fmt.Println("out")
		}
	}
}

func chan_stu4() {
	dataC := make(chan int, 5)
	go func() {
		for {
			v, ok := <-dataC
			if !ok {
				log.Println("closed, exiting")
				return
			}

			log.Printf("recieve: %d\n", v)
		}
	}()

	for i := 0; i < 20; i++ {
		dataC <- i
		log.Printf("send: %d\n", i)

	}

	close(dataC)
	log.Println("close dataC")
	time.Sleep(time.Second * 2)
}

// 如何退出子协程
func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true
}

func timeout(f func(chan bool)) error {
	done := make(chan bool) // bad, 协程泄漏，优化为	done := make(chan bool, 1)
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true:
	default:
		return
	}
}
