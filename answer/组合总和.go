package answer

import (
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sort.Ints(candidates)

	var dfs func(start int, target int)
	dfs = func(start int, target int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			dfs(i, target-candidates[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}
