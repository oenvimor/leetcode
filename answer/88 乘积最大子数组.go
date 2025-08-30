package answer

func maxProduct(nums []int) int {
	minVal, maxVal := nums[0], nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			minVal, maxVal = maxVal, minVal
		}
		maxVal = Max(nums[i], maxVal*nums[i])
		minVal = Min(nums[i], minVal*nums[i])
		res = Max(res, maxVal)
	}
	return res
}

//[-2,3,-4]
