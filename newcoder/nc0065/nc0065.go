package nc0065

// 最长公共子序列 (二)
func LCS(s1 string, s2 string) (ans string) {
	n, m := len(s1), len(s2)
	dp := arr2d(n+1, m+1)
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	var str []byte
	for i, j := n, m; dp[i][j] > 0; {
		if s1[i-1] == s2[j-1] {
			str = append(str, s1[i-1])
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}

	ans = string(str)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func arr2d(n, m int) (ans [][]int) {
	for i := 0; i < n; i++ {
		ans = append(ans, make([]int, m))
	}
	return
}
