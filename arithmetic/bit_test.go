package arithmetic

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Run("AND", func(t *testing.T) {
		// 负数转换为补码再按位运算

		fmt.Printf("%08b, %08b \n", 3, -1)
		v := 3 & 1
		fmt.Printf("%v, %08b \n", v, v)
		v = 3 & -1
		fmt.Printf("%v, %08b \n", v, v)
		v = -3 & 1
		fmt.Printf("%v, %08b \n", v, v)
	})
	t.Run("OR", func(t *testing.T) {
		v := 3 | 2
		fmt.Printf("%v, %08b \n", v, v)
		v = 3 | -2 // 0010, 1110
		fmt.Printf("%v, %08b \n", v, v)
		v = -3 | 1 // 1101 | 0001
		fmt.Printf("%v, %08b \n", v, v)
	})

	// 对一个数 a 进行按位非运算，结果等于 -(a + 1)。
	t.Run("NOT", func(t *testing.T) {
		v := ^10
		fmt.Printf("%v, %08b \n", v, v)
		v = ^-10
		fmt.Printf("%v, %08b \n", v, v)
	})

	t.Run("Shift Left", func(t *testing.T) {
		v := 5 << 1
		fmt.Printf("%v, %08b \n", v, v)
		v = -5 << 1
		fmt.Printf("%v, %08b \n", v, v)
	})

	t.Run("Shift Right", func(t *testing.T) {
		v := 5 >> 1
		fmt.Printf("%v, %08b \n", v, v)
		v = -5 >> 1 
		fmt.Printf("%v, %08b \n", v, v)
	})

	//规则：两个位不同时结果为 1，相同时为 0。
	/*
			用途：
		翻转特定比特位（如 num ^ mask 可将 mask 中为 1 的位翻转）。
		交换两个数（无需临时变量：a = a ^ b; b = a ^ b; a = a ^ b;）。
		判断两个数是否相等（a ^ b == 0 则相等）, a^a=0, a^0=a, a^b=b^a
	*/
	t.Run("XOR", func(t *testing.T) {
		v := 3 ^ 1
		fmt.Printf("%v, %08b \n", v, v)
		v = 3 ^ -1 // 将3的所有位取反，等价于-4
		fmt.Printf("%v, %08b \n", v, v)
	})

	// 位清零, x &^ y 等价于 x & (^y)
	//核心记忆口诀：y 是掩码。y 中为 1 的位是“清除指令”，会把 x 中对应的位清零。y 中为 0 的位是“保留指令”，会保留 x 中对应的位。
	t.Run("AND NOT", func(t *testing.T) {
		var x byte = 0b10101010 // 170
		var y byte = 0b11110000 // 240
		// 目标是：清除 x 的高 4 位，保留低 4 位

		result := x &^ y
		fmt.Printf("x     = %08b\n", x)
		fmt.Printf("y     = %08b\n", y)
		fmt.Printf("result= %08b\n", result)
	})

	t.Run("BitDemo", func(t *testing.T) {
		BitDemo()
	})
}
