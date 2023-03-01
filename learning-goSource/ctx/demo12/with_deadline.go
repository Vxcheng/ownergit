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

	ctx2, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second))

	defer cancel()

	wg.Add(1)

	go func() {

		defer wg.Done()

		tick := time.NewTicker(300 * time.Millisecond)

		for {

			select {

			case <-ctx2.Done():

				fmt.Println(ctx2.Err())

				return

			case t := <-tick.C:

				fmt.Println(t.Nanosecond())

			}

		}

	}()

	wg.Wait()

}
