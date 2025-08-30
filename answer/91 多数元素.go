package answer

import "slices"

//func majorityElement(nums []int) int {
//	res := 0
//	cnt := make(map[int]int)
//	for _, n := range nums {
//		cnt[n]++
//		if cnt[n] > len(nums)/2 {
//			res = n
//		}
//	}
//	return res
//}

func majorityElement(nums []int) int {
	slices.Sort(nums)
	return nums[len(nums)/2]
}
