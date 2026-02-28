package main

import (
	"fmt"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// main_count()
	// main_token()
	main_limiter()
}

// FixedWindowLimiter 实现了固定窗口算法的限流器
type FixedWindowLimiter struct {
	limit      int64 // 每秒最大请求数
	counter    int64 // 当前计数
	lastSecond int64 // 上一次重置的时间戳（秒）
}

func NewFixedWindowLimiter(limit int64) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:      limit,
		lastSecond: time.Now().Unix(),
	}
}

// Allow 判断是否允许通过
func (l *FixedWindowLimiter) Allow() bool {
	now := time.Now().Unix()
	// 如果跨秒，重置计数器（乐观锁思想）
	if now > atomic.LoadInt64(&l.lastSecond) {
		// 尝试更新 lastSecond，只有更新成功的线程负责重置计数器
		if atomic.CompareAndSwapInt64(&l.lastSecond, l.lastSecond, now) {
			fmt.Println("重置计数器", now, "旧值", l.lastSecond, "计数器值", l.counter)
			atomic.StoreInt64(&l.counter, 0)
		}
	}
	// 计数+1
	current := atomic.AddInt64(&l.counter, 1)
	return current <= l.limit
}

func main_count() {
	limiter := NewFixedWindowLimiter(1000)

	// 模拟并发请求
	for i := 0; i < 2000; i++ {
		go func(i int) {
			time.Sleep(time.Duration(i%1000) * time.Millisecond)
			if limiter.Allow() {
				fmt.Printf("请求 %d: ✅ 通过\n", i)
			} else {
				fmt.Printf("请求 %d: ❌ 限流\n", i)
			}
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("请求结束", limiter.counter)
}

// TokenBucketLimiter 实现了令牌桶算法的限流器
type TokenBucketLimiter struct {
	rate     int64         // 每秒放入令牌数
	capacity int64         // 桶容量
	tokens   int64         // 当前令牌数
	lastFill time.Time     // 上次填充时间
	fillChan chan struct{} // 触发填充的信号（可选）
	stopChan chan struct{}
}

func NewTokenBucketLimiter(rate, capacity int64) *TokenBucketLimiter {
	tb := &TokenBucketLimiter{
		rate:     rate,
		capacity: capacity,
		tokens:   capacity, // 初始满桶
		lastFill: time.Now(),
		stopChan: make(chan struct{}),
	}
	// 启动后台定时填充
	go tb.startFiller()
	return tb
}

func (tb *TokenBucketLimiter) startFiller() {
	ticker := time.NewTicker(time.Second / time.Duration(tb.rate))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			tb.fill()
		case <-tb.stopChan:
			return
		}
	}
}

func (tb *TokenBucketLimiter) fill() {
	now := time.Now()
	elapsed := now.Sub(tb.lastFill).Seconds()
	tb.lastFill = now

	// 计算应该添加的令牌数
	newTokens := int64(elapsed * float64(tb.rate))
	if newTokens > 0 {
		tb.tokens += newTokens
		if tb.tokens > tb.capacity {
			tb.tokens = tb.capacity
		}
	}
}

// Allow 尝试获取一个令牌
func (tb *TokenBucketLimiter) Allow() bool {
	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucketLimiter) Stop() {
	close(tb.stopChan)
}

func main_token() {
	limiter := NewTokenBucketLimiter(1000, 1000) // 每秒1000个，桶容量1000
	defer limiter.Stop()

	for i := 0; i < 2000; i++ {
		go func(i int) {
			if limiter.Allow() {
				fmt.Printf("请求 %d: ✅ 通过\n", i)
			} else {
				fmt.Printf("请求 %d: ❌ 限流\n", i)
			}
		}(i)
	}
	time.Sleep(2 * time.Second)
}

func main_limiter() {
	// 每秒产生 1000 个令牌，桶容量 1000（允许瞬时 1000 并发）
	limiter := rate.NewLimiter(1000, 1000)
	count := 0
	for i := 0; i < 2000; i++ {
		go func(i int) {
			if limiter.Allow() {
				fmt.Printf("请求 %d: ✅ 通过\n", i)
				count++
			} else {
				fmt.Printf("请求 %d: ❌ 限流\n", i)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("总通过请求数:", count)
}
