package answer

func MinPathSum(grid [][]int) int {
	length := len(grid)
	width := len(grid[0])
	dp := make([][]int, length)
	for i, _ := range dp {
		dp[i] = make([]int, width)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < length; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < width; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for m := 1; m < length; m++ {
		for n := 1; n < width; n++ {
			dp[m][n] = grid[m][n] + Min(dp[m-1][n], dp[m][n-1])
		}
	}
	return dp[length-1][width-1]
}

//func Min(a, b int) int {
//	if a > b {
//		return b
//	} else {
//		return a
//	}
//}
