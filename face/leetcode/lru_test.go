package leetcode

import (
	"testing"
	"time"
)

func TestLRUCache_BasicSetGet(t *testing.T) {
	c := NewLRUCache(3)
	defer c.Close()

	c.Set("a", 1, 0)
	c.Set("b", 2, 0)

	if v, ok := c.Get("a"); !ok || v != 1 {
		t.Errorf("expected 1, got %v", v)
	}
	if _, ok := c.Get("x"); ok {
		t.Error("expected miss for key x")
	}
}

func TestLRUCache_Eviction(t *testing.T) {
	c := NewLRUCache(2)
	defer c.Close()

	c.Set("a", 1, 0)
	c.Set("b", 2, 0)
	c.Get("a")      // a 最近使用，b 变成最久未使用
	c.Set("c", 3, 0) // 触发淘汰，b 应被淘汰

	if _, ok := c.Get("b"); ok {
		t.Error("b should have been evicted")
	}
	if _, ok := c.Get("a"); !ok {
		t.Error("a should still exist")
	}
}

func TestLRUCache_TTLExpire(t *testing.T) {
	c := NewLRUCache(10)
	defer c.Close()

	c.Set("k", "v", 50*time.Millisecond)
	if _, ok := c.Get("k"); !ok {
		t.Fatal("key should exist before expiry")
	}

	time.Sleep(60 * time.Millisecond)
	if _, ok := c.Get("k"); ok {
		t.Error("key should have expired")
	}
}

func TestLRUCache_DelAndClear(t *testing.T) {
	c := NewLRUCache(5)
	defer c.Close()

	c.Set("a", 1, 0)
	c.Set("b", 2, 0)
	c.Del("a")
	if _, ok := c.Get("a"); ok {
		t.Error("a should be deleted")
	}

	c.Clear()
	if c.Len() != 0 {
		t.Errorf("expected len 0 after clear, got %d", c.Len())
	}
}
