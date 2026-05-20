package leetcode

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*
P14201.8位LED控制器

题目理解
题目模拟一个8位LED控制器，8个灯编号0-7（0是最低位，7是最高位），初始状态全部熄灭，即二进制 00000000，对应整数 0。

三种指令：

# Lx：点亮第 x 号灯 → 对应位设为 1

# Dx：熄灭第 x 号灯 → 对应位设为 0

Tx：切换第 x 号灯 → 对应位取反（0变1，1变0）

输入一个指令字符串，按顺序执行后输出最终状态的整数值。

示例："L0L1L2T1"

# L0 → 灯0亮：00000001

# L1 → 灯1亮：00000011

# L2 → 灯2亮：00000111

T1 → 灯1切换：00000101 → 二进制101 = 整数 5

核心考点
这道题主要考察三个点：

位运算基础：用整数的低8位模拟8个LED状态

字符串解析：每2个字符为一条指令，第一个字符是操作，第二个字符是灯编号

运算对应：

点亮（置1）：state |= (1 << x)

熄灭（置0）：state &= ~(1 << x)

切换（取反）：state ^= (1 << x)

时间复杂度：O(n)，其中 n 是指令字符串的长度（每2个字符处理一次）
空间复杂度：O(1)，只使用了常数空间来存储状态
*/
func ledController(commands string) int {
	state := 0
	n := len(commands)

	// 每两个字符为一条指令
	for i := 0; i+1 < n; i += 2 {
		op := commands[i]             // 操作符: 'L', 'D', 'T'
		x := int(commands[i+1] - '0') // 灯编号: 0-7

		switch op {
		case 'L':
			// 点亮：对应位置1
			state |= (1 << x)
		case 'D':
			// 熄灭：对应位置0
			state &= ^(1 << x)
		case 'T':
			// 切换：对应位取反
			state ^= (1 << x)
		}
	}

	return state
}

/*
一、题目理解
给定多个分辨率字符串（如 "1920x1080"），需要按以下规则从大到小排序：

1. 清晰度映射规则
清晰度大小：720P < 1080P < 2K < 4K

匹配规则：宽且高同时 ≥ 清晰度标准时，才满足该清晰度，且优先匹配更高清晰度。

清晰度	标准宽	标准高
720P	1280	720
1080P	1920	1080
2K	2560	1440
4K	3840	2160
特殊情况：

所有低于720P的 → 归为720P

所有满足4K的（宽≥3840且高≥2160） → 归为4K，即使更大也仍是4K

不考虑交换宽高（即 2500x3200 不能当2K匹配，因为没有3200这个标准）

2. 排序规则（优先级从高到低）
清晰度等级：4K > 2K > 1080P > 720P

面积（宽×高）：面积大的排前面

宽度：宽大的排前面

3. 输入输出
输入："1920x1080 1280x720 3840x2160 2560x1440"

输出：从大到小排序后的分辨率字符串，空格分隔

二、解题思路
解析每个分辨率字符串，得到宽和高

根据宽高判断清晰度等级（从高到低判断：4K→2K→1080P→720P）

定义排序规则：先比等级，再比面积，最后比宽度

排序后输出

时间复杂度：O(n log n)，其中 n 是分辨率的数量（排序的时间复杂度）
空间复杂度：O(n)，用于存储分辨率信息
*/

// Resolution 分辨率结构体
type Resolution struct {
	Width  int
	Height int
	Str    string
	Level  int // 清晰度等级: 0=720P, 1=1080P, 2=2K, 3=4K
	Area   int // 面积
}

// getLevel 根据宽高获取清晰度等级
func getLevel(width, height int) int {
	// 从高到低匹配
	// 4K: 3840x2160
	if width >= 3840 && height >= 2160 {
		return 3
	}
	// 2K: 2560x1440
	if width >= 2560 && height >= 1440 {
		return 2
	}
	// 1080P: 1920x1080
	if width >= 1920 && height >= 1080 {
		return 1
	}
	// 低于720P也算720P
	return 0
}

// sortResolutions 排序分辨率
func sortResolutions(resolutions []Resolution) []Resolution {
	sort.Slice(resolutions, func(i, j int) bool {
		// 按清晰度降序
		if resolutions[i].Level != resolutions[j].Level {
			return resolutions[i].Level > resolutions[j].Level
		}
		// 按面积降序
		if resolutions[i].Area != resolutions[j].Area {
			return resolutions[i].Area > resolutions[j].Area
		}
		// 按宽度降序
		return resolutions[i].Width > resolutions[j].Width
	})
	return resolutions
}

func resolution(input string) string {
	parts := strings.Split(input, " ")
	resolutions := make([]Resolution, len(parts))

	for i, part := range parts {
		// 解析 "宽x高"
		wh := strings.Split(part, "x")
		width, _ := strconv.Atoi(wh[0])
		height, _ := strconv.Atoi(wh[1])

		resolutions[i] = Resolution{
			Width:  width,
			Height: height,
			Str:    part,
			Level:  getLevel(width, height),
			Area:   width * height,
		}
	}

	// 排序
	sorted := sortResolutions(resolutions)

	// 输出结果
	result := make([]string, len(sorted))
	for i, r := range sorted {
		result[i] = r.Str
	}
	fmt.Println(strings.Join(result, " "))
	return strings.Join(result, " ")
}

