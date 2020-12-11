package main

import (
	"fmt"
	"unsafe"
)

func UnSafePoint() {
	u := uint32(32)
	i := int32(1)
	fmt.Printf("u: %v, i: %v\n", &u, &i)

	p := &i
	p = (*int32)(unsafe.Pointer(&u))
	fmt.Printf("p: %v\n", p)

	b := []byte{'a', 'b', 'c'}
	c := &b[0]
	fmt.Printf("b: %v, c: %v\n", b, c)
	a := unsafe.Pointer(c)
	a1 := uintptr(a)
	a2 := unsafe.Pointer(a1 + uintptr(1))
	fmt.Printf("a: %v, a1: %v, a2: %v\n", a, a1, a2)
	a3 := (*byte)(a2)
	fmt.Printf("after: %v, %v, %v\n", a3, *a3, &a3)
}

func main() {
	UnSafePoint()
}
