package answer

// 贪心算法 + 二分查找
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ast := make([]int, 0)
	ast = append(ast, nums[0])
	for i := 1; i < n; i++ {
		if nums[i] > ast[len(ast)-1] {
			ast = append(ast, nums[i])
		} else {
			pos := binarySearch(ast, 0, len(ast)-1, nums[i])
			ast[pos] = nums[i]
		}
	}
	return len(ast)
}

func binarySearch(nums []int, left, right, target int) int {
	for left <= right {
		mid := (right + left) / 2
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// 动态规划
//func lengthOfLIS(nums []int) int {
//	n := len(nums)
//	if n == 0 {
//		return 0
//	}
//	dp := make([]int, n)
//	for i, _ := range dp {
//		dp[i] = 1
//	}
//	maxLen := 1
//	for i := 1; i < n; i++ {
//		for j := 0; j < i; j++ {
//			if nums[j] < nums[i] {
//				dp[i] = Max(dp[i], dp[j]+1)
//			}
//		}
//		maxLen = Max(dp[i], maxLen)
//	}
//	return maxLen
//}
