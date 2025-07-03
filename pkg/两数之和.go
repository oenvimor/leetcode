package pkg

func twoSum(nums []int, target int) []int {
	// 循环遍历数组，将其中的数字两两相加，然后与target进行比较得出结果
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = nums[i]
				result[1] = nums[j]
			}
		}
	}
	return result
}
