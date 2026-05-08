package built

import "testing"

// RED: And - 按位与
func TestAnd(t *testing.T) {
	a := 2 * (1 << 3) // 2 * 16 = 32
	t.Log(a)

	if And(0b1010, 0b1100) != 0b1000 {
		t.Errorf("And(0b1010, 0b1100) want 0b1000, got %b", And(0b1010, 0b1100))
	}
	if And(0xFF, 0x0F) != 0x0F {
		t.Errorf("And(0xFF, 0x0F) want 0x0F, got %x", And(0xFF, 0x0F))
	}
}

// RED: Or - 按位或
func TestOr(t *testing.T) {
	if Or(0b1010, 0b0101) != 0b1111 {
		t.Errorf("Or(0b1010, 0b0101) want 0b1111, got %b", Or(0b1010, 0b0101))
	}
	if Or(0, 0xFF) != 0xFF {
		t.Errorf("Or(0, 0xFF) want 0xFF, got %x", Or(0, 0xFF))
	}
}

// RED: Xor - 按位异或
func TestXor(t *testing.T) {
	if Xor(0b1010, 0b1100) != 0b0110 {
		t.Errorf("Xor(0b1010, 0b1100) want 0b0110, got %b", Xor(0b1010, 0b1100))
	}
	// 相同数异或为0
	if Xor(42, 42) != 0 {
		t.Errorf("Xor(42, 42) want 0, got %d", Xor(42, 42))
	}
}

// RED: Not - 按位取反（uint8范围）
func TestNot8(t *testing.T) {
	if Not8(0b00001111) != 0b11110000 {
		t.Errorf("Not8(0b00001111) want 0b11110000, got %b", Not8(0b00001111))
	}
	if Not8(0x00) != 0xFF {
		t.Errorf("Not8(0x00) want 0xFF, got %x", Not8(0x00))
	}
}

// RED: LeftShift - 左移
func TestLeftShift(t *testing.T) {
	if LeftShift(1, 3) != 8 {
		t.Errorf("LeftShift(1, 3) want 8, got %d", LeftShift(1, 3))
	}
	if LeftShift(0b0001, 4) != 0b10000 {
		t.Errorf("LeftShift(0b0001, 4) want 0b10000, got %b", LeftShift(0b0001, 4))
	}
}

// RED: RightShift - 右移
func TestRightShift(t *testing.T) {
	if RightShift(8, 3) != 1 {
		t.Errorf("RightShift(8, 3) want 1, got %d", RightShift(8, 3))
	}
	if RightShift(0b10000, 4) != 0b0001 {
		t.Errorf("RightShift(0b10000, 4) want 0b0001, got %b", RightShift(0b10000, 4))
	}
}

// RED: SetBit - 将第n位设为1（从0开始）
func TestSetBit(t *testing.T) {
	if SetBit(0b0000, 2) != 0b0100 {
		t.Errorf("SetBit(0b0000, 2) want 0b0100, got %b", SetBit(0b0000, 2))
	}
	if SetBit(0b0101, 1) != 0b0111 {
		t.Errorf("SetBit(0b0101, 1) want 0b0111, got %b", SetBit(0b0101, 1))
	}
}

// RED: ClearBit - 将第n位清为0
func TestClearBit(t *testing.T) {
	if ClearBit(0b1111, 2) != 0b1011 {
		t.Errorf("ClearBit(0b1111, 2) want 0b1011, got %b", ClearBit(0b1111, 2))
	}
	if ClearBit(0b0100, 2) != 0b0000 {
		t.Errorf("ClearBit(0b0100, 2) want 0b0000, got %b", ClearBit(0b0100, 2))
	}
}

// RED: ToggleBit - 翻转第n位
func TestToggleBit(t *testing.T) {
	if ToggleBit(0b0000, 3) != 0b1000 {
		t.Errorf("ToggleBit(0b0000, 3) want 0b1000, got %b", ToggleBit(0b0000, 3))
	}
	if ToggleBit(0b1000, 3) != 0b0000 {
		t.Errorf("ToggleBit(0b1000, 3) want 0b0000, got %b", ToggleBit(0b1000, 3))
	}
}

// RED: TestBit - 检测第n位是否为1
func TestTestBit(t *testing.T) {
	if !TestBit(0b1010, 1) {
		t.Errorf("TestBit(0b1010, 1) want true")
	}
	if TestBit(0b1010, 0) {
		t.Errorf("TestBit(0b1010, 0) want false")
	}
}

// RED: CountOnes - 统计二进制中1的个数
func TestCountOnes(t *testing.T) {
	if CountOnes(0b1010) != 2 {
		t.Errorf("CountOnes(0b1010) want 2, got %d", CountOnes(0b1010))
	}
	if CountOnes(0xFF) != 8 {
		t.Errorf("CountOnes(0xFF) want 8, got %d", CountOnes(0xFF))
	}
	if CountOnes(0) != 0 {
		t.Errorf("CountOnes(0) want 0, got %d", CountOnes(0))
	}
}

// RED: IsPowerOfTwo - 判断是否为2的幂
func TestIsPowerOfTwo(t *testing.T) {
	for _, n := range []int{1, 2, 4, 8, 16, 1024} {
		if !IsPowerOfTwo(n) {
			t.Errorf("IsPowerOfTwo(%d) want true", n)
		}
	}
	for _, n := range []int{0, 3, 5, 6, 7, 100} {
		if IsPowerOfTwo(n) {
			t.Errorf("IsPowerOfTwo(%d) want false", n)
		}
	}
}

// RED: LowBit - 取最低位的1（lowbit，树状数组核心操作）
func TestLowBit(t *testing.T) {
	if LowBit(0b1100) != 0b0100 {
		t.Errorf("LowBit(0b1100) want 0b0100, got %b", LowBit(0b1100))
	}
	if LowBit(0b1010) != 0b0010 {
		t.Errorf("LowBit(0b1010) want 0b0010, got %b", LowBit(0b1010))
	}
	if LowBit(8) != 8 {
		t.Errorf("LowBit(8) want 8, got %d", LowBit(8))
	}
}

// RED: ReverseBits8 - 反转8位整数的所有位
func TestReverseBits8(t *testing.T) {
	if ReverseBits8(0b10000000) != 0b00000001 {
		t.Errorf("ReverseBits8(0b10000000) want 0b00000001, got %b", ReverseBits8(0b10000000))
	}
	if ReverseBits8(0b10110001) != 0b10001101 {
		t.Errorf("ReverseBits8(0b10110001) want 0b10001101, got %b", ReverseBits8(0b10110001))
	}
}

// RED: SwapWithoutTemp - 不用临时变量交换两个整数
func TestSwapWithoutTemp(t *testing.T) {
	a, b := SwapWithoutTemp(3, 5)
	if a != 5 || b != 3 {
		t.Errorf("SwapWithoutTemp(3,5) want (5,3), got (%d,%d)", a, b)
	}
	a, b = SwapWithoutTemp(0, 100)
	if a != 100 || b != 0 {
		t.Errorf("SwapWithoutTemp(0,100) want (100,0), got (%d,%d)", a, b)
	}
}
