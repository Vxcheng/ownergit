package main

import (
	"fmt"
	"testing"
)

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	t1
	t2

	mutexWaiterShift = iota
)

func TestMutex(t *testing.T) {
	t.Run("", func(t *testing.T) {
		fmt.Printf("%x, %x, %x, %x, %x, %x\n", mutexLocked, mutexWoken, mutexStarving, mutexWaiterShift, t1, t2)
	})
}