/*
od笔试题：P14203.wif设备网络规划

一、题目理解
核心规则：

网格大小 ≤ 50×50

# 表示墙壁（不能放AP）

. 表示空地（必须被覆盖）

每个AP覆盖自身为中心的 3×3 区域（包括上下左右及对角）

任意两个AP的覆盖区域不能重叠

输出最少AP数，若无法覆盖返回 -1

关键约束：覆盖区域不能重叠，这意味着每个格子最多只能被一个AP覆盖。

二、解题思路
这个问题本质上是带约束的集合覆盖问题，由于网格规模小，可以采用DFS + 回溯 + 剪枝的方法：

遍历所有可能的AP放置位置（空地格子）

每放置一个AP，标记该AP的3×3覆盖区域

确保新AP的覆盖区域与已放置AP的覆盖区域无重叠

记录覆盖所有空地所需的最小AP数量
*/

func minWifiAPs(grid [][]byte) int {
	rows, cols := len(grid), len(grid[0])

	// 统计空地数量
	emptyCount := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '.' {
				emptyCount++
			}
		}
	}
	if emptyCount == 0 {
		return 0
	}

	// 记录每个格子是否被覆盖
	covered := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		covered[i] = make([]bool, cols)
	}

	minAPs := math.MaxInt32

	// 检查位置 (x,y) 放置AP是否会与现有覆盖区域重叠
	isOverlap := func(x, y int) bool {
		for i := max(0, x-1); i <= min(rows-1, x+1); i++ {
			for j := max(0, y-1); j <= min(cols-1, y+1); j++ {
				if covered[i][j] {
					return true
				}
			}
		}
		return false
	}

	// 放置AP，标记覆盖区域
	placeAP := func(x, y int) {
		for i := max(0, x-1); i <= min(rows-1, x+1); i++ {
			for j := max(0, y-1); j <= min(cols-1, y+1); j++ {
				if grid[i][j] == '.' {
					covered[i][j] = true
				}
			}
		}
	}

	// 移除AP，清除覆盖标记（需要知道之前哪些格子是被这个AP覆盖的）
	removeAP := func(x, y int, beforeState [][]bool) {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				covered[i][j] = beforeState[i][j]
			}
		}
	}

	// 检查是否所有空地都被覆盖
	allCovered := func() bool {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				if grid[i][j] == '.' && !covered[i][j] {
					return false
				}
			}
		}
		return true
	}

	// 复制覆盖状态
	copyCovered := func() [][]bool {
		newCovered := make([][]bool, rows)
		for i := 0; i < rows; i++ {
			newCovered[i] = make([]bool, cols)
			copy(newCovered[i], covered[i])
		}
		return newCovered
	}

	var dfs func(idx, count int)
	dfs = func(idx, count int) {
		if count >= minAPs {
			return
		}

		if allCovered() {
			if count < minAPs {
				minAPs = count
			}
			return
		}

		if idx >= rows*cols {
			return
		}

		x, y := idx/cols, idx%cols

		// 跳过墙壁
		if grid[x][y] == '#' {
			dfs(idx+1, count)
			return
		}

		// 如果当前格子已被覆盖，可以跳过放置
		if covered[x][y] {
			dfs(idx+1, count)
		}

		// 尝试在此位置放置AP
		if !isOverlap(x, y) {
			beforeState := copyCovered()
			placeAP(x, y)
			dfs(idx+1, count+1)
			removeAP(x, y, beforeState)
		}

		// 不在此位置放置AP（跳过此格子）
		dfs(idx+1, count)
	}

	dfs(0, 0)

	if minAPs == math.MaxInt32 {
		return -1
	}
	return minAPs
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func mainAP() {
	// 测试用例1
	grid1 := [][]byte{
		{'.', '.', '.', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.'},
		{'.', '.', '.', '#', '.', '.', '.'},
	}
	fmt.Println(minWifiAPs(grid1)) // 预期输出: 6

	// 测试用例2
	grid2 := [][]byte{
		{'.', '.', '#', '.', '.'},
		{'.', '.', '#', '.', '.'},
		{'.', '.', '#', '.', '.'},
		{'.', '.', '#', '.', '.'},
		{'.', '.', '#', '.', '.'},
	}
	fmt.Println(minWifiAPs(grid2)) // 预期输出: 4

	// 测试用例3
	grid3 := [][]byte{
		{'.'}, {'.'}, {'.'}, {'.'}, {'.'},
	}
	fmt.Println(minWifiAPs(grid3)) // 预期输出: 2

	// 测试用例4（无法覆盖）
	grid4 := [][]byte{
		{'.', '#', '.', '#'},
		{'#', '.', '#', '.'},
		{'.', '#', '.', '#'},
		{'#', '.', '#', '.'},
	}
	fmt.Println(minWifiAPs(grid4)) // 预期输出: -1
}

/*
a(b)(

(ab()
找到剔除没成对的括号，最多一个

*/

func handleStr1(context string) (display string) {
	// 只删除一个不成对的括号
	// 使用栈来跟踪未匹配的左括号索引，
	// 匹配到右括号时弹出栈顶，如果遇到无法匹配的右括号，直接删除它；
	// 如果最后栈中还有未匹配的左括号，删除最后一个未匹配的左括号。
	stack := make([]int, 0, len(context))
	for i, r := range context {
		switch r {
		case '(':
			stack = append(stack, i)
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // 匹配到一个左括号，弹出栈顶
			} else {
				return context[:i] + context[i+1:]
			}
		}
	}

	if len(stack) > 0 {
		idx := stack[len(stack)-1]
		fmt.Printf("idx: %d\n", idx)
		return context[:idx] + context[idx+1:]
	}

	return context
}
