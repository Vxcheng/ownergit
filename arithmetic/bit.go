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

type Number struct {
	dec  int8
	orig string // 原码
	rev  string // 反码
	comp string // 补码
}

func computeCodes(n int8) Number {
	u := uint8(n)   // 将 int8 直接转换为 uint8，用于获取负数的二进制补码表示
	abs := uint8(n) // 初始绝对值取值为 n 的无符号形式
	if n < 0 {
		abs = uint8(-n) // 负数时取其绝对值的无符号表示
	}

	// 原码
	signBit := uint8(0) // 先假定符号位为 0
	if n < 0 {
		signBit = 1 << 7 // 负数时设置最高位为 1
	}
	orig := signBit | (abs & 0x7F) // 将符号位与数值部分组合成原码，abs & 0x7F 的作用是清除最高位（符号位），只保留低 7 位数值

	// 反码
	var revCode uint8
	if n >= 0 {
		revCode = orig // 正数反码与原码相同
	} else {
		revCode = (^orig) & 0xFF // 负数反码为原码按位取反，截取 8 位，可能有问题
	}

	// 补码
	var compCode uint8
	if n >= 0 {
		compCode = orig // 正数补码与原码相同
	} else {
		compCode = u // 负数补码直接使用 uint8 转换后的值
	}

	return Number{
		dec:  n,                             // 保存原始十进制值
		orig: fmt.Sprintf("%08b", orig),     // 格式化为 8 位原码字符串
		rev:  fmt.Sprintf("%08b", revCode),  // 格式化为 8 位反码字符串
		comp: fmt.Sprintf("%08b", compCode), // 格式化为 8 位补码字符串
	}
}

func Complement() {
	fmt.Println("十进制 → 原码 → 反码 → 补码")
	for i := int8(-7); i <= 7; i++ {
		n := computeCodes(i)
		fmt.Printf("%3d → %s → %s → %s\n", n.dec, n.orig, n.rev, n.comp)
	}
}
