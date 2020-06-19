package main

import (
	"fmt"
	"unsafe"
)

func Add1(num int) int {
	var r int
	r = num + 1
	return r
}

func SubNoPoint(num int) {
	num -= 1
}

func Sub1(num *int) {
	*num--
}

type User struct {
	Name string
	Sex  bool
	Age  int16
}

type Phone struct{}

var sexCollection = map[bool]string{true: "boy", false: "girl"}

func (u User) TalkName() {
	fmt.Printf("hi, my name is %s\n", u.Name)
	u.Name = "unknow"
	// return nil
}

func (u *User) SaySex() error {
	fmt.Printf("i am a %s\n", sexCollection[u.Sex])
	return nil
}

func (u *User) Growing() {
	fmt.Printf("u point: %v\n", u)
	fmt.Printf("u derefrence: %v\n", *u)
	p := *u
	fmt.Printf("p point: %v\n", p)

	fmt.Printf("i am %d, i want growing to %d!\n", (u).Age, (*u).Age+10)
	(*u).Age += 10
}

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
