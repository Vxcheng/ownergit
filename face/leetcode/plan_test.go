package leetcode

import (
	"fmt"
	"testing"
)

func TestCountSeatSolutions(t *testing.T) {
	// 测试用例
	fmt.Println(countSeatSolutions(2, 2, 2, 2)) // 示例：d=2, t=2, M=2, N=2

	// 更多测试
	fmt.Println(countSeatSolutions(1, 1, 2, 1)) // 不能连续超过1个
	fmt.Println(countSeatSolutions(3, 3, 3, 2))
}
