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
	// outFunc()
	// outChan()
	// stu1_context()
	stu2_context()
	log.Println("finish")
}

func stu2_context() {
	expired := time.Second * 1
	{
		// inner
		ctx, cancel := context.WithTimeout(context.Background(), expired)
		f := func(c context.Context) {
			log.Println("queue")
			time.Sleep(expired * 2)
		in:
			for {
				select {
				case <-ctx.Done():
					log.Println("time out")
					break in
				}
			}
		}
		f(ctx)
		cancel()
	}

	{
		// outer
		ctx, cancel := context.WithTimeout(context.Background(), expired)
		f := func(c context.Context) {
			log.Println("queue")
			time.Sleep(expired * 2)
		}
		f(ctx)
	out:
		for {
			select {
			case <-ctx.Done():
				log.Println("time out")
				break out
			}
		}
		cancel()
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), expired)
		go func(c context.Context) {
			log.Println("parallel")
			time.Sleep(time.Second * 2)

		in:
			for {
				select {
				case <-ctx.Done():
					log.Println("time out")
					cancel()
					break in
				}
			}

		}(ctx)
		time.Sleep(time.Second * 3)
		log.Println("wait")
	}

	{
		ctx, cancel := context.WithTimeout(context.Background(), expired)
		go func(c context.Context) {
			log.Println("parallel")
			time.Sleep(time.Second * 2)
		}(ctx)

	out_p:
		for {
			select {
			case <-ctx.Done():
				log.Println("time out")
				cancel()
				break out_p
			}
		}
	}
}

func stu1_context() {

	collect := func(ctx context.Context, item int, errC chan<- error) {
		var err error
		time.Sleep(time.Millisecond * 500) //
		if item%2 != 1 {                   // yu
			err = fmt.Errorf("mock error, i: %d", item)
		} else {
			log.Printf("do collect %d", item)
			err = nil
		}
		errC <- err
	}

	{

		f := func(ctx context.Context, parallel bool, sli []int) error {
			var errInfo string
			errC := make(chan error, len(sli))
			for _, item := range sli {
				if parallel {
					go func() {
						collect(ctx, item, errC)
					}()
				} else {
					collect(ctx, item, errC)
				}
			}

			for i := 0; i < len(sli); i++ {
				select {
				case <-ctx.Done():
					return fmt.Errorf("time out")
				case err := <-errC:
					if err != nil {
						errInfo += fmt.Sprintf("%s, ", err.Error())
					}
				}
			}

			if errInfo != "" {
				return fmt.Errorf(errInfo)
			}
			return nil
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		parallel := false
		items := []int{1, 2, 3}
		err := f(ctx, parallel, items)
		if err != nil {
			log.Printf("collect failed, err: %v", err)
		}
		cancel()
	}
}

func outChan() {
	ch := make(chan int)
	go func(c chan int) {
		time.Sleep(time.Second * 2)
		c <- 1
	}(ch)
	for {
		select {
		case <-time.After(4 * time.Second):
			log.Println("time out")
		case a := <-ch:
			log.Printf("a: %d\n", a)
			return
		}
	}
}

func outFunc() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	go func(c context.Context) {
		time.Sleep(time.Second * 4)
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
	defer cancel()
	chiHanBao(ctx)
	exec(ctx)
}

func chiHanBao(ctx context.Context) {
	n := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context stop ")
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
	time.Sleep(time.Second * 4)
	log.Println("sleep....")
}
