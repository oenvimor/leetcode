package answer

func numIslands(grid [][]byte) int {
	if len(grid) <= 0 {
		return 0
	}
	rows := len(grid)    // 行
	cols := len(grid[0]) // 列
	counts := 0          // 岛屿数量
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i < 0 || i >= rows || j < 0 || j > cols || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0' // 将岛屿其他部分置为 0
		dfs(i+1, j)      // 上
		dfs(i-1, j)      // 下
		dfs(i, j-1)      // 左
		dfs(i, j+1)      // 右
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				counts++
				dfs(i, j)
			}
		}
	}
	return counts
}
