package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	// stu_waitGroup()
	stu1_lock()
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
