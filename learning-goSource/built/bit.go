package built

// golang 所有位操作，示例、核心函数。
/*
操作	表达式	说明
设置第 n 位	x | (1 << n)	将第 n 位设为 1
清除第 n 位	x &^ (1 << n)	将第 n 位设为 0
翻转第 n 位	x ^ (1 << n)	将第 n 位取反
检查第 n 位	(x >> n) & 1	返回 0 或 1
最低位 1	x & -x	提取最低位的 1
清除最低位 1	x & (x - 1)	将最低位的 1 清零
判断奇偶	x & 1	1 为奇数，0 为偶数
乘 2^n	x << n	左移 n 位
除 2^n	x >> n	右移 n 位
取模 2^n	x & (2^n - 1)	仅对 2 的幂有效
*/

// And 按位与
func And(a, b int) int { return a & b }

// Or 按位或
func Or(a, b int) int { return a | b }

// Xor 按位异或
func Xor(a, b int) int { return a ^ b }

// Not8 按位取反（uint8 范围）
func Not8(a uint8) uint8 { return ^a }

// LeftShift 左移 n 位
func LeftShift(a, n int) int { return a << n }

// RightShift 右移 n 位
func RightShift(a, n int) int { return a >> n }

// SetBit 将第 n 位（从0开始）设为1
func SetBit(a, n int) int { return a | (1 << n) }

// ClearBit 将第 n 位清为0
func ClearBit(a, n int) int { return a &^ (1 << n) }

// ToggleBit 翻转第 n 位
func ToggleBit(a, n int) int { return a ^ (1 << n) }

// TestBit 检测第 n 位是否为1
func TestBit(a, n int) bool { return a&(1<<n) != 0 }

// CountOnes 统计二进制中1的个数（Brian Kernighan 算法）
func CountOnes(a int) int {
	count := 0
	for a != 0 {
		a &= a - 1
		count++
	}
	return count
}

// IsPowerOfTwo 判断是否为2的幂
func IsPowerOfTwo(n int) bool { return n > 0 && n&(n-1) == 0 }

// LowBit 取最低位的1（树状数组核心操作）
func LowBit(a int) int { return a & (-a) }

// ReverseBits8 反转8位整数的所有位
func ReverseBits8(a uint8) uint8 {
	var result uint8
	for i := 0; i < 8; i++ {
		result = (result << 1) | (a & 1)
		a >>= 1
	}
	return result
}

// SwapWithoutTemp 不用临时变量交换两个整数（异或实现）
func SwapWithoutTemp(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}
