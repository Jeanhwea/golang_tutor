package nc0076

// 正则表达式匹配

// 递归解法
func match(str string, pat string) (ans bool) {
	n, m := len(str), len(pat)

	switch m {
	case 0:
		return n == 0
	case 1:
		if n != 1 {
			return false
		}
		return pat[0] == '.' || pat[0] == str[0]
	default: // m >= 2
		if pat[1] == '*' {
			if n == 0 {
				return match(str, pat[2:])
			}

			if pat[0] == '.' || pat[0] == str[0] {
				return match(str[1:], pat) || match(str, pat[2:])
			} else {
				return match(str, pat[2:])
			}
		} else {
			if n == 0 {
				return false
			}
			if pat[0] == '.' || pat[0] == str[0] {
				return match(str[1:], pat[1:])
			}
		}
	}

	return false
}

// 动规解法
func match2(str string, pat string) (ans bool) {
	n, m := len(str), len(pat)
	dp := [27][27]bool{} // dp[i][j] 表示 str[0:i] 是否匹配 pat[0:j]

	// pat 和 str 都为空结果为 true
	dp[0][0] = true

	// 情形一: pat 为空, str 长度大于 0 都是 false
	for i := 1; i <= n; i++ {
		dp[i][0] = false
	}

	// 情形二: str 空, pat 长度大于 0
	for j := 1; j <= m; j++ {
		if pat[j-1] == '*' {
			if j == 2 {
				dp[0][j] = true
			} else {
				dp[0][j] = dp[0][j-2]
			}
		}
	}

	// 迭代计算其它场景
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if pat[j-1] == '*' {
				if j-2 < 0 {
					continue
				}

				// 忽略 x* 字符串, 可以得到第一 case
				dp[i][j] = dp[i][j] || dp[i][j-2]

				// 如果满足以下条件, 可能匹配 str[i-1] 或者忽略 str[i-1]
				if pat[j-2] == '.' || pat[j-2] == str[i-1] {
					dp[i][j] = dp[i][j] || dp[i-1][j-2] || dp[i-1][j]
				}
			} else if pat[j-1] == '.' || pat[j-1] == str[i-1] {
				// 如果没有 * 字符, 则只需匹配一个字符
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			} else {
				dp[i][j] = false
			}
		}
	}

	ans = dp[n][m]
	return
}
