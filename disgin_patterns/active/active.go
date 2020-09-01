package main

import (
	"log"
	"time"
)

/**
这种模式包含6个组件：
	proxy: 定义了客户端要调用的Active Object接口。当客户端调用它的方法是，方法调用被转换成method request放入到scheduler的activation queue之中。
	method request: 用来封装方法调用的上下文
	activation queue:待处理的 method request队列
	scheduler:一个独立的线程，管理activation queue，调度方法的执行
	servant:active object的方法执行的具体实现，
	future:当客户端调用方法时，一个future对象会立即返回，允许客户端可以获取返回结果。
*/

type MethodRequest int

const (
	Incr MethodRequest = iota
	Decr
)

func Stu_ActiveObject() {
	log.Printf("开始学习Active Object\n")
	simpleExample()
}

type Future interface{}

func simpleExample() {
	printFuture := func(f ...Future) {
		log.Printf("printFuture: %v \n", f)
	}
	s := NewService(10)
	s.Incr()
	s.Decr()
	time.Sleep(time.Second)
	log.Printf("step1: s.Incr(), step2: s.Decr(), v: %d\n", s.v)
	printFuture("I", 0, "U")
	return
}

type Service struct {
	v     int
	queue chan MethodRequest
}

func (s *Service) Incr() {
	s.queue <- Incr
}

func (s *Service) Decr() {
	s.queue <- Decr
}

func (s *Service) schedule() {
	for r := range s.queue {
		log.Printf("<-s.queue: %d\t", s.v)

		if r == Incr {
			s.v++
		} else {
			s.v--
		}
	}
}

func NewService(buffer int) *Service {
	s := &Service{
		queue: make(chan MethodRequest, buffer),
	}
	go s.schedule()
	return s
}

func main() {
	Stu_ActiveObject()
}
