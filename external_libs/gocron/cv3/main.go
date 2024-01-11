package main

import (
	"context"
	"fmt"
	"github.com/panjf2000/ants/v2"
	cronV3 "github.com/robfig/cron/v3"
	"golang.org/x/time/rate"
	"sync"
	"time"
)

func main() {
	c := C{
		Cron: cronV3.New(),
	}

	//c.removeDemo()

	//c.LastDemo()
	//c.addDemo()
	//c.sameExp()

	//c.limitDemo1()

	c.limitDemo2()

	//c.limitDemo3()
	//c.lockDemo2()

	//c.WrapSkipIfStillRunning()
}

type C struct {
	*cronV3.Cron // multi entry
}

func (c C) addDemo() (err error) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	_, err = c.Cron.AddFunc("*/1 * * * *", func() {
		fmt.Printf("every minute: %d\n", time.Now().Unix())
	})

	_, err = c.Cron.AddFunc("0 0-18 * * *", func() {
		fmt.Printf("every day: %d\n", time.Now().Unix())
	})

	_, err = c.Cron.AddFunc("8 10 * * 3,4,5,7", func() {
		fmt.Printf("every week: %d\n", time.Now().Unix())
	})

	_, err = c.Cron.AddFunc("0 0,18 1-14 * *", func() {
		fmt.Printf("every month: %d\n", time.Now().Unix())
	})

	// day L-3
	_, err = c.Cron.AddFunc("2-10 1 * * 3", func() {
		now := time.Now()
		lastDay := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, now.Location())
		thirdToLastDay := lastDay.AddDate(0, 0, -2)
		if now.Day() == thirdToLastDay.Day() {
			fmt.Printf("every month: %d\n", time.Now().Unix())
		}
	})

	c.Cron.Start()

	defer c.Stop()

	// Give cron 2 seconds to run our job (which is always activated).
	select {
	case <-time.After(time.Minute * 10):
		println("expected job runs")
	case <-wait(wg):
	}
	println("done")
	return
}

func (c C) removeDemo() (err error) {
	_, err = c.Cron.AddFunc("*/1  * * * *", func() {
		fmt.Printf("every minute: %d\n", time.Now().Unix())
	})
	c.Cron.Start()

	defer c.Stop()
	for {
		if time.Now().Second() > 20 {
			for _, ent := range c.Cron.Entries() {
				c.Cron.Remove(ent.ID)
			}
			break
		}
	}

	return
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
	last := 49
	_, err = c.Cron.AddFunc("* * * * *", func() {
		now := time.Now()
		if now.Minute() == 60-last {
			fmt.Printf("every minute: %d\n", time.Now().Unix())
		} else {
			fmt.Printf("continue every minute: %d\n", time.Now().Unix())
		}

	})
	c.Cron.Start()
	defer c.Stop()

	_, err = c.Cron.AddFunc("*/3 * * * *", func() {
	})
	c.Cron.Start()

	select {}
	return
}

func (c C) sameExp() (err error) {
	_, err = c.Cron.AddFunc("40 * * * *", func() {
		fmt.Printf("continue every minute1: %d\n", time.Now().Unix())
	})
	c.Cron.Start()
	defer c.Stop()

	_, err = c.Cron.AddFunc("* * * * *", func() {
		fmt.Printf("continue every minute2: %d\n", time.Now().Unix())
	})
	c.Cron.Start()

	select {}
	return
}

func do(ent int) {
	fmt.Printf("continue every %d: %v\n", ent, time.Now())
	time.Sleep(time.Second)
}

func (c C) limitDemo1() (err error) {
	p, _ := ants.NewPoolWithFunc(1, func(i interface{}) { // common
		do(i.(int))
	})
	//defer p.Release()

	_, err = c.Cron.AddFunc("* * * * *", func() {
		p.Invoke(1)
	})
	c.Cron.Start()
	defer c.Stop()

	_, err = c.Cron.AddFunc("* * * * *", func() {
		p.Invoke(2)
	})
	c.Cron.Start()

	select {}
	return
}

