package nc0068

func minPathSum(matrix [][]int) (ans int) {
	n, m := len(matrix), len(matrix[0])
	dp := [500][500]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i-1 >= 0 && j-1 >= 0 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
			} else if i-1 >= 0 {
				dp[i][j] = dp[i-1][j] + matrix[i][j]
			} else if j-1 >= 0 {
				dp[i][j] = dp[i][j-1] + matrix[i][j]
			} else {
				dp[i][j] = matrix[i][j]
			}
		}
	}
	ans = dp[n-1][m-1]
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
