package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestOrderedMap(t *testing.T) {
	t.Run("OrderedMap", func(t *testing.T) {
		OrderedMapMain()
	})
}

func OrderedMapMain() {
	om := NewOrderedMap()

	om.Set("orange", 5)
	om.Set("apple", 3)
	om.Set("banana", 2)

	fmt.Println("按插入顺序:")
	om.Iterate(func(key string, value interface{}) {
		fmt.Printf("%s: %v\n", key, value)
	})

	fmt.Println("\n按键排序后:")
	om.SortKeys()
	om.Iterate(func(key string, value interface{}) {
		fmt.Printf("%s: %v\n", key, value)
	})

	fmt.Println("\n按值排序后:")
	om.SortByValue()
	om.Iterate(func(key string, value interface{}) {
		fmt.Printf("%s: %v\n", key, value)
	})
}

// OrderedMap 自定义有序map
type OrderedMap struct {
	keys   []string
	values map[string]interface{}
}

// NewOrderedMap 创建新的有序map
func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		keys:   make([]string, 0),
		values: make(map[string]interface{}),
	}
}

// Set 设置键值对
func (om *OrderedMap) Set(key string, value interface{}) {
	if _, exists := om.values[key]; !exists {
		om.keys = append(om.keys, key)
	}
	om.values[key] = value
}

// Get 获取值
func (om *OrderedMap) Get(key string) (interface{}, bool) {
	val, exists := om.values[key]
	return val, exists
}

// Keys 返回所有键（按插入顺序）
func (om *OrderedMap) Keys() []string {
	return om.keys
}

// SortKeys 按键排序
func (om *OrderedMap) SortKeys() {
	sort.Strings(om.keys)
}

// SortByValue 按值排序（需要知道值的类型）
func (om *OrderedMap) SortByValue() {
	sort.Slice(om.keys, func(i, j int) bool {
		// 这里假设值是int类型，可以根据实际情况调整
		return om.values[om.keys[i]].(int) < om.values[om.keys[j]].(int)
	})
}

// Iterate 按顺序迭代
func (om *OrderedMap) Iterate(f func(key string, value interface{})) {
	for _, key := range om.keys {
		f(key, om.values[key])
	}
}
