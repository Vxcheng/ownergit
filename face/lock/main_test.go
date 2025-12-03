package main

import "testing"



func TestRedisLock(t *testing.T) {
	t.Run("redisLockBenchMain", func(t *testing.T) {
		redisLockBenchMain()
	})

	t.Run("redisLockMain", func(t *testing.T) {
		redisLockMain()
	})
}
