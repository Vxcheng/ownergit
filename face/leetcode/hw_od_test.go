package leetcode

import (
	"fmt"
	"testing"
)

func TestLedController(t *testing.T) {
	// 测试用例
	fmt.Println(ledController("L0L1L2T1"))    // 输出: 5
	fmt.Println(ledController("L0L1L2D1"))    // 输出: 5 (L0=1,L1=1,L2=1,然后D1熄灭L1→101)
	fmt.Println(ledController(""))            // 空指令，输出: 0
	fmt.Println(ledController("L7"))          // 点亮7号灯 → 10000000 = 128
	t.Logf("%b\n", ledController("L0L7T0T7")) // 点亮0和7号灯，然后切换0和7 → 00000000 = 0
}

func TestResolution(t *testing.T) {
	// 测试用例1
	input := "1920x1080 1280x720 3840x2160 2560x1440"
	// 期望输出: 3840x2160 2560x1440 1920x1080 1280x720

	// 测试用例2
	// input := "2600x1400 1920x1080 1280x720"
	// 2600x1400: 宽≥1920且高≥1080 → 1080P, 面积3640000
	// 1920x1080: 1080P, 面积2073600
	// 1280x720: 720P, 面积921600
	// 期望: 2600x1400 1920x1080 1280x720
	v := resolution(input)
	t.Logf("Result: %s\n", v)
}

func TestMainAP(t *testing.T) {
	mainAP()
}
