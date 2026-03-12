package leetcode

// countSeatSolutions 统计将 M 个 D 类和 N 个 T 类排成一排的合法方案数。
// 约束：D 类连续不超过 d 个，T 类连续不超过 t 个。
// 状态：dp[i][j][k][l] = 已放 i 个D、j 个T，末尾类型为 k（0=D,1=T）且连续 l 个时的方案数。
func countSeatSolutions(d int, t int, M int, N int) int {
	// 第三维长度依末尾类型而定：D 最多连续 d 个，T 最多连续 t 个
	dp := make([][][][]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([][][]int, N+1)
		for j := 0; j <= N; j++ {
			dp[i][j] = make([][]int, 2)
			for k := 0; k < 2; k++ {
				maxLen := d
				if k == 1 {
					maxLen = t
				}
				dp[i][j][k] = make([]int, maxLen+1)
			}
		}
	}

	// 边界：队列只有 1 个人的两种起始状态
	if M > 0 {
		dp[1][0][0][1] = 1
	}
	if N > 0 {
		dp[0][1][1][1] = 1
	}

	for i := 0; i <= M; i++ {
		for j := 0; j <= N; j++ {
			for k := 0; k < 2; k++ {
				maxLen := d
				if k == 1 {
					maxLen = t
				}
				for l := 1; l <= maxLen; l++ {
					if dp[i][j][k][l] == 0 {
						continue
					}
					val := dp[i][j][k][l]

					if k == 0 { // 末尾是 D
						if l < d && i < M { // 继续放 D，连续计数 +1
							dp[i+1][j][0][l+1] += val
						}
						if j < N { // 切换放 T，连续计数重置为 1
							dp[i][j+1][1][1] += val
						}
					} else { // 末尾是 T
						if l < t && j < N { // 继续放 T，连续计数 +1
							dp[i][j+1][1][l+1] += val
						}
						if i < M { // 切换放 D，连续计数重置为 1
							dp[i+1][j][0][1] += val
						}
					}
				}
			}
		}
	}

	// 累加所有 D、T 恰好用完的终态方案数
	result := 0
	for k := 0; k < 2; k++ {
		maxLen := d
		if k == 1 {
			maxLen = t
		}
		for l := 1; l <= maxLen; l++ {
			result += dp[M][N][k][l]
		}
	}
	return result
}
