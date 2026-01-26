package main

import (
	"testing"
)

func TestRedis(t *testing.T) {
	cfg := ConfigST{
		User: "root",
		Pw:   "",
		Host: "localhost:6379",
		Db:   0,
	}
	InitRedis(cfg)
	t.Run("expire key []", func(t *testing.T) {
		WriteMapDataWithExpire(map[string]interface{}{"aaa": "aaa"}, int64(100))
	})
}
