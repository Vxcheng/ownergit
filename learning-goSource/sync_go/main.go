package main

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	poolDemo()
	atomicDemo1()
	atomicDemo2()
	// stu_waitGroup()
	// stu1_lock()
	time.Sleep(time.Second)
	log.Println("finished")
}

type User struct {
	wg  sync.WaitGroup
	mu  sync.Mutex
	rw  sync.RWMutex
	num int
}

func stu_waitGroup() {
	u := &User{}
	nums := 3
	for i := 0; i < nums; i++ {
		u.wg.Add(1)

		go func() {
			defer u.wg.Done()
			u.Count()
		}()
	}
	u.wg.Wait()
	log.Println("num:", u.num)
	return
}

func (u *User) Count() {
	u.num++
}

func stu1_lock() {
	nums := 5
	{
		count := 0
		for i := 0; i < nums; i++ {
			go func() {
				count++
				log.Printf("iner count, ponit: %v, value: %d\n", &count, count)

			}()
			log.Printf("out count: %d\n", count)
		}
		log.Printf("final count: %d\n", count)
	}

	{
		count := 0
		for i := 0; i < nums; i++ {
			go func() {
				count++
				log.Printf("iner count, ponit: %v, value: %d\n", &count, count)
			}()
			log.Printf("out count: %d\n", count)
		}
		time.Sleep(time.Millisecond * 200) //
		log.Printf("final count: %d\n", count)
	}

	{
		count := 0
		for i := 0; i < nums; i++ {
			go func(val *int) {
				*val++
				log.Printf("iner count, ponit: %v, value: %d\n", val, *val)
			}(&count)
			log.Printf("out count: %d\n", count)
		}
		log.Printf("final count: %d\n", count)
	}

	{
		u := &User{}
		for i := 0; i < nums; i++ {
			u.mu.Lock()
			go func() {
				u.Count()
				log.Printf("iner num: %d\n", u.num)
			}()
			log.Printf("out num: %d\n", u.num)
			u.mu.Unlock()
		}

		log.Printf("final u.num: %d\n", u.num)
	}
}

func atomicDemo1() {
	// cas
	v := int64(0)
	atomic.StoreInt64(&v, 1)
	atomic.LoadInt64(&v)
	atomic.CompareAndSwapInt64(&v, 2, 3)
	atomic.AddInt64(&v, 1)
	atomic.SwapInt64(&v, 1)
	atomic.AndInt64(&v, 1)
	atomic.OrInt64(&v, 1)
}
func atomicDemo2() {
	// cas
	a := atomic.Int64{}
	a.Store(1)
	a.Load()
	a.CompareAndSwap(1, 2)
	a.Add(1)
	a.Swap(1)
	a.And(1)
	a.Or(1)
}

// 定义一个临时对象的结构
type TempObject struct {
	data string
}

// 初始化 sync.Pool
var pool = sync.Pool{
	New: func() interface{} {
		return &TempObject{}
	},
}

func poolDemo() {
	const numJobs = 100
	var wg sync.WaitGroup

	// 启动多个并发任务
	for i := 0; i < numJobs; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// 从池中获取一个对象
			obj := pool.Get().(*TempObject)
			obj.data = fmt.Sprintf("Job %d", i)

			// 模拟处理任务
			fmt.Println("Processing", obj.data)
			time.Sleep(time.Duration(i) * time.Millisecond)

			// 将对象放回池中
			pool.Put(obj)
		}(i)
	}

	// 等待所有任务完成
	wg.Wait()
	fmt.Println("All jobs completed")
}
