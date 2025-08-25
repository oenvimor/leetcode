package answer

func maximalSquare(matrix [][]byte) int {
	maxLen := 0
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}
	for i := 1; i <= rows; i++ {
		for j := 1; j <= cols; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = MinNum(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
				}
			}

		}
	}
	return maxLen * maxLen
}

func MinNum(num ...int) int {
	minNum := int(1e9)
	for _, v := range num {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}
