package answer

func SubarraySum(nums []int, k int) int {
	count := 0
	preNum := 0
	cnt := make(map[int]int)
	cnt[0] = 1
	for i := 0; i < len(nums); i++ {
		preNum += nums[i]
		if _, found := cnt[preNum-k]; found {
			count += cnt[preNum-k]
		}
		cnt[preNum]++
	}
	return count
}
