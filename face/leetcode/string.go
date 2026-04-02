package leetcode

import (
	"math"
	"sort"
	"strings"
	"sync"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 || len(intervals) == 1 {
		return intervals
	}
	var result [][]int
	var preRight int
	for _, area := range intervals {
		if preRight == 0 || area[0] > preRight {
			preRight = area[1]
			result = append(result, []int{area[0], area[1]})
			continue
		}
		if area[1] > preRight {
			preRight = area[1]
			result[len(result)-1][1] = area[1]
		}
	}
	return result
}

func LongestDupSubstring(S string) string {
	low, high := 0, len(S)-1
	v := ""
	for low <= high {
		mid := low + (high-low)/2
		v1 := RabinKarp(S, mid)
		if v1 == "" {
			high = mid - 1
		} else {
			low = mid + 1
			v = v1
		}
	}
	return v
}
func RabinKarp(s string, length int) string {
	m := make(map[int]bool)
	now := 0
	r, mod := 256, 6*(1<<20)+1
	for i := 0; i < length; i++ {
		now = ((now * r) + int(s[i])) % mod
	}
	rm := 1
	for i := 1; i < length; i++ {
		rm = (rm * r) % mod
	}
	m[now] = true
	for i := length; i < len(s); i++ {
		now = (now - rm*int(s[i-length])%mod + mod) % mod
		now = (now*r + int(s[i])) % mod
		if m[now] && strings.Contains(s[:i], s[i-length+1:i+1]) {
			return s[i-length+1 : i+1]
		}
		m[now] = true
	}
	return ""
}

func para() {

	count := 5
	vec := make(map[int]int)
	wg := &sync.WaitGroup{}
	// mu := &sync.Mutex{}
	wg.Add(count)
	for i := 0; i < count; i++ {
		// tmp := i
		func() {
			wg.Done()

			// mu.Lock()

			// vec[tmp] = tmp
			vec[i] = i
			// mu.Unlock()
		}()
	}

	wg.Wait()
	print("len: ", len(vec))
	return
}

/*
49. 字母异位词分组

作者：力扣官方题解
链接：https://leetcode.cn/problems/group-anagrams/solutions/520469/zi-mu-yi-wei-ci-fen-zu-by-leetcode-solut-gyoc/

给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
*/
// 原理：1. 字符串排序 2. 哈希表记录相同字符串分组
func groupAnagrams(strs []string) [][]string {
	var res [][]string
	diffM := make(map[string][]string)
	for _, ss := range strs {
		bytes := []byte(ss)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})

		unique := string(bytes)
		// 记录到diffM，比较其它元素，更新到diffM
		diffM[unique] = append(diffM[unique], ss)
	}

	for _, ss := range diffM {
		res = append(res, ss)
	}

	return res
}

// 原理：1. 统计字符串中每个字符出现的次数 2. 哈希表记录相同字符串分组
func groupAnagrams1(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}

/*
128. 最长连续序列
已解答
中等
相关标签
premium lock icon
相关企业
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
*/
func longestConsecutive(nums []int) int {
	// 去重
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	//暴力
	longest := 0
	for num := range numSet {
		if !numSet[num-1] { // 前一位不存在，连续数组开头
			currentNum := num
			currentStack := 1
			for numSet[currentNum+1] { // 后一位存在，继续往后找
				currentNum++
				currentStack++
			}

			if longest < currentStack {
				longest = currentStack
			}
		}
	}

	return longest
}

/*
560. 和为 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。

作者：力扣官方题解
链接：https://leetcode.cn/problems/subarray-sum-equals-k/solutions/238572/he-wei-kde-zi-shu-zu-by-leetcode-solution/
原理：
前缀和 + 哈希表
前缀和：pre[i] = nums[0] + nums[1] + ... + nums[i-1]
则 pre[i] - pre[j] = nums[j] + nums[j+1] + ... + nums[i-1]
如果 pre[i] - pre[j] == k，则说明 nums[j] + nums[j+1] + ... + nums[i-1] == k，即找到了一个和为 k 的子数组。
哈希表：用一个哈希表 m 来记录前缀和出现的次数。初始时，m[0] = 1，表示前缀和为 0 的情况出现了一次。
遍历数组 nums，对于每个元素 nums[i]，计算当前的前缀和 pre，并检查 m 中是否存在 pre - k。如果存在，则说明找到了 m[pre - k] 个和为 k 的子数组，将计数器 count 加上 m[pre - k] 的值。
最后，将当前的前缀和 pre 记录到哈希表 m 中，表示前缀和 pre 出现了一次。
*/
func subarraySum(nums []int, k int) int {
	count, pre := 0, 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if _, ok := m[pre-k]; ok { // pre-k 出现过，说明找到了 m[pre-k] 个和为 k 的子数组
			count += m[pre-k]
		}
		m[pre] += 1
	}
	return count
}

