package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// main_count()
	// main_token()
	// main_limiter()
	// main_sliding()
	main_redis_limit()
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

// SlidingWindowLog 基于时间戳队列的精确滑动窗口限流器
//
// 原理：用一个大小为 limit 的循环队列记录每个通过请求的时间戳。
// 每次 Allow 时，先丢弃队列中早于 (now - window) 的旧时间戳，
// 再判断剩余数量是否已达上限。精确但内存占用与 limit 成正比。
type SlidingWindowLog struct {
	limit      int     // 窗口内最大请求数
	window     int64   // 窗口大小（纳秒）
	timestamps []int64 // 循环队列，存储各请求时间戳（UnixNano）
	head       int     // 队列头（最旧元素）索引
	count      int     // 当前窗口内有效请求数
	mu         sync.Mutex
}

func NewSlidingWindowLog(limit int, window time.Duration) *SlidingWindowLog {
	return &SlidingWindowLog{
		limit:      limit,
		window:     window.Nanoseconds(),
		timestamps: make([]int64, limit),
	}
}

func (l *SlidingWindowLog) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now().UnixNano()
	cutoff := now - l.window

	// 移除窗口外的过期时间戳
	for l.count > 0 && l.timestamps[l.head] <= cutoff {
		l.head = (l.head + 1) % l.limit
		l.count--
	}

	if l.count >= l.limit {
		return false
	}

	// 写入本次请求时间戳
	tail := (l.head + l.count) % l.limit
	l.timestamps[tail] = now
	l.count++
	return true
}

// SlidingWindowCounter 基于计数器加权的近似滑动窗口限流器
//
// 原理：同时维护"当前窗口"和"上一个窗口"的计数器。
// 估算值 = prevCount × (窗口剩余比例) + curCount
// 这是 Redis + Nginx 常用的近似方案，内存占用固定，误差 < 1/limit。
type SlidingWindowCounter struct {
	limit     int64
	window    int64 // 窗口大小（纳秒）
	curCount  int64
	prevCount int64
	curWindow int64 // 当前窗口起始时间（纳秒）
	mu        sync.Mutex
}

func NewSlidingWindowCounter(limit int64, window time.Duration) *SlidingWindowCounter {
	return &SlidingWindowCounter{
		limit:     limit,
		window:    window.Nanoseconds(),
		curWindow: time.Now().UnixNano(),
	}
}

func (l *SlidingWindowCounter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now().UnixNano()
	elapsed := now - l.curWindow

	if elapsed >= l.window*2 {
		// 两个窗口均已过期，直接归零
		l.prevCount = 0
		l.curCount = 0
		l.curWindow = now
		elapsed = 0
	} else if elapsed >= l.window {
		// 当前窗口到期，向前滑动一格
		l.prevCount = l.curCount
		l.curCount = 0
		l.curWindow += l.window
		elapsed = now - l.curWindow
	}

	// 上个窗口在当前滑动窗口内的加权贡献
	ratio := float64(l.window-elapsed) / float64(l.window)
	estimate := float64(l.prevCount)*ratio + float64(l.curCount)

	if int64(estimate) >= l.limit {
		return false
	}

	l.curCount++
	return true
}

func main_sliding() {
	const limit = 5
	window := time.Second

	fmt.Println("=== 精确滑动窗口（Log）===")
	log := NewSlidingWindowLog(limit, window)
	for i := 1; i <= 8; i++ {
		if log.Allow() {
			fmt.Printf("请求 %d: ✅ 通过\n", i)
		} else {
			fmt.Printf("请求 %d: ❌ 限流\n", i)
		}
	}
	fmt.Println("--- 等待 600ms ---")
	time.Sleep(600 * time.Millisecond)
	for i := 9; i <= 12; i++ {
		if log.Allow() {
			fmt.Printf("请求 %d: ✅ 通过\n", i)
		} else {
			fmt.Printf("请求 %d: ❌ 限流\n", i)
		}
	}

	fmt.Println("\n=== 近似滑动窗口（Counter）===")
	counter := NewSlidingWindowCounter(limit, window)
	for i := 1; i <= 8; i++ {
		if counter.Allow() {
			fmt.Printf("请求 %d: ✅ 通过\n", i)
		} else {
			fmt.Printf("请求 %d: ❌ 限流\n", i)
		}
	}
	fmt.Println("--- 等待 600ms ---")
	time.Sleep(600 * time.Millisecond)
	for i := 9; i <= 12; i++ {
		if counter.Allow() {
			fmt.Printf("请求 %d: ✅ 通过\n", i)
		} else {
			fmt.Printf("请求 %d: ❌ 限流\n", i)
		}
	}
}
