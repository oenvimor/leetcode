package answer

func SortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// 确保 mid 位置的数字为中间值
	mid := left + (right-left)/2
	if nums[mid] < nums[left] {
		nums[mid], nums[left] = nums[left], nums[mid]
	}
	if nums[right] < nums[left] {
		nums[left], nums[right] = nums[right], nums[left]
	}
	if nums[mid] > nums[right] {
		nums[mid], nums[right] = nums[right], nums[mid]
	}
	nums[left], nums[mid] = nums[mid], nums[left]
	// 选取中间值为基准
	pivot := nums[left]
	i, j := left, right
	for i < j {
		for i < j && nums[j] >= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot
	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
