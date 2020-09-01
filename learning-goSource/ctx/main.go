package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	log.Println("start")
	WithTimeout()
	log.Println("finish")
}

func WithTimeout() {
	fmt.Printf("%03d", 2)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10))

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// chiHanBao(ctx)
	exec(ctx)
	defer cancel()
}

func chiHanBao(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stop \n")
			return
		default:
			incr := rand.Intn(5)
			n += incr
			fmt.Printf("我吃了 %d 个汉堡\n", n)
		}
		time.Sleep(time.Second)
	}
}

func exec(ctx context.Context) {
	time.Sleep(time.Second * 5)
}
