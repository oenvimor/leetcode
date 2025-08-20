package answer

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(start int)
	dfs = func(start int) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}
