package parallel

import (
	"log"
	"runtime"
	"sync"
	// "time"
)

func Stu_go() {
	log.Println("学习go关键字")
	stu1_go()
}

func stu1_go() {
	num := runtime.GOMAXPROCS(runtime.NumCPU())
	log.Printf("进程数：%d.", num)

	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(i int) {
			log.Printf("%d, ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
