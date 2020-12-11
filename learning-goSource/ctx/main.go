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
	// WithTimeout()
	outFunc()
	log.Println("finish")
}

func outChan() {
	ch := make(chan int)
	go func(c chan int) {
		time.Sleep(time.Second * 2)
		log.Println("out")
	}(ch)
	for {
		select {
		case <-time.After(4 * time.Second):
		}
	}

}

func outFunc() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	go func(c context.Context) {
		time.Sleep(time.Second * 6)
		log.Println("out")
	}(ctx)
	select {
	case <-ctx.Done():
	}
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
			fmt.Println("stop ")
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
