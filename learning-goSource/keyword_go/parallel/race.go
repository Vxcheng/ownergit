package parallel

import (
	"fmt"
	"sync"
)

func race_stu1() {
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Println(count)
			count++

			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Printf("out: %d\n", count)
	return
}

func race_stu2() {
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("num: %d, count: %d\n", i, count)
			count++

			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Printf("out: %d\n", count)
	return
}

func race_stu3() {
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(num int) {
			fmt.Printf("num: %d, count: %d\n", num, count)
			count++

			wg.Done()
		}(i)

	}
	wg.Wait()
	fmt.Printf("out: %d\n", count)
	return
}
