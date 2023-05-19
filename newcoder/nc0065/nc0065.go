package nc0065

const (
	MAX_LEN = 2001
)

// 最长公共子序列 (二)
func LCS(s1 string, s2 string) (ans string) {
	n, m := len(s1), len(s2)
	dp := [MAX_LEN][MAX_LEN]int{}
	choose := [MAX_LEN]bool{}
	for k := 0; k < max(n, m)*2+1; k++ {
		for i, j := 0, k-1; i >= 0 && j >= 0; i, j = i+1, j-1 {
			if i > n || j > m {
				continue
			}
			maxlen := 0
			if i-1 >= 0 && dp[i-1][j] > maxlen {
				maxlen = dp[i-1][j]
			}
			if j-1 >= 0 && dp[i][j-1] > maxlen {
				maxlen = dp[i][j-1]
			}
			if i-1 >= 0 && j-1 >= 0 {
				b := dp[i-1][j-1]
				if s1[i-1] == s2[j-1] {
					b++
				}
				if b > maxlen {
					maxlen = b
					choose[i-1] = true
				}
			}
			dp[i][j] = maxlen
		}
	}
	for i := 0; i < n; i++ {
		if choose[i] {
			ans = ans + string(s1[i])
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
