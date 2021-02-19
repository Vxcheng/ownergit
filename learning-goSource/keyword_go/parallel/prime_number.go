package parallel

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

// 返回生成自然数序列的管道: 2, 3, 4, ...
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// 管道过滤器: 删除能被素数整除的数
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func RunPrimeNumer() {
	ch := GenerateNatural()
	// 自然数序列: 2, 3, 4, ...
	for i := 0; i < 100; i++ {
		prime := <-ch // 新出现的素数
		fmt.Printf("%v: %v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) // 基于新素数构造的过滤器
	}
}

func worker(wg *sync.WaitGroup, cannel chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("hello")
		case <-cannel:
			return
		}
	}
}

func RunClose() {
	cancel := make(chan bool)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}

func workerC(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			log.Println("hello")
			time.Sleep(time.Second * 2)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
func RunCloseByContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			if err := workerC(ctx, &wg); err != nil {
				log.Fatalf("err: %#v", err)
			}
		}()
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
