package main

import "testing"

/*
性能差异原理
缓存工作机制
CPU 缓存以 缓存行（通常64字节） 为单位加载数据

连续内存访问可以充分利用缓存行

跳跃式访问会导致大量 缓存缺失

bar() 的性能问题
缓存颠簸：每次访问 matrixB[j][i] 都在不同的缓存行

预取器失效：CPU 无法预测跳跃式访问模式

TLB 缺失：频繁跨越不同内存页

缓存污染：加载的缓存数据很少被重复使用

BenchmarkFoo-8               277          24690552 ns/op        65634311 B/op       4002 allocs/op
BenchmarkBar-8                86          67162294 ns/op        65634376 B/op       4002 allocs/op
*/

func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo()
	}

}

func BenchmarkBar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bar()
	}
}

/*
无符号整数下溢
*/

func TestUnderflow(t *testing.T) {
	cases := []struct {
		x, y, want uint
	}{
		{1, 2, 2 ^ 64 - 1},
		{1, 10, 10 ^ 64 - 9},
	}
	a := 2 ^ 64
	println(a)
	for _, c := range cases {
		t.Run("", func(t *testing.T) {
			println(c.x - c.y)
			if got := c.x - c.y; got != c.want {
				t.Errorf("x - y = %d, want %d", got, c.want)
			}
		})
	}

}
