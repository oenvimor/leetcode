package answer

func TwoSum(nums []int, target int) []int {
	// 创建一个 map 用作 hash 表
	numsMap := make(map[int]int)
	// 遍历数组
	for i, num := range nums {
		if idx, ok := numsMap[target-num]; ok {
			return []int{idx, i}
		}
		numsMap[num] = i
	}
	return nil
}
