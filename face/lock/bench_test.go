package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gomodule/redigo/redis"
)

// DistributedLockTester 分布式锁压测器
type DistributedLockTester struct {
	redisPool         *redis.Pool
	totalRequests     int32
	successCount      int32
	failureCount      int32
	timeoutCount      int32
	lockConflictCount int32
	totalLatency      int64
}

// NewDistributedLockTester 创建压测器
func NewDistributedLockTester(redisAddr string) *DistributedLockTester {
	pool := &redis.Pool{
		MaxIdle:     50,
		MaxActive:   200,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
	return &DistributedLockTester{redisPool: pool}
}

// tryAcquireLock 尝试获取分布式锁
func (t *DistributedLockTester) tryAcquireLock(ctx context.Context, lockKey string, expireTime time.Duration) (string, bool, error) {
	conn := t.redisPool.Get()
	defer conn.Close()

	lockValue, err := generateLockValue()
	if err != nil {
		return "", false, fmt.Errorf("generate lock value failed: %v", err)
	}

	reply, err := redis.String(conn.Do("SET", lockKey, lockValue, "NX", "EX", int(expireTime/time.Second)))
	if err != nil {
		if err == redis.ErrNil {
			return "", false, nil
		}
		return "", false, fmt.Errorf("failed to acquire lock: %v", err)
	}

	if reply != "OK" {
		return "", false, nil
	}

	return lockValue, true, nil
}

// releaseLock 释放分布式锁
func (t *DistributedLockTester) releaseLock(ctx context.Context, lockKey string, lockValue string) error {
	conn := t.redisPool.Get()
	defer conn.Close()

	luaScript := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
	`

	script := redis.NewScript(1, luaScript)
	_, err := redis.Int(script.Do(conn, lockKey, lockValue))
	if err != nil {
		return fmt.Errorf("failed to release lock: %v", err)
	}
	return nil
}

// TestLockContention 测试锁竞争场景
func (t *DistributedLockTester) TestLockContention(ctx context.Context, concurrentWorkers, requestsPerWorker int, lockKeyPrefix string) {
	var wg sync.WaitGroup
	startTime := time.Now()

	log.Printf("Starting lock contention test: %d workers, %d requests each", concurrentWorkers, requestsPerWorker)

	for i := 0; i < concurrentWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			for j := 0; j < requestsPerWorker; j++ {
				atomic.AddInt32(&t.totalRequests, 1)
				requestID := workerID*requestsPerWorker + j

				t.testSingleLock(ctx, workerID, requestID, lockKeyPrefix)
			}
		}(i)
	}

	wg.Wait()

	duration := time.Since(startTime)
	t.printTestResults(duration)
}

// testSingleLock 测试单个锁操作
func (t *DistributedLockTester) testSingleLock(ctx context.Context, workerID, requestID int, lockKeyPrefix string) {
	start := time.Now()

	// 模拟不同的锁键模式
	lockKey := fmt.Sprintf("%s:resource_%d", lockKeyPrefix, requestID%10) // 10个不同的资源

	const lockExpire = 5 * time.Second
	timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var acquired bool
	var lockValue string
	var err error

	// 重试机制
	for retry := 0; retry < 3; retry++ {
		lockValue, acquired, err = t.tryAcquireLock(timeoutCtx, lockKey, lockExpire)
		if err != nil {
			log.Printf("Worker %d, Request %d: Lock acquire error: %v", workerID, requestID, err)
			time.Sleep(time.Duration(retry+1) * 100 * time.Millisecond)
			continue
		}
		if acquired {
			break
		}

		atomic.AddInt32(&t.lockConflictCount, 1)
		time.Sleep(time.Duration(retry+1) * 50 * time.Millisecond)
	}

	if !acquired {
		atomic.AddInt32(&t.failureCount, 1)
		log.Printf("Worker %d, Request %d: Failed to acquire lock after retries", workerID, requestID)
		return
	}

	// 模拟业务处理时间
	processingTime := time.Duration(randInt(100)+50) * time.Millisecond // 50-150ms
	time.Sleep(processingTime)

	// 释放锁
	if err := t.releaseLock(ctx, lockKey, lockValue); err != nil {
		log.Printf("Worker %d, Request %d: Failed to release lock: %v", workerID, requestID, err)
	}

	latency := time.Since(start).Milliseconds()
	atomic.AddInt64(&t.totalLatency, latency)
	atomic.AddInt32(&t.successCount, 1)

	if requestID%100 == 0 {
		log.Printf("Worker %d, Request %d: Lock acquired and released successfully, latency: %dms",
			workerID, requestID, latency)
	}
}

// TestSameKeyContention 测试同一把锁的高竞争场景
func (t *DistributedLockTester) TestSameKeyContention(ctx context.Context, concurrentWorkers int) {
	var wg sync.WaitGroup
	sameLockKey := "high_contention_lock:same_key"
	startTime := time.Now()
	t.totalRequests = 1

	log.Printf("Starting same key contention test: %d workers competing for same lock", concurrentWorkers)

	for i := 0; i < concurrentWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()

			start := time.Now()
			lockValue, acquired, err := t.tryAcquireLock(ctx, sameLockKey, 2*time.Second)

			if err != nil {
				atomic.AddInt32(&t.failureCount, 1)
				log.Printf("SameKey Worker %d: Error: %v", workerID, err)
				return
			}

			if !acquired {
				atomic.AddInt32(&t.lockConflictCount, 1)
				log.Printf("SameKey Worker %d: Lock not available", workerID)
				return
			}

			defer t.releaseLock(ctx, sameLockKey, lockValue)

			// 持有锁一段时间
			time.Sleep(100 * time.Millisecond)

			latency := time.Since(start).Milliseconds()
			atomic.AddInt64(&t.totalLatency, latency)
			atomic.AddInt32(&t.successCount, 1)

			log.Printf("SameKey Worker %d: Acquired lock successfully, latency: %dms", workerID, latency)
		}(i)

		// 错开启动时间
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
	duration := time.Since(startTime)
	t.printTestResults(duration)
}

// printTestResults 打印测试结果
func (t *DistributedLockTester) printTestResults(duration time.Duration) {
	total := atomic.LoadInt32(&t.totalRequests)
	success := atomic.LoadInt32(&t.successCount)
	failure := atomic.LoadInt32(&t.failureCount)
	conflicts := atomic.LoadInt32(&t.lockConflictCount)
	totalLatency := atomic.LoadInt64(&t.totalLatency)

	avgLatency := float64(0)
	if success > 0 {
		avgLatency = float64(totalLatency) / float64(success)
	}

	qps := float64(total) / duration.Seconds()

	fmt.Printf("\n=== TEST RESULTS ===\n")
	fmt.Printf("Duration: %v\n", duration)
	fmt.Printf("Total Requests: %d\n", total)
	fmt.Printf("Success: %d (%.1f%%)\n", success, float64(success)/float64(total)*100)
	fmt.Printf("Failure: %d (%.1f%%)\n", failure, float64(failure)/float64(total)*100)
	fmt.Printf("Lock Conflicts: %d\n", conflicts)
	fmt.Printf("QPS: %.1f requests/second\n", qps)
	fmt.Printf("Average Latency: %.2f ms\n", avgLatency)
	fmt.Printf("===================\n")
}

// ResetStats 重置统计信息
func (t *DistributedLockTester) ResetStats() {
	atomic.StoreInt32(&t.totalRequests, 0)
	atomic.StoreInt32(&t.successCount, 0)
	atomic.StoreInt32(&t.failureCount, 0)
	atomic.StoreInt32(&t.timeoutCount, 0)
	atomic.StoreInt32(&t.lockConflictCount, 0)
	atomic.StoreInt64(&t.totalLatency, 0)
}

func redisLockBenchMain() {
	// 初始化压测器
	tester := NewDistributedLockTester(fmt.Sprintf("%s:6379", host))
	defer tester.redisPool.Close()

	// 测试Redis连接
	conn := tester.redisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	ctx := context.Background()

	// 测试场景1: 一般锁竞争
	fmt.Println("\n1. Testing general lock contention...")
	tester.ResetStats()
	tester.TestLockContention(ctx, 20, 50, "test_lock") // 20个worker，每个50次请求

	// 测试场景2: 高竞争同一把锁
	fmt.Println("\n2. Testing high contention for same lock...")
	tester.ResetStats()
	tester.TestSameKeyContention(ctx, 30) // 30个worker竞争同一把锁

	// 测试场景3: 大量不同锁
	fmt.Println("\n3. Testing many different locks...")
	tester.ResetStats()
	tester.TestLockContention(ctx, 10, 100, "test_lock") // 10个worker，每个100次请求

	// 测试场景4: 极端高并发
	fmt.Println("\n4. Testing extreme concurrency...")
	tester.ResetStats()
	tester.TestLockContention(ctx, 50, 20, "stress_test") // 50个worker，每个20次请求
}
