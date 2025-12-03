package main

import "testing"

/*
有限并行，在资源有限的情况下完成大量独立任务
该代码用典型的有界并发+管道+取消+同步等待模式，安全高效地并发处理目录下所有文件的哈希计算。

通用规则（总结）

谁发送谁关闭：通常由 channel 的单一发送方负责关闭；不要让接收方去 close。
只有当接收方使用 for ... range 或需要“无更多值”的信号时，才必须 close。
关闭一个可能仍有发送者的 channel 会导致 panic。用 WaitGroup /其它同步确保没有并发发送后再 close。
信号型 channel（只作广播）用 close 作为广播是惯用做法（如 done）。
*/
func TestBounded_parallelismn(t *testing.T) {
	bounded_parallelism()
}

/**
并行性	完成大量独立任务

*/
func TestParallelism(t *testing.T) {
	parallelism()
}

/*
生成器每次产生一个值序列。
*/
func TestGenerator(t *testing.T) {
	generator()

}

func TestFanin(t *testing.T) {
	fanInMain()

}
func TestFanout(t *testing.T) {
	fanOutMain()
}