func do2(ent int, c chan struct{}) {
	c <- struct{}{}
	defer func() {
		<-c
	}()

	fmt.Printf("continue every %d: %d\n", ent, time.Now().Unix())
	time.Sleep(time.Second)
	return
}

func (c C) limitDemo2() (err error) {
	limtC := make(chan struct{}, 1)

	_, err = c.Cron.AddFunc("* * * * *", func() {
		do2(1, limtC)
	})
	c.Cron.Start()
	defer c.Stop()

	_, err = c.Cron.AddFunc("* * * * *", func() {
		do2(2, limtC)
	})
	c.Cron.Start()

	select {}
	return
}

func do3(ent int, c *rate.Limiter) {
	err := c.Wait(context.TODO())
	if err != nil {
		println(err)
		return
	}

	fmt.Printf("continue every %d: %v\n", ent, time.Now())
	time.Sleep(time.Second)
	return
}

func (c C) limitDemo3() (err error) {
	limiter := rate.NewLimiter(1, 1)

	_, err = c.Cron.AddFunc("* * * * *", func() {
		do3(1, limiter)
	})
	c.Cron.Start()
	defer c.Stop()

	_, err = c.Cron.AddFunc("* * * * *", func() {
		do3(2, limiter)
	})
	_, err = c.Cron.AddFunc("* * * * *", func() {
		do3(3, limiter)
	})
	c.Cron.Start()

	select {}
	return
}

func (c C) WrapSkipIfStillRunning() (err error){
	cc := cronV3.New(
		cronV3.WithChain(
			cronV3.SkipIfStillRunning(cronV3.DefaultLogger)))
	//cc.Then()
	_, err = cc.AddFunc("* * * * *", func() {
		do(1)
	})
	cc.Start()
	defer cc.Stop()

	_, err = cc.AddFunc("* * * * *", func() {
		do(2)
	})
	_, err = cc.AddFunc("* * * * *", func() {
		do(3)
	})

	select {}
	return
}

var (
	mu    sync.Mutex  // 互斥锁
	entry cronV3.EntryID // 保留的任务 ID
	c *cronV3.Cron
)

func lockDemo1() {
	c = cronV3.New()
	// 添加任务
	_, _ = c.AddFunc("*/1 * * * * *", func() {
		// 尝试获取互斥锁，如果已被其他任务持有，则直接返回
		entry++
		if !tryLock() {
			return
		}

		// 执行任务
		fmt.Println("Running task")

		// 释放互斥锁
		unlock()
	})

	c.Start()

	// 阻塞主线程
	select {}
}

// 尝试获取互斥锁，如果锁已被其他任务持有，则返回 false
func tryLock() bool {
	entry++

	mu.Lock()
	defer mu.Unlock()

	// 如果锁已被持有，则返回 false
	if entry != 0 {
		return false
	}

	// 记录当前任务 ID
	//entry = c.CurrentEntry().ID
	return true
}

// 释放互斥锁
func unlock() {
	mu.Lock()
	defer mu.Unlock()

	// 清除任务记录
	entry = 0
}

func (c C) lockDemo2() (err error){
	limiter := make(chan struct{}, 1)

	_, err = c.Cron.AddFunc("* * * * *", func() {
		select {
		case limiter<- struct{}{}:
		default:
			return
		}

		do(1)
		<-limiter
	})
	c.Cron.Start()
	defer c.Stop()

	_, err = c.Cron.AddFunc("* * * * *", func() {
		select {
		case limiter<- struct{}{}:
		default:
			return
		}

		do(2)
		<-limiter
	})
	_, err = c.Cron.AddFunc("* * * * *", func() {
		select {
		case limiter<- struct{}{}:
		default:
			return
		}

		do(3)
		<-limiter
	})

	select {}
	return
}
