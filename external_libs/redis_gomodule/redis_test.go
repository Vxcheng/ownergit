package main

import (
	"testing"
)

func TestRedis(t *testing.T) {
	cfg := ConfigST{
		User: "root",
		Pw:   "redis123!",
		Host: "192.168.10.64:6379",
		Db:   3,
	}
	InitRedis(cfg)
	t.Run("expire key []", func(t *testing.T) {
		WriteMapDataWithExpire(map[string]interface{}{"aaa": "aaa"}, int64(10))
	})
}
