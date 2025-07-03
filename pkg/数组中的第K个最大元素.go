package pkg

func FindKthLargest(nums []int, k int) int {
	n := len(nums)
	return QuickSelect(nums, 0, n-1, n-k)
}

func QuickSelect(nums []int, left, right, targetIdx int) int {
	if left == right {
		return nums[left]
	}
	// 指定最左边的值为基准
	partition := nums[left]
	// 初始化左右指针
	i := left - 1
	j := right + 1
	// 当左指针小于右指针时，交换两个指针指向的元素
	for i < j {
		// 如果左指针的值小于基准值，则左指针向右移动一位
		for i++; nums[i] < partition; i++ {
		}
		// 如果右指针的值大于基准值，则右指针向左移动一位
		for j--; nums[j] > partition; j-- {
		}
		// 如果左指针小于右指针，交换两个指针指向的元素
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 当目标下标大于j，就到右分区去找；否则，就到左分区去找
	if j < targetIdx {
		return QuickSelect(nums, j+1, right, targetIdx)
	} else {
		return QuickSelect(nums, left, j, targetIdx)
	}
}

// QuickSelectTimeOut 超时
func QuickSelectTimeOut(nums []int, left, right, targetIdx int) int {
	if left == right {
		return nums[left]
	}
	p := nums[right]
	n := len(nums)
	i := left
	for j := left; j < n-1; j++ {
		if nums[j] < p {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	if i == targetIdx {
		return nums[i]
	} else if i < targetIdx {
		return QuickSelect(nums, i+1, right, targetIdx)
	} else if i > targetIdx {
		return QuickSelect(nums, left, i-1, targetIdx)
	}
	return 0
}
