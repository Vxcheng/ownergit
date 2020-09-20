package parallel

import (
	"fmt"
	"log"
	"time"
)

func Stu_select() {
	log.Println("学习select关键字")
	stu5_select()
}

// multiple channel
func stu5_select() {
	aCh, bCh := make(chan int), make(chan int)
	go func() {
		for {
			aCh <- 1
			time.Sleep(time.Millisecond * 500)
		}

	}()
	go func() {
		for {
			bCh <- 2
			time.Sleep(time.Millisecond * 500)
		}
	}()

	for {
		select {
		case val := <-aCh:
			log.Println("val: ", val)
		case t := <-bCh:
			log.Println("t: ", t)
		case <-time.After(time.Millisecond * 500):
			log.Println("timeout")
		}
	}
}

func stu4() {
	go func() {
		ticker := time.NewTicker(time.Duration(time.Second * 3))
		for {
			select {
			case <-ticker.C:
				log.Println("select 阻塞")
			}
		}
	}()
	select {
	// default:
	}
}

func stu2_select() {
	// count := 0
	ch := make(chan error)

	go func(c chan error) {
		time.Sleep(time.Second * 1)
		c <- fmt.Errorf("this is a err")
	}(ch)

	for {
		select {
		case err := <-ch:
			if err != nil {
				log.Printf("err: %v\n", err)
			}
			log.Println("success")
		case <-time.After(time.Second * 2):
			log.Printf("2 second timeout\n")
		}
	}

}

func stu1_select() {
	ch := make(chan int)
	quit := make(chan bool)

	go func() { //子go 程获取数据

		for {
			select {

			case <-time.After(3 * time.Second):
				quit <- true
				goto lable
			//return   退出func(),如果用break只会跳出当前case
			//runtime.Goexit()  结束这个goroutine
			case num := <-ch:
				log.Println("num = ", num)
			}
		}
	lable: //lable可以在func（）内任意位置，但不可在函数之外
		log.Println("break to label ---")
	}()

	for i := 0; i < 3; i++ {
		ch <- i
		log.Println("i = ", i)

		time.Sleep(2 * time.Second) //每隔2秒向ch写入数据
	}

	<-quit //主go程阻塞等待子go程通知，退出
	log.Println("finish !!!")
}

func Run(task_id, sleeptime, timeout int, ch chan string) {
	ch_run := make(chan string)
	go run(task_id, sleeptime, ch_run)
	select {
	case re := <-ch_run:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", task_id)
		ch <- re
	}
}

func run(task_id, sleeptime int, ch chan string) {

	time.Sleep(time.Duration(sleeptime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", task_id, sleeptime)
	return
}

func stu3_select() {
	input := []int{3, 2, 1}
	timeout := 2
	chLimit := make(chan bool, 1)
	chs := make([]chan string, len(input))
	limitFunc := func(chLimit chan bool, ch chan string, task_id, sleeptime, timeout int) {
		Run(task_id, sleeptime, timeout, ch)
		<-chLimit
	}
	startTime := time.Now()
	log.Println("Multirun start")
	for i, sleeptime := range input {
		chs[i] = make(chan string, 1)
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleeptime, timeout)
	}

	for _, ch := range chs {
		log.Println(<-ch)
	}
	endTime := time.Now()
	log.Printf("Multissh finished. Process time %s. Number of task is %d\n", endTime.Sub(startTime), len(input))
}
