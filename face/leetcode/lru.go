package leetcode

import (
	"container/list"
	"sync"
	"time"
)

// LRUCache 带过期时间的LRU缓存
type LRUCache struct {
	capacity int
	mu       sync.RWMutex
	cache    map[interface{}]*list.Element // 快速查找
	lruList  *list.List                    // LRU双向链表
	stopCh   chan struct{}                 // 停止清理协程
}

// cacheItem 缓存项
type cacheItem struct {
	key        interface{}
	value      interface{}
	expireTime time.Time // 过期时间，零值表示永不过期
}

// NewLRUCache 创建新缓存
func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		capacity = 100 // 默认容量
	}
	c := &LRUCache{
		capacity: capacity,
		cache:    make(map[interface{}]*list.Element),
		lruList:  list.New(),
		stopCh:   make(chan struct{}),
	}
	// 启动过期清理协程
	go c.cleanupExpired()
	return c
}

// Set 设置缓存，支持过期时间
func (c *LRUCache) Set(key, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var expireTime time.Time
	if ttl > 0 {
		expireTime = time.Now().Add(ttl)
	}
	// 如果key已存在，更新值
	if elem, ok := c.cache[key]; ok {
		item := elem.Value.(*cacheItem)
		item.value = value
		item.expireTime = expireTime
		c.lruList.MoveToFront(elem) // 移到最前
		return
	}
	// 检查容量，如果满了淘汰最久未使用的
	if c.lruList.Len() >= c.capacity {
		c.evictLocked()
	}
	// 插入新节点
	item := &cacheItem{
		key:        key,
		value:      value,
		expireTime: expireTime,
	}
	elem := c.lruList.PushFront(item)
	c.cache[key] = elem
}

// Get 获取缓存
func (c *LRUCache) Get(key interface{}) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	elem, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	item := elem.Value.(*cacheItem)
	// 检查是否过期
	if !item.expireTime.IsZero() && time.Now().After(item.expireTime) {
		c.removeElementLocked(elem)
		return nil, false
	}
	// 移到最前（最近使用）
	c.lruList.MoveToFront(elem)
	return item.value, true
}

// Del 删除缓存
func (c *LRUCache) Del(key interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.cache[key]; ok {
		c.removeElementLocked(elem)
	}
}

// Len 返回当前缓存数量
func (c *LRUCache) Len() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lruList.Len()
}

// Clear 清空缓存
func (c *LRUCache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache = make(map[interface{}]*list.Element)
	c.lruList.Init()
}

// Close 关闭缓存，停止清理协程
func (c *LRUCache) Close() {
	close(c.stopCh)
}

// evictLocked 淘汰最久未使用的（调用前必须加锁）
func (c *LRUCache) evictLocked() {
	elem := c.lruList.Back()
	if elem != nil {
		c.removeElementLocked(elem)
	}
}

// removeElementLocked 删除节点（调用前必须加锁）
func (c *LRUCache) removeElementLocked(elem *list.Element) {
	item := elem.Value.(*cacheItem)
	delete(c.cache, item.key)
	c.lruList.Remove(elem)
}

// cleanupExpired 定时清理过期缓存
func (c *LRUCache) cleanupExpired() {
	ticker := time.NewTicker(30 * time.Second) // 每30秒清理一次
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			c.cleanExpiredOnce()
		case <-c.stopCh:
			return
		}
	}
}

// cleanExpiredOnce 单次清理过期项
func (c *LRUCache) cleanExpiredOnce() {
	c.mu.Lock()
	defer c.mu.Unlock()

	now := time.Now()
	for elem := c.lruList.Back(); elem != nil; {
		next := elem.Prev() // 先保存前一个，因为可能删除当前
		item := elem.Value.(*cacheItem)
		if !item.expireTime.IsZero() && now.After(item.expireTime) {
			c.removeElementLocked(elem)
		}
		elem = next
	}
}
