package answer

func NextPermutation(nums []int) {
	p := -1
	n := len(nums)
	for i := n - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			p = i
			break
		}
	}
	if p == -1 {
		reverseV1(nums)
	} else {
		for j := n - 1; j >= 0; j-- {
			if nums[j] > nums[p] {
				nums[j], nums[p] = nums[p], nums[j]
				break
			}
		}
		reverseV1(nums[p+1:])
	}

}

func reverseV1(nums []int) {
	n := len(nums)
	for i := 0; i < n/2; i++ {
		nums[i], nums[n-i-1] = nums[n-i-1], nums[i]
	}
}
