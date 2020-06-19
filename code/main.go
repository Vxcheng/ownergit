package main

import (
	"fmt"
	"time"
)

type Semaphore chan struct{}

func NewSemaphore(size int) Semaphore {
	return make(Semaphore, size)
}

func (s Semaphore) Lock() {
	s <- struct{}{}
}

func (s Semaphore) Unlock() {
	<-s
}

type Mutex Semaphore

func NewMutex() Mutex {
	return Mutex(NewSemaphore(1)) // signal
}

// func main() {
// 	count()
// }

func count() {
	// m := NewMutex()

	var sum int
	for i := 0; i < 5; i++ {
		go func() {
			// Semaphore(m).Lock()
			// defer Semaphore(m).Unlock()

			value := sum
			value++
			time.Sleep(1 * time.Nanosecond)
			sum = value
			fmt.Printf("num+: %d \n", sum)
		}()
	}

	time.Sleep(time.Millisecond * 10)
	fmt.Printf("sum: %d \n", sum)
}
