package lc0064

// 最小路径和
//  采用动态规划的解法
//  dp[i][j] 表示走到 grid[i][j] 的最小路径和
//  dp[i][j] := grid[i][j] + max(dp[i-1][j], dp[i][j-1])   i > 0 && j > 0
//           |= grid[i][j] + dp[i][j-1]                    i == 0
//           |= grid[i][j] + dp[i-1][j]                    j == 0
//
func minPathSum(grid [][]int) (ans int) {
	const N = 200
	dp := [N][N]int{}
	n, m := len(grid), len(grid[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = grid[i][j] + min(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	ans = dp[n-1][m-1]
	return
}

func min(a, b int) (ans int) {
	if a < b {
		return a
	}
	return b
}
