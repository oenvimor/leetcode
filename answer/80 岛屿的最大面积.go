package answer

func MaxAreaOfIsland(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	maxCount := 0
	tmp := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= rows || j < 0 || j >= cols || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		tmp++
		if tmp > maxCount {
			maxCount = tmp
		}
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				tmp = 0
			}
		}
	}
	return maxCount
}
