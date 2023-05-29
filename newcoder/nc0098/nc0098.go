package nc0098

// 螺旋矩阵
func spiralOrder(matrix [][]int) (ans []int) {
	n := len(matrix)
	if n <= 0 {
		return
	}
	m := len(matrix[0])
	if m <= 0 {
		return
	}

	visited := [10][10]bool{}
	dx, dy, d := []int{0, 1, 0, -1}, []int{1, 0, -1, 0}, 0
	for i, j := 0, 0; ; {
		ans = append(ans, matrix[i][j])
		visited[i][j] = true
		if len(ans) >= n*m {
			break
		}
		for {
			p, q := i+dx[d], j+dy[d]
			if p >= 0 && p < n && q >= 0 && q < m && !visited[p][q] {
				i, j = p, q
				break
			} else {
				d = (d + 1) % 4
			}
		}
	}
	return
}
