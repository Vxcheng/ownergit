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

type s struct {
	a int
	b string
	c []int
	d func()
	e interface{}
}

func unsafeFmt() {
	a := int(1)
	println(unsafe.Alignof(a))
	println(unsafe.Sizeof(a))
	//
	s := s{}
	fmt.Println(unsafe.Sizeof(s), unsafe.Alignof(s))
	fmt.Println(unsafe.Sizeof(s.a), unsafe.Alignof(s.a))
	fmt.Println(unsafe.Sizeof(s.b), unsafe.Alignof(s.b))
	fmt.Println(unsafe.Sizeof(s.c), unsafe.Alignof(s.c))
	fmt.Println(unsafe.Sizeof(s.d), unsafe.Alignof(s.d))
	fmt.Println(unsafe.Sizeof(s.e), unsafe.Alignof(s.e))

	p := (*int16)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.a)))
	fmt.Println(p, s.a)
	*p = 42
	fmt.Println(s.a)
}

func float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}
