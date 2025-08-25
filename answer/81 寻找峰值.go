package answer

func FindPeakElement(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}
	left, right := 1, n-2
	for left <= right {
		if nums[left] > nums[left-1] && nums[left] > nums[left+1] {
			return left
		}
		if nums[right] > nums[right-1] && nums[right] > nums[right+1] {
			return right
		}
		left++
		right--
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	return 0
}
