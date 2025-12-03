package arithmetic

import "fmt"

// 判断奇偶性（比取模快）
func isEven(n int) bool {
	return n&1 == 0
}

// 交换两个数（不需要临时变量）
func swap(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}

// 绝对值（快速版本）
func abs(n int) int {
	mask := n >> 63
	return (n + mask) ^ mask
}

func BitDemo() {
	a := 0b1010 // 10 十进制
	b := 0b1100 // 12 十进制

	fmt.Printf("a = %d (%b)\n", a, a)
	fmt.Printf("b = %d (%b)\n", b, b)

	// 按位与
	andResult := a & b
	fmt.Printf("a & b = %d (%b)\n", andResult, andResult) // 8 (1000)

	// 按位或
	orResult := a | b
	fmt.Printf("a | b = %d (%b)\n", orResult, orResult) // 14 (1110)

	// 按位异或
	xorResult := a ^ b
	fmt.Printf("a ^ b = %d (%b)\n", xorResult, xorResult) // 6 (0110)

	// 按位非
	notResult := ^a
	fmt.Printf("^a = %d (%b)\n", notResult, notResult) // -11 (11110101) // 注意，这是-11的补码表示

	// 左移
	leftShift := a << 2
	fmt.Printf("a << 2 = %d (%b)\n", leftShift, leftShift) // 40 (101000)

	// 右移
	rightShift := b >> 2
	fmt.Printf("b >> 2 = %d (%b)\n", rightShift, rightShift) // 3 (11)

	// 位运算的实际应用示例
	// 1. 判断奇偶
	num := 7
	if num&1 == 1 {
		fmt.Printf("%d 是奇数\n", num)
	}

	// 2. 切换特定位
	flag := 0b1000
	mask := 0b0100
	newFlag := flag ^ mask
	fmt.Printf("切换特定位: %b -> %b\n", flag, newFlag) // 1000 -> 1100

	// 3. 清除特定位
	clearFlag := newFlag &^ mask                        // 等价于 newFlag & ^mask
	fmt.Printf("清除特定位: %b -> %b\n", newFlag, clearFlag) // 1100 -> 1000
}
