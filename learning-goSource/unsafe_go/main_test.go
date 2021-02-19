package main

import (
	"fmt"
	"testing"
)

func TestUnSafePoint(t *testing.T) {
	t.Run("UnSafePoint", func(t *testing.T) {
		UnSafePoint()
	})
}

func TestUnsafeFmt(t *testing.T) {
	t.Run("", func(t *testing.T) {
		unsafeFmt()
		fmt.Printf("%#016x\n", float64bits(1.0))
	})
}
