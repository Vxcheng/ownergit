package basic

import (
	"fmt"
	"log"
)

func Stu_defer() {
	log.Println("learn defer")
	// stu1_defer()
}

type Person struct {
	age int
}

func deferStu1() {
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

func deferStu2() {
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

func deferStu3() {
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

}

func deferStu4() (str string, err error) {
	defer func() {
		str, err = "hello", nil
	}()
	return "world", fmt.Errorf("world")
}

func deferInc() (v int) {
	defer func() { v++ }()
	return 42
}

func deferPrint() {
	for i := 0; i < 3; i++ {
		defer func() {
			println(i)
		}()
	}

}

func deferPrint1() {
	for i := 0; i < 3; i++ {
		i := i // 定义一个循环体内局部变量i
		defer func() {
			println(i)
		}()
	}
}

func deferPrint2() {
	for i := 0; i < 3; i++ { // 通过函数传入i // defer 语句会马上对调用参数求值
		defer func(i int) {
			println(i)
		}(i)
	}
}

func returnButDefer() (t int) { //t初始化0， 并且作用域为该函数全域

	defer func() {
		t = t * 10
	}()

	return 1
}

func function(index int, value int) int {

	fmt.Println(index)

	return index
}

func stackPush() {
	defer function(1, function(3, 0))
	defer function(2, function(4, 0))
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}

func DeferFunc4() (t int) {
	defer func(i int) {
		fmt.Println(i)
		fmt.Println(t)
	}(t)
	t = 1
	return 2
}

func example1() {
	fmt.Println(DeferFunc1(1))
	fmt.Println(DeferFunc2(1))
	fmt.Println(DeferFunc3(1))
	DeferFunc4()
}
