package answer

func MaxSubArray(nums []int) int {
	currSum := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = Max(nums[i], currSum+nums[i])
		maxSum = Max(currSum, maxSum)
	}
	return maxSum
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
