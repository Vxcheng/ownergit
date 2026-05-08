package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*
12.1 开启100个协程，顺序打印1-1000，且保证协程号1的，打印尾数为1的数字
12.2 三个goroutinue交替打印abc 10次
12.3 用不超过10个goroutine不重复的打印slice中的100个元素
12.4 两个协程交替打印奇偶数
12.5 用单个channel实现0,1的交替打印
12.6 sync.Cond实现多生产者多消费者
12.7 使用go实现1000个并发控制并设置执行超时时间1秒
12.8 使用两个Goroutine,向标准输出中按顺序按顺序交替打出字母与数字,输出是a1b2c3
12.9 编写一个程序限制10个goroutine执行,每执行完一个goroutine就放一个新的goroutine进来
*/

// 是一个生产者消费者的题目，就是10个消费者去消费一个数组100个元素，然后把结果再累加回去
func producerComsumer1() {
	nums := make([]int, 100)
	for i := 0; i < len(nums); i++ {
		nums[i] = i + 1
	}

	total := parallelSum(nums, 10)
	fmt.Printf("最终累加结果：%d\n", total)
}

func parallelSum(nums []int, consumerCount int) int64 {
	jobs := make(chan int, len(nums))
	var wg sync.WaitGroup
	var total int64

	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for n := range jobs {
				atomic.AddInt64(&total, int64(n))
			}
		}(i + 1)
	}

	go func() {
		for _, n := range nums {
			jobs <- n
		}
		close(jobs)
	}()

	wg.Wait()
	return total
}

func producerComsumer2(nums []int) int {
	taskCh := make(chan int, len(nums))
	resultCh := make(chan int, len(nums))

	// 1. 生产者：将数组元素发送到任务队列
	go func() {
		for _, n := range nums {
			taskCh <- n
		}
		close(taskCh)
	}()

	// 2. 消费者：10个协程并发消费
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for n := range taskCh {
				// 模拟处理：这里就是累加，实际可以是复杂计算
				resultCh <- n
			}
		}(i)
	}

	// 3. 等待所有消费者完成后关闭结果通道
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 4. 汇总结果
	sum := 0
	for r := range resultCh {
		sum += r
	}
	return sum
}
