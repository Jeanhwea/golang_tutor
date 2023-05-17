package nc0057

func solve(grid [][]byte) (ans int) {
	n, m := len(grid), len(grid[0])

	visited = make([][]bool, 0, n)
	for i := 0; i < n; i++ {
		visited = append(visited, make([]bool, m, m))
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				dfs(grid, i, j)
				ans++
			}
		}
	}
	return
}

var (
	visited [][]bool
)

func dfs(grid [][]byte, x0, y0 int) {
	n, m := len(grid), len(grid[0])

	dx := [4]int{-1, 1, 0, 0}
	dy := [4]int{0, 0, -1, 1}
	for i := 0; i < 4; i++ {
		x, y := x0+dx[i], y0+dy[i]
		if x >= n || x < 0 || y >= m || y < 0 {
			continue
		}
		if grid[x][y] == '1' && !visited[x][y] {
			visited[x][y] = true
			dfs(grid, x, y)
		}
	}
}
