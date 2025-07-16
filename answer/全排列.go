package answer

func Permute(nums []int) [][]int {
	n := len(nums)
	if n <= 0 {
		return [][]int{}
	}
	used := make([]bool, n)
	path := make([]int, 0, n)
	ans := make([][]int, 0)
	var backtrack func()
	backtrack = func() {
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return ans
}
