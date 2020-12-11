package main

import (
	"fmt"
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
