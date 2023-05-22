package nc0075

// 编辑距离(一)
func editDistance(str1 string, str2 string) (ans int) {
	n, m := len(str1), len(str2)
	dp := [1001][1001]int{} // str1[0:i] 和 str2[0:j] 的编辑距离
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1])
			}
		}
	}

	ans = dp[n][m]
	return
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	}
	if b < c {
		return b
	}
	return c
}
