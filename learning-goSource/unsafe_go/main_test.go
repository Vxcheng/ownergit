package main

import (
	"fmt"
	"testing"
	"unsafe"
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

	t.Run("unsafeMap", func(t *testing.T) {
		unsafeMap()
	})

	t.Run("unsafeStruct", func(t *testing.T) {
		unsafeStruct()
	})
}

type demo1 struct {
	a int8
	b int16
	c int32
}

type demo2 struct {
	a int8
	c int32
	b int16
}

type demo3 struct {
	c int32
	a struct{}
}

type demo4 struct {
	a struct{}
	c int32
}

func TestUnsafeAlign(t *testing.T) {
	t.Run("demo1", func(t *testing.T) {
		fmt.Printf("demo1: %d\n", unsafe.Sizeof(demo1{}))
		fmt.Printf("demo1: %d\n", unsafe.Alignof(demo1{}))

		fmt.Printf("demo2: %d\n", unsafe.Sizeof(demo2{}))
		fmt.Printf("demo2: %d\n", unsafe.Alignof(demo2{}))
	})

	t.Run("demo4", func(t *testing.T) {
		fmt.Printf("demo3: %d\n", unsafe.Sizeof(demo3{}))
		fmt.Printf("demo3: %d\n", unsafe.Alignof(demo3{}))

		fmt.Printf("demo4: %d\n", unsafe.Sizeof(demo4{}))
		fmt.Printf("demo4: %d\n", unsafe.Alignof(demo4{}))
	})

}
