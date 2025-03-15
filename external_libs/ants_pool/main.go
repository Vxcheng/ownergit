package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func myFunc(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("run with %d\n", n)
}

func demoFunc() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("Hello World!")
}

func main() {
	// demo1()
	// demo2()
	demo3()
}

func demo1() {
	defer ants.Release()

	runTimes := 1000

	// Use the common pool.
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		demoFunc()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks.\n")
	select {}
	// Use the pool with a function,
	// set 10 to the capacity of goroutine pool and 1 second for expired duration.
	p, _ := ants.NewPoolWithFunc(10, func(i interface{}) {
		myFunc(i)
		wg.Done()
	})
	defer p.Release()
	// Submit tasks one by one.
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		_ = p.Invoke(int32(i))
	}
	wg.Wait()
	fmt.Printf("running goroutines: %d\n", p.Running())
	fmt.Printf("finish all tasks, result is %d\n", sum)

}

type A struct {
	Items []float64
	DataC chan float64
}

func demo2() {

	defer ants.Release()

	runTimes := 20
	a := &A{
		DataC: make(chan float64, runTimes),
	}
	// Use the common pool.
	syncCalculateSum := func() {
		done(a)
	}
	for i := 0; i < runTimes; i++ {
		_ = ants.Submit(syncCalculateSum)
	}

	for i := 0; i < runTimes; i++ {
		value := <-a.DataC
		a.Items = append(a.Items, value)
	}

	fmt.Printf("running goroutines: %d\n", ants.Running())
	fmt.Printf("finish all tasks, result is %v\n", a.Items)
}

func done(a *A) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := r.Float64()
	a.DataC <- v
}

func demo3() {
	var wg sync.WaitGroup

	// 创建一个容量为 10 的 Goroutine 池
	p, err := ants.NewPool(10)
	if err != nil {
		log.Fatal("goroutine pool create err:", err)
	}
	defer p.Release() // 函数结束后关闭此池并释放工作队列

	// 提交任务到 Goroutine 池
	for i := 0; i < 100; i++ {
		wg.Add(1)
		_ = p.Submit(func() {
			taskFunc()
			wg.Done()
		})
	}

	wg.Wait() // 阻塞等待任务执行完成
	fmt.Println("所有任务已完成")
}

// taskFunc 执行耗时任务
func taskFunc() {
	time.Sleep(1 * time.Second) // 模拟耗时任务
	fmt.Println("任务完成")
}
