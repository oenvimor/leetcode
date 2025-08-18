package answer

import "math"

func CoinChange(coins []int, amount int) int {
	INF := math.MaxInt32
	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != INF {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}
	if dp[amount] == INF {
		return -1
	}
	return dp[amount]
}
