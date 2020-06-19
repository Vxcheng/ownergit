package parallel

import (
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
