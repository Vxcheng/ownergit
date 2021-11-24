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
	stu2_chan()
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
