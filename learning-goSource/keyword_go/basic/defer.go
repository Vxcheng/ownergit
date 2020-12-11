package basic

import (
	"fmt"
	"log"
)

func Stu_defer() {
	log.Println("learn defer")
	stu1_defer()
	stu2_defer()
	stu3_defer()
}

type Person struct {
	age int
}

func stu1_defer() {
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

func stu2_defer() {
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

type nodeStatus struct {
	ip  string
	err error
}

type statusErr struct{}

func (s *statusErr) Error() string {
	return ""
}

func stu3_defer() {
	{
		value := 0
		defer func() {
			log.Printf("value: %d", value)
		}()
		value = 1
	}

	{

		status := &nodeStatus{
			ip: "127.0.0.1",
		}
		var err error
		defer func(s *nodeStatus) {
			log.Printf("status: %+v", s)
		}(status)
		err = fmt.Errorf("unknown")
		status.err = err
	}

	{
		str, err := stu3_defer_0()
		log.Printf("str: %s, err: %v", str, err)

	}
}

func stu3_defer_0() (str string, err error) {
	defer func() {
		str, err = "hello", nil
	}()
	return "world", fmt.Errorf("world")
}