/*
239. 滑动窗口最大值
已解答
困难
相关标签
premium lock icon
相关企业
提示
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

返回 滑动窗口中的最大值 。
作者：力扣官方题解
链接：https://leetcode.cn/problems/sliding-window-maximum/solutions/543426/hua-dong-chuang-kou-zui-da-zhi-by-leetco-ki6m/
简短原理：
使用一个双端队列（deque）来维护当前窗口内的元素索引，确保队列中的索引对应的元素值是递减的。
当窗口向右移动时，首先将新元素的索引加入队列，同时移除队列中所有对应元素值小于新元素的索引，以保持队列的递减性质。
如果队列头部的索引已经不在当前窗口范围内，则将其移除。
每次窗口移动后，队列头部的索引对应的元素就是当前窗口中的最大值。

*/

func maxSlidingWindow(nums []int, k int) []int {
	// q 维护当前窗口中候选最大值的索引，保证对应的 nums 值单调递减
	q := []int{}

	// push 将索引 i 推入双端队列，同时保持队列对应值递减
	push := func(i int) {
		// 当前元素比队尾索引对应元素大时，队尾元素不可能成为窗口内最大值，弹出队尾
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}

	// 先将第一个窗口的前 k 个元素初始化到队列中
	for i := 0; i < k; i++ {
		push(i)
	}

	n := len(nums)
	// 结果切片容量设为 n-k+1，避免多次扩容
	ans := make([]int, 1, n-k+1)
	// 当前窗口的最大值对应队头索引
	ans[0] = nums[q[0]]

	for i := k; i < n; i++ {
		// 将新进元素 i 入队，自动淘汰比它小的后续元素
		push(i)

		// 如果队头索引已经滑出当前窗口，就从队头弹出
		for q[0] <= i-k {
			q = q[1:]
		}

		// 追加当前窗口的最大值
		ans = append(ans, nums[q[0]])
	}
	return ans
}

/*
76. 最小覆盖子串
困难
相关标签
premium lock icon
相关企业
提示
给定两个字符串 s 和 t，长度分别是 m 和 n，返回 s 中的 最短窗口 子串，使得该子串包含 t 中的每一个字符（包括重复字符）。如果没有这样的子串，返回空字符串 ""

作者：力扣官方题解
链接：https://leetcode.cn/problems/minimum-window-substring/solutions/257359/zui-xiao-fu-gai-zi-chuan-by-leetcode-solution/
原理：
滑动窗口（左右指针） + 哈希表（字符计数）
使用滑动窗口的左右指针 [l, r] 来维护一个窗口，初始时 l = 0，r = 0。
使用两个哈希表 ori 和 cnt 来记录 t 中每个字符所需出现的次数和当前窗口中每个字符已经出现的次数。
当 r 向右移动时，如果当前字符在 t 中出现过，则更新窗口计数 cnt。
当当前窗口已经包含 t 中所有字符时，尝试收缩左边界 l，直到窗口不再满足条件。在收缩过程中，记录更短的窗口边界。
最后，如果没有找到满足条件的窗口，返回空字符串；否则返回找到的最短窗口子串。
*/

func minWindow(s string, t string) string {
	// ori 记录 t 中每个字符所需出现的次数
	// cnt 记录当前窗口中每个字符已经出现的次数
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	// minLen 用于记录当前找到的最短窗口长度
	minLen := math.MaxInt32
	ansL, ansR := -1, -1

	// check 判断当前窗口是否已经覆盖了 t 中的所有字符
	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}

	// 使用滑动窗口的左右指针 [l, r]
	for l, r := 0, 0; r < sLen; r++ {
		// 如果当前字符在 t 中出现过，则更新窗口计数
		if ori[s[r]] > 0 {
			cnt[s[r]]++
		}

		// 当当前窗口已经包含 t 中所有字符时，尝试收缩左边界
		for check() && l <= r {
			// 记录更短的窗口边界
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL, ansR = l, l+minLen
			}

			// 减少左边界字符的计数，用于收缩窗口
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]]--
			}
			l++
		}
	}

	// 如果没有找到满足条件的窗口，返回空字符串
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}
