package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"sync"
	"time"

	"github.com/gomodule/redigo/redis"
)

const host = "172.26.4.121"

// OrderService 订单服务
type OrderService struct {
	redisPool *redis.Pool
}

// NewOrderService 创建订单服务实例
func NewOrderService(redisAddr string) *OrderService {
	pool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
	return &OrderService{redisPool: pool}
}

// Close 关闭Redis连接池
func (s *OrderService) Close() {
	s.redisPool.Close()
}

// generateLockValue 生成唯一的锁值
func generateLockValue() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// tryAcquireLock 尝试获取分布式锁
func (s *OrderService) tryAcquireLock(ctx context.Context, lockKey string, expireTime time.Duration) (string, bool, error) {
	conn := s.redisPool.Get()
	defer conn.Close()

	lockValue, err := generateLockValue()
	if err != nil {
		return "", false, fmt.Errorf("generate lock value failed: %v", err)
	}

	// 使用 SET key value NX EX 命令原子性地设置锁
	reply, err := redis.String(conn.Do("SET", lockKey, lockValue, "NX", "EX", int(expireTime/time.Second)))
	if err != nil {
		if err == redis.ErrNil {
			// 锁已存在，获取失败
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
func (s *OrderService) releaseLock(ctx context.Context, lockKey string, lockValue string) error {
	conn := s.redisPool.Get()
	defer conn.Close()

	// 使用 Lua 脚本保证原子性：只有锁的值匹配时才删除
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

// renewLock 续期锁
func (s *OrderService) renewLock(ctx context.Context, lockKey string, lockValue string, expireTime time.Duration) (bool, error) {
	conn := s.redisPool.Get()
	defer conn.Close()

	// 使用 Lua 脚本保证原子性续期
	luaScript := `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("expire", KEYS[1], ARGV[2])
	else
		return 0
	end
	`

	script := redis.NewScript(1, luaScript)
	result, err := redis.Int(script.Do(conn, lockKey, lockValue, int(expireTime/time.Second)))
	if err != nil {
		return false, fmt.Errorf("failed to renew lock: %v", err)
	}

	return result == 1, nil
}

// acquireLockWithWatchdog 获取带看门狗的锁
func (s *OrderService) acquireLockWithWatchdog(ctx context.Context, lockKey string, expireTime time.Duration) (string, context.CancelFunc, error) {
	lockValue, acquired, err := s.tryAcquireLock(ctx, lockKey, expireTime)
	if err != nil || !acquired {
		return "", nil, err
	}

	watchdogCtx, cancel := context.WithCancel(context.Background())

	// 启动看门狗协程自动续期
	go s.watchdog(watchdogCtx, lockKey, lockValue, expireTime)

	return lockValue, cancel, nil
}

// watchdog 看门狗协程，定期续期
func (s *OrderService) watchdog(ctx context.Context, lockKey, lockValue string, expireTime time.Duration) {
	ticker := time.NewTicker(expireTime / 2) // 在过期时间的一半时续期
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return // 主任务完成，退出看门狗
		case <-ticker.C:
			renewed, err := s.renewLock(ctx, lockKey, lockValue, expireTime)
			if err != nil || !renewed {
				log.Printf("Watchdog renew lock failed: %v, renewed: %v", err, renewed)
				return
			}
			log.Printf("Lock %s renewed successfully", lockKey)
		}
	}
}

// CreateOrder 创建订单（使用分布式锁保护）
func (s *OrderService) CreateOrder(ctx context.Context, userID, productID int64, quantity int) error {
	// 1. 生成锁的Key
	lockKey := fmt.Sprintf("order_lock:user_%d:product_%d", userID, productID)
	const lockExpire = 10 * time.Second

	// 2. 尝试获取锁（带重试机制）
	var lockValue string
	var acquired bool
	var err error

	for i := 0; i < 3; i++ { // 最大重试3次
		lockValue, acquired, err = s.tryAcquireLock(ctx, lockKey, lockExpire)
		if err != nil {
			log.Printf("Acquire lock attempt %d failed: %v", i+1, err)
			continue
		}
		if acquired {
			break
		}

		// 等待随机时间后重试，避免活锁
		waitTime := time.Duration(100+i*100) * time.Millisecond
		log.Printf("Lock not available, retrying in %v...", waitTime)
		time.Sleep(waitTime)
	}

	if !acquired {
		return fmt.Errorf("create order failed: cannot acquire lock after retries")
	}

	// 3. 确保最终释放锁
	defer func() {
		if releaseErr := s.releaseLock(ctx, lockKey, lockValue); releaseErr != nil {
			log.Printf("WARNING: failed to release lock %s: %v", lockKey, releaseErr)
		} else {
			log.Printf("Lock %s released successfully", lockKey)
		}
	}()

	log.Printf("Lock acquired successfully for order creation: user %d, product %d", userID, productID)

	// 4. 执行核心业务逻辑
	if err := s.processOrderCreation(ctx, userID, productID, quantity); err != nil {
		return fmt.Errorf("order processing failed: %v", err)
	}

	log.Printf("Order created successfully for user %d, product %d", userID, productID)
	return nil
}

// processOrderCreation 处理订单创建的核心业务逻辑
func (s *OrderService) processOrderCreation(ctx context.Context, userID, productID int64, quantity int) error {
	// 这里应该是你的实际订单创建逻辑，例如：
	// - 检查库存
	// - 创建订单记录
	// - 扣减库存
	// - 其他业务操作...

	// 模拟业务处理时间
	processingTime := time.Duration(100+randInt(200)) * time.Millisecond
	time.Sleep(processingTime)

	// 模拟随机失败
	if randInt(10) == 0 {
		return fmt.Errorf("random order processing failure")
	}

	return nil
}

// randInt 生成随机整数
func randInt(max int) int {
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return int(n.Int64())
}

func redisLockMain() {
	// 初始化订单服务
	orderService := NewOrderService(fmt.Sprintf("%s:6379", host))
	defer orderService.Close()

	// 测试Redis连接
	conn := orderService.redisPool.Get()
	defer conn.Close()

	if _, err := conn.Do("PING"); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully")

	// 模拟并发创建订单
	var wg sync.WaitGroup
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(orderIndex int) {
			defer wg.Done()

			userID := int64(orderIndex%3 + 1)       // 3个用户
			productID := int64(1000 + orderIndex%2) // 2个商品

			log.Printf("Starting order creation %d: user %d, product %d",
				orderIndex, userID, productID)

			err := orderService.CreateOrder(ctx, userID, productID, 1)
			if err != nil {
				log.Printf("Order %d failed: %v", orderIndex, err)
			} else {
				log.Printf("Order %d completed successfully", orderIndex)
			}
		}(i)

		// 稍微错开启动时间
		time.Sleep(10 * time.Millisecond)
	}

	wg.Wait()
	log.Println("All order creation attempts completed")
}
