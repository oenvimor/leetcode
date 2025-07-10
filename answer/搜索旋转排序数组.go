package answer

func Search(nums []int, target int) int {
	if len(nums) <= 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left <= right {
		// 获取中间值的下标
		mid := (left + right) / 2
		// 如果找到则返回
		if nums[mid] == target {
			return mid
		}
		// 左边是有序区间
		if nums[left] <= nums[mid] {
			// 如果目标在左区间
			if nums[left] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 如果目标在右区间
			if nums[mid] <= target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
