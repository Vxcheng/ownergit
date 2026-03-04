package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// RedisTokenBucketLimiter 基于 Redis + Lua 的分布式令牌桶限流器。
//
// 核心思路：将"当前令牌数"和"上次填充时间"存储在 Redis Hash 中，
// 通过 Lua 脚本保证 读取 → 计算 → 写入 三步的原子性，
// 从而在多实例部署场景下实现全局一致的限流，无需本地状态。
type RedisTokenBucketLimiter struct {
	pool     *redis.Pool
	rate     float64      // 每秒产生的令牌数
	capacity float64      // 桶的最大容量
	script   *redis.Script
}

// tokenBucketLua 令牌桶 Lua 脚本
//
// KEYS[1]  限流 key，例如 "ratelimit:user:123"
// ARGV[1]  rate     每秒令牌数（float）
// ARGV[2]  capacity 桶容量（float）
// ARGV[3]  now      当前毫秒时间戳
// ARGV[4]  n        本次消耗令牌数（通常为 1）
//
// 返回 1 = 允许，0 = 拒绝
const tokenBucketLua = `
local key      = KEYS[1]
local rate     = tonumber(ARGV[1])
local capacity = tonumber(ARGV[2])
local now      = tonumber(ARGV[3])
local n        = tonumber(ARGV[4])

-- 读取上次保存的状态
local info        = redis.call("HMGET", key, "tokens", "last_refill")
local tokens      = tonumber(info[1])
local last_refill = tonumber(info[2])

-- 首次请求：初始化为满桶
if tokens == nil then
    tokens      = capacity
    last_refill = now
end

-- 按经过时间（毫秒 → 秒）补充令牌，不超过桶容量
local elapsed  = math.max(0, (now - last_refill) / 1000.0)
local refilled = math.min(capacity, tokens + elapsed * rate)

-- 判断令牌是否充足
local allowed = 0
if refilled >= n then
    refilled = refilled - n
    allowed  = 1
end

-- 持久化状态；TTL = 桶从空填满所需秒数 + 1s 缓冲，避免僵尸 key
local ttl = math.ceil(capacity / rate) + 1
redis.call("HMSET", key, "tokens", refilled, "last_refill", now)
redis.call("EXPIRE", key, ttl)

return allowed
`

func NewRedisTokenBucketLimiter(pool *redis.Pool, rate, capacity float64) *RedisTokenBucketLimiter {
	return &RedisTokenBucketLimiter{
		pool:     pool,
		rate:     rate,
		capacity: capacity,
		script:   redis.NewScript(1, tokenBucketLua),
	}
}

// Allow 消耗 1 个令牌，判断 key 对应的请求是否允许通过。
func (l *RedisTokenBucketLimiter) Allow(key string) (bool, error) {
	return l.AllowN(key, 1)
}

// AllowN 消耗 n 个令牌（适用于带权重的请求）。
func (l *RedisTokenBucketLimiter) AllowN(key string, n int) (bool, error) {
	conn := l.pool.Get()
	defer conn.Close()

	result, err := redis.Int(l.script.Do(conn, key,
		l.rate,
		l.capacity,
		time.Now().UnixMilli(),
		n,
	))
	if err != nil {
		return false, fmt.Errorf("redis lua 执行失败: %w", err)
	}
	return result == 1, nil
}

func newRedisPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		MaxActive:   100,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

func main_redis_limit() {
	pool := newRedisPool("172.26.4.121:6379")
	defer pool.Close()

	// 验证连接
	conn := pool.Get()
	if _, err := conn.Do("PING"); err != nil {
		fmt.Printf("Redis 连接失败: %v\n", err)
		conn.Close()
		return
	}
	conn.Close()

	// 每秒 5 个令牌，桶容量 5（初始满桶）
	limiter := NewRedisTokenBucketLimiter(pool, 5, 5)
	key := "ratelimit:demo:user_1"

	fmt.Println("=== Redis 分布式令牌桶限流（rate=5/s, capacity=5）===")
	fmt.Println("--- 连续发送 8 个请求（前 5 个通过，后 3 个被限流）---")
	for i := 1; i <= 8; i++ {
		ok, err := limiter.Allow(key)
		if err != nil {
			fmt.Printf("请求 %2d: ⚠️  %v\n", i, err)
			continue
		}
		if ok {
			fmt.Printf("请求 %2d: ✅ 通过\n", i)
		} else {
			fmt.Printf("请求 %2d: ❌ 限流\n", i)
		}
	}

	fmt.Println("\n--- 等待 1s，令牌补充至 5 ---")
	time.Sleep(time.Second)

	fmt.Println("--- 再发 6 个请求（前 5 个通过，第 6 个被限流）---")
	for i := 9; i <= 14; i++ {
		ok, err := limiter.Allow(key)
		if err != nil {
			fmt.Printf("请求 %2d: ⚠️  %v\n", i, err)
			continue
		}
		if ok {
			fmt.Printf("请求 %2d: ✅ 通过\n", i)
		} else {
			fmt.Printf("请求 %2d: ❌ 限流\n", i)
		}
	}

	// 演示 AllowN：消耗 3 个令牌的重量级请求
	fmt.Println("\n--- AllowN(3)：等待 600ms 后发 2 次批量请求 ---")
	time.Sleep(600 * time.Millisecond)
	for i := 1; i <= 2; i++ {
		ok, err := limiter.AllowN(key, 3)
		if err != nil {
			fmt.Printf("批量请求 %d: ⚠️  %v\n", i, err)
			continue
		}
		if ok {
			fmt.Printf("批量请求 %d (cost=3): ✅ 通过\n", i)
		} else {
			fmt.Printf("批量请求 %d (cost=3): ❌ 限流\n", i)
		}
	}
}
