package answer

func MinDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for k := 0; k <= m; k++ {
		dp[k][0] = k
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = MinNum(
					dp[i][j-1]+1,   // word1 插入
					dp[i-1][j]+1,   // word1 删除
					dp[i-1][j-1]+1, // word1 替换
				)
			}
		}
	}
	return dp[m][n]
}

//func MinNum(num ...int) int {
//	minNum := int(1e9)
//	for _, v := range num {
//		if v < minNum {
//			minNum = v
//		}
//	}
//	return minNum
//}
