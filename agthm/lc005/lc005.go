package lc005

// 最长回文串
//
//  动态规划解法:
//    dp[i,j] = true/false, 表示 s[i:j+1] 是否是回文串
//
//    dp[i,j] = dp[i+1,j-1] && s[i] == s[j]
//            | true iff i==j
//
func longestPalindrome(s string) (ans string) {
	n := len(s)
	if n <= 0 {
		return
	}

	longest := 1
	ans = s[0:1]

	const N = 1001
	dp := [N][N]bool{}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	for k := 2; k <= n; k++ { // k 表示子串的长度
		for i := 0; i <= n-k; i++ {
			j := i + k - 1
			if s[i] == s[j] && (k == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if k > longest {
					longest = k
					ans = s[i : j+1]
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	return
}

// 暴力解法:
//
//   1. 直接枚举长度 k 的子串
//   2. 判断如果 k 的子串是回文串则返回
//
func longestPalindrome2(s string) (ans string) {
	n := len(s)
	for k := n; k > 0; k-- {
		for i := 0; i <= n-k; i++ {
			found := true
			for p, q := i, i+k-1; p < q; p, q = p+1, q-1 {
				if s[p] != s[q] {
					found = false
					break
				}
			}
			if found {
				ans = s[i : i+k]
				return
			}
		}
	}
	return
}
