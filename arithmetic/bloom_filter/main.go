package main

import (
	"fmt"
	"hash"
	"hash/fnv"
	"math"
)

// BloomFilter 布隆过滤器结构体
type BloomFilter struct {
	bitSet      []bool        // 位数组（实际使用bool切片模拟）
	bitSize     uint          // 位数组大小（位数）
	hashFuncNum uint          // 哈希函数数量
	hashFuncs   []hash.Hash64 // 哈希函数实例
}

// NewBloomFilter 创建布隆过滤器
// expectedInsertions: 预期插入元素数量
// falsePositiveRate: 期望误判率 (0 < rate < 1)
func NewBloomFilter(expectedInsertions uint, falsePositiveRate float64) *BloomFilter {
	if falsePositiveRate <= 0 || falsePositiveRate >= 1 {
		panic("误判率必须在0和1之间")
	}

	// 计算最优位数组大小
	bitSize := uint(-float64(expectedInsertions) * math.Log(falsePositiveRate) / (math.Log(2) * math.Log(2)))
	// 计算最优哈希函数数量
	hashFuncNum := uint(math.Log(2) * float64(bitSize) / float64(expectedInsertions))
	if hashFuncNum < 1 {
		hashFuncNum = 1
	}

	// 初始化哈希函数（使用FNV-1a算法，性能较好）
	hashFuncs := make([]hash.Hash64, hashFuncNum)
	for i := uint(0); i < hashFuncNum; i++ {
		hashFuncs[i] = fnv.New64a()
	}

	return &BloomFilter{
		bitSet:      make([]bool, bitSize),
		bitSize:     bitSize,
		hashFuncNum: hashFuncNum,
		hashFuncs:   hashFuncs,
	}
}

// Add 添加元素到布隆过滤器
func (bf *BloomFilter) Add(data []byte) {
	for i := uint(0); i < bf.hashFuncNum; i++ {
		// 重置哈希函数以复用
		bf.hashFuncs[i].Reset()
		bf.hashFuncs[i].Write(data)
		hashValue := bf.hashFuncs[i].Sum64()
		position := uint(hashValue % uint64(bf.bitSize))
		bf.bitSet[position] = true
	}
}

// AddString 添加字符串元素（便捷方法）
func (bf *BloomFilter) AddString(s string) {
	bf.Add([]byte(s))
}

// MightContain 检查元素是否可能存在
// 返回true: 可能存在（可能有误判）
// 返回false: 一定不存在
func (bf *BloomFilter) MightContain(data []byte) bool {
	for i := uint(0); i < bf.hashFuncNum; i++ {
		bf.hashFuncs[i].Reset()
		bf.hashFuncs[i].Write(data)
		hashValue := bf.hashFuncs[i].Sum64()
		position := uint(hashValue % uint64(bf.bitSize))
		if !bf.bitSet[position] {
			return false // 只要有一位为0，则一定不存在
		}
	}
	return true // 所有位都为1，则可能存在
}

// MightContainString 检查字符串元素是否可能存在（便捷方法）
func (bf *BloomFilter) MightContainString(s string) bool {
	return bf.MightContain([]byte(s))
}

// UtilizationRate 获取位数组使用率（监控用）
func (bf *BloomFilter) UtilizationRate() float64 {
	count := 0
	for _, bit := range bf.bitSet {
		if bit {
			count++
		}
	}
	return float64(count) / float64(len(bf.bitSet))
}

func main() {
	// 创建布隆过滤器：预期插入10000个元素，误判率1%
	bf := NewBloomFilter(10000*10000, 0.001)
	fmt.Printf("初始化: 位大小=%d, 哈希函数数=%d\n", bf.bitSize, bf.hashFuncNum)

	// 添加元素
	bf.AddString("https://www.example.com")
	bf.AddString("user12345")
	bf.AddString("data_item_678")

	// 测试存在性
	fmt.Println("测试 'user12345':", bf.MightContainString("user12345")) // 应返回 true
	fmt.Println("测试 'unknown':", bf.MightContainString("unknown"))     // 很可能返回 false

	// 误判率测试（演示用）
	totalTests := 10000
	falsePositives := 0
	for i := 0; i < totalTests; i++ {
		testItem := fmt.Sprintf("未添加的项目_%d", i)
		if bf.MightContainString(testItem) {
			falsePositives++
		}
	}
	actualFalsePositiveRate := float64(falsePositives) / float64(totalTests)
	fmt.Printf("实际误判率: %.4f%%\n", actualFalsePositiveRate*100)
	fmt.Printf("位数组使用率: %.2f%%\n", bf.UtilizationRate()*100)
}
