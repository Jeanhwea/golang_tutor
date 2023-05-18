package nc0061

// 矩阵最长递增路径
func solve(matrix [][]int) (ans int) {
	const S = 1000
	n, m := len(matrix), len(matrix[0])
	dx, dy := []int{1, -1, 0, 0}, []int{0, 0, 1, -1}

	// dfs 搜索最长路径
	var dfs func(x0, y0, count int)
	dfs = func(x0, y0, count int) {
		if count > ans {
			ans = count
		}
		for i := 0; i < 4; i++ {
			x, y := x0+dx[i], y0+dy[i]
			if x >= 0 && x < n && y >= 0 && y < m && matrix[x][y] > matrix[x0][y0] {
				dfs(x, y, count+1)
			}
		}
	}

	for x0 := 0; x0 < n; x0++ {
		for y0 := 0; y0 < m; y0++ {
			dfs(x0, y0, 1)
		}
	}

	return
}
