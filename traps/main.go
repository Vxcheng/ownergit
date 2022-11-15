package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

var (
	G *A
)

func main() {
	// int overflow
	overflow(2147483647)
	if G != nil {
		println("true")
	} else {
		println("false")
	}

	//SF()

	//unsafePointer()
}

type A struct {
	B  *B
	B1 B
	V  string
}

type B struct {
	V string
}

// bad: 未限制长度，导致整数溢出
func overflow(numControlByUser int32) {
	var numInt int32 = 0
	numInt = numControlByUser + 1
	// 对长度限制不当，导致整数溢出
	fmt.Printf("%d\n", numInt)
	// 使用numInt，可能导致其他错误
}

type Data struct {
	o *Data
}

// bad
func SetFinal() {
	var a, b Data
	a.o = &b
	b.o = &a
	// 指针循环引用，SetFinalizer()无法正常调用
	runtime.SetFinalizer(&a, func(d *Data) {
		fmt.Printf("a %p final.\n", d)
		fmt.Println("a")
	})
	runtime.SetFinalizer(&b, func(d *Data) {
		fmt.Printf("b %p final.\n", d)
		fmt.Println("b")
	})
}

func SF() {
	for {
		SetFinal()
		time.Sleep(time.Millisecond)
	}
}

func unsafePointer() {
	b := make([]byte, 1)
	foo := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(0xfffffffe)))
	fmt.Print(*foo + 1)
}
