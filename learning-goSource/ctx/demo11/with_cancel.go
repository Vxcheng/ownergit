package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	ctx := context.Background()

	ctx1, cancel := context.WithCancel(ctx)

	wg.Add(1)

	go func() {

		defer wg.Done()

		tick := time.NewTicker(300 * time.Millisecond)

		for {

			select {

			case <-ctx1.Done():

				fmt.Println(ctx1.Err())

				return

			case t := <-tick.C:

				fmt.Println(t.Nanosecond())

			}

		}

	}()

	time.Sleep(time.Second)

	cancel()

	wg.Wait()

}
