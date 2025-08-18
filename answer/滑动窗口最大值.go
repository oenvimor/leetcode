package answer

// 暴力解法， 时间复杂度高
//func MaxSlidingWindow(nums []int, k int) []int {
//	if len(nums) == 0 {
//		return []int{}
//	}
//	res := make([]int, 0)
//	n := len(nums) - k + 1
//	for i := 0; i < n; i++ {
//		res = append(res, getMax(nums[i:i+k]))
//	}
//	return res
//}
//
//func getMax(nums []int) int {
//	var i int
//	n := len(nums)
//	tmp := make([]int, n)
//	copy(tmp, nums)
//	for i = 0; i < n-1; i++ {
//		if tmp[i] > tmp[i+1] {
//			tmp[i], tmp[i+1] = tmp[i+1], tmp[i]
//		}
//	}
//	return tmp[i]
//}

func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	res := make([]int, 0)
	deQueue := make([]int, 0)

	n := len(nums)
	for i := 0; i < n; i++ {
		if len(deQueue) > 0 && deQueue[0] <= i-k {
			deQueue = deQueue[1:]
		}
		for len(deQueue) > 0 && nums[deQueue[len(deQueue)-1]] < nums[i] {
			deQueue = deQueue[:len(deQueue)-1]
		}
		deQueue = append(deQueue, i)
		if i >= k-1 {
			res = append(res, nums[deQueue[0]])
		}
	}
	return res
}
