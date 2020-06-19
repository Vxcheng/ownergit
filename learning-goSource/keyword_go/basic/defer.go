package basic

import (
	"fmt"
	"log"
)

func Stu_defer() {
	p := &Person{28}

	defer log.Printf("a: %d\n", p.age)
	defer func(person *Person) {
		log.Printf("b: %d\n", person.age)
	}(p)
	defer func() {
		log.Printf("c: %d\n", p.age)
	}()
	p.age = 29
	return
}

type Person struct {
	age int
}

func stu1() {
	p := &Person{28}

	defer log.Printf("a1: %d\n", p.age)
	defer func(person *Person) {
		log.Printf("b1: %d\n", person.age)
	}(p)
	defer func() {
		log.Printf("c1: %d\n", p.age)
	}()
	p = &Person{29}
	return
}

func init() {
	stu2()
}

func stu2() {
	str, err := stu3()
	log.Printf("str: %s, err: %v\n", str, err)
}

func stu3() (str string, err error) {
	defer func() {
		str, err = "hello", nil
	}()
	return "world", fmt.Errorf("world")
}
