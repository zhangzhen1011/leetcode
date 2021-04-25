// 计算岛屿数量
// 连通的1的部分
func numIslands(grid [][]byte) int {
	var count int
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				dfs(grid, i, j)
				count++
			}
		}
	}
	return count
}

// 只做标记
func dfs(grid [][]byte, i, j int) {
	m := len(grid)
	n := len(grid[0])

	if i < 0 || j < 0 || i >= m || j >= n {
		return
	}
	if grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2' // 标记为已经访问的大陆
	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}

// tips: 网格图的问题, 搜索标记
// dfs、bfs, 标记连通分量
