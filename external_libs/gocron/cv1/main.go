package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron"
)

func main() {
	c := C{
		Cron: cron.New(),
	}

	c.addDemo()

	//c.LastDemo()
}

type C struct {
	*cron.Cron // multi entry
}

func (c C) addDemo() (err error) {
	fmt.Printf("now: %d\n", time.Now().Unix())

	wg := &sync.WaitGroup{}
	wg.Add(1)

	err = c.Cron.AddFunc("@every 1m", func() {
		fmt.Printf("@every minute: %d\n", time.Now().Unix())
	})

	err = c.Cron.AddFunc("* * * * * *", func() {
		fmt.Printf("every minute: %d\n", time.Now().Unix())
	})

	err = c.Cron.AddFunc("0 0-18 * * *", func() {
		fmt.Printf("every day: %d\n", time.Now().Unix())
	})

	err = c.Cron.AddFunc("0 12 * * 1,2", func() {
		fmt.Printf("every week: %d\n", time.Now().Unix())
	})

	err = c.Cron.AddFunc("0 0,18 1-14 * *", func() {
		fmt.Printf("every month: %d\n", time.Now().Unix())
	})

	err = c.AddFunc("0 0 23 L * ?", func() {
		fmt.Println("Task executed on the second-to-last day of the month")
	})
	c.Cron.Start()

	defer c.Stop()

	// Give cron 2 seconds to run our job (which is always activated).
	select {
	case <-time.After(time.Minute * 5):
		println("expected job runs")
	case <-wait(wg):
	}
	println("done")
	return
}

func (c C) removeDemo() {

}

func wait(wg *sync.WaitGroup) chan bool {
	ch := make(chan bool)
	go func() {
		wg.Wait()
		ch <- true
	}()
	return ch
}

func (c C) LastDemo() (err error) {
	//last := 12
	c.Cron.AddFunc("* * * * * *", func() {
		fmt.Printf("every month: %d\n", time.Now().Unix())
	})
	c.Cron.Start()
	defer c.Stop()

	select {}
	return
}
