package pkg

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	// 将数组排序
	sort.Ints(nums)
	// 固定第一个数字
	for i := 0; i < len(nums)-2; i++ {
		// 外部去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			// 满足条件则加入返回数组
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 判断下一个数字是否与当前数字相同，相同则跳过
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
			} else if sum < 0 {
				// 如果总和小于0，则左指针右移
				left++
				continue
			} else {
				// 如果总和大于0，则右指针左移
				right--
				continue
			}
			left++
			right--
		}
	}
	return res
}
