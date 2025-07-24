package main

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		Demo()
	})

	t.Run("test", func(t *testing.T) {
		Demo2()
	})
}

func Demo() {
	num := 5
	// 不好的做法
	for i := 0; i < num; i++ {
		defer fmt.Println("d1: ", i) // 会累积100个延迟调用
	}

	fmt.Println("--------------------------------")
	// 改进做法
	for i := 0; i < num; i++ {
		func(i int) {
			defer fmt.Println("d2: ", i) // 每个循环迭代独立的defer
			fmt.Println("d2 print: ", i)
		}(i)
	}
}

// defer 内嵌套多个defer， 生成Demo2函数
func Demo2() {
	var i int
	defer func() {
		defer func() {
			i++
			fmt.Println("i1: ", i)
		}()
		defer func(n int) {
			n++
			fmt.Println("i2: ", n, i)
		}(i)
		i++
		fmt.Println("s2: ", i)
	}()
	i++
	fmt.Println("s1: ", i)

}
