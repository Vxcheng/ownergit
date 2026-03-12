package main

import (
	"fmt"
	"sync"
	"time"
)

type MapRW interface {
	Read(key string) string
	Write(key, value string)
}

type MyMap struct {
	data     map[string]string
	dataLock map[string]chan bool
	lock     sync.RWMutex
}

func NewMyMap() *MyMap {
	return &MyMap{data: make(map[string]string), dataLock: make(map[string]chan bool)}
}

func (m *MyMap) Read(key string) string {
	m.lock.Lock()
	if v, ok := m.data[key]; ok {
		m.lock.Unlock()
		return v
	}

	if ch, ok := m.dataLock[key]; ok {
		m.lock.Unlock()
		<-ch
		return m.Read(key)
	} else {
		ch := make(chan bool)
		m.dataLock[key] = ch
		m.lock.Unlock()
		<-ch
		return m.Read(key)
	}
}

func (m *MyMap) Write(key, value string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.data[key] = value
	if ch, ok := m.dataLock[key]; ok {
		close(ch)
	}
}

func main() {
	m := NewMyMap()
	m.Write("ni", "hao")
	var wg sync.WaitGroup
	wg.Add(4)
	go func() {
		fmt.Println("begin goroutine 1")
		m.Read("hello")
		fmt.Println("finish goroutine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("begin goroutine 2")
		m.Read("hello")
		fmt.Println("finish goroutine 2")
		wg.Done()
	}()
	go func() {
		fmt.Println("begin goroutine 3")
		m.Read("ni")
		fmt.Println("finish goroutine 3")
		wg.Done()
	}()
	go func() {
		time.Sleep(10 * time.Second)
		m.Write("hello", "world")
		wg.Done()
	}()
	wg.Wait()
}
