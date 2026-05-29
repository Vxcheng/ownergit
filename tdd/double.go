package tdd

/*
题1：
有一种数组压缩机制，将连续的重复数字以【数字，重复次数】的形式记录。
例如，一维数组[1，1，28，28，28，3]压缩后形成二维数组[[1，2]，[28，3]，[3，1]]。
现有两个压缩后的二维数组numsA与numsB，对其进行解压，解压后的两个新数组若长度不同，在较短的数组后补0使两个数
组长度相等。
然后将两个新数组的相同索引的数字相乘，所得结果存人临时数组，最后将临时数组以同样的压缩机制进行压缩后返回，


输入：
numsA=[[3,3],[4,11]
numsB=[[2,1],[4,2],[3,2]
输出：[6,1],[12,3],[0,1]]

解法提示：
*/
//暴力解法
func twoSum(numsA, numsB [][2]int) (res [][2]int) {
	oneA := uncompact(numsA)
	oneB := uncompact(numsB)

	oneA, oneB = fillZero(oneA, oneB)

	tmp := multiply(oneA, oneB)

	res = compact(tmp)
	return
}

// 相乘
func multiply(numsA, numsB []int) (res []int) {
	if len(numsA) == 0 || len(numsB) == 0 {
		return nil
	}

	n := len(numsA)
	if len(numsB) < n {
		n = len(numsB)
	}

	res = make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = numsA[i] * numsB[i]
	}
	return
}

// 压缩
func compact(nums []int) (res [][2]int) {
	if len(nums) == 0 {
		return nil
	}

	currentValue := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == currentValue {
			count++
			continue
		}
		res = append(res, [2]int{currentValue, count})
		currentValue = nums[i]
		count = 1
	}
	res = append(res, [2]int{currentValue, count})
	return
}

// 解压
func uncompact(nums [][2]int) (res []int) {
	for _, item := range nums {
		value := item[0]
		count := item[1]
		if count <= 0 {
			continue
		}
		for i := 0; i < count; i++ {
			res = append(res, value)
		}
	}
	return
}

// 填充0
func fillZero(numsA, numsB []int) (newA, newB []int) {
	if len(numsA) >= len(numsB) {
		extra := len(numsA) - len(numsB)
		newA = append([]int(nil), numsA...)
		newB = append([]int(nil), numsB...)
		for i := 0; i < extra; i++ {
			newB = append(newB, 0)
		}
		return
	}

	extra := len(numsB) - len(numsA)
	newA = append([]int(nil), numsA...)
	newB = append([]int(nil), numsB...)
	for i := 0; i < extra; i++ {
		newA = append(newA, 0)
	}
	return
}

// 双指针解法：直接在压缩格式上按段处理，避免先解压再计算
func doubleSum(numsA, numsB [][2]int) (res [][2]int) {
	i, j := 0, 0
	// 当前段剩余数量
	aRem, bRem := 0, 0
	// 当前段的值
	aVal, bVal := 0, 0

	for {
		// 如果 A 当前段已经耗尽，则取下一段；A 为空时使用 0 扩展
		if aRem == 0 {
			if i < len(numsA) {
				aVal = numsA[i][0]
				aRem = numsA[i][1]
				i++
			} else if j < len(numsB) || bRem > 0 {
				aVal = 0
				if bRem > 0 {
					aRem = bRem // 继续按 B 当前段长度补0
				} else {
					aRem = numsB[j][1] // B 还有后续段，则补0直到该段长度
				}
			} else {
				break
			}
		}

		// 如果 B 当前段已经耗尽，则取下一段；B 为空时使用 0 扩展
		if bRem == 0 {
			if j < len(numsB) {
				bVal = numsB[j][0]
				bRem = numsB[j][1]
				j++
			} else if i < len(numsA) || aRem > 0 {
				bVal = 0
				if aRem > 0 {
					bRem = aRem // 继续按 A 当前段长度补0
				} else {
					bRem = numsA[i][1] // A 还有后续段，则补0直到该段长度
				}
			} else {
				break
			}
		}

		if aRem == 0 && bRem == 0 {
			break
		}

		// 当前重叠区间长度
		minCnt := aRem
		if bRem < minCnt {
			minCnt = bRem
		}

		prod := aVal * bVal
		if len(res) == 0 || res[len(res)-1][0] != prod {
			// 新值与前一段不同，则新增一段
			res = append(res, [2]int{prod, minCnt})
		} else {
			// 新值与前一段相同，则合并计数
			res[len(res)-1][1] += minCnt
		}

		aRem -= minCnt
		bRem -= minCnt
	}

	return
}
