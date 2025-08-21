package answer

func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	idx := 0
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			idx = mid
			break
		}
	}
	if left > right {
		return []int{-1, -1}
	}
	l := idx - 1
	r := idx + 1
	for l >= 0 && nums[l] == target {
		l--
	}
	for r < len(nums) && nums[r] == target {
		r++
	}
	return []int{l + 1, r - 1}
}
