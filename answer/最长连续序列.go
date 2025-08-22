package answer

import "sort"

func LongestConsecutive(nums []int) int {
	//if len(nums) == 0 {
	//	return 0
	//}
	//set := make(map[int]bool)
	//for _, num := range nums {
	//	set[num] = true
	//}
	//longest := 0
	//for num := range set {
	//	if !set[num-1] {
	//		length := 1
	//		cur := num
	//		for set[cur+1] {
	//			cur++
	//			length++
	//		}
	//		if length > longest {
	//			longest = length
	//		}
	//	}
	//}
	//return longest

	if len(nums) == 0 {
		return 0
	}
	longest := 0
	cur := 1
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cur++
		} else if nums[i] == nums[i-1] {

		} else {
			cur = 1
		}
		if cur > longest {
			longest = cur
		}
	}
	return longest
}
