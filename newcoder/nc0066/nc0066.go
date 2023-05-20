package nc0066

// 最长公共子串

// 动态规划: dp[i][j] 表示以 str1[i] 和 str2[j] 结尾的最长公共子串长度
func LCS(str1 string, str2 string) (ans string) {
	n, m := len(str1), len(str2)
	dp := [5001][5001]int{}
	maxlen, end := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if str1[i] == str2[j] {
				if i-1 >= 0 && j-1 >= 0 {
					dp[i][j] = dp[i-1][j-1] + 1
				} else {
					dp[i][j] = 1
				}
				// fmt.Printf("dp[i][j]: %v\n", dp[i][j])
			}
			if dp[i][j] > maxlen {
				maxlen = dp[i][j]
				end = i
			}
		}
	}

	ans = str1[end-maxlen+1 : end+1]
	return
}

// 暴力解法
func LCS2(str1 string, str2 string) (ans string) {
	n, m := len(str1), len(str2)
	// 取 str1[i] 固定
	for i := 0; i < n; i++ {
		// 遍历 str2, 找到与 str[i] 相同的字符作为起点进行搜索
		for j := 0; j < m; j++ {
			if str1[i] != str2[j] {
				continue
			}

			// 搜索得到 str1 和 str2 的相同长度的字符串
			length := 0
			for ; i+length < n && j+length < m && str1[i+length] == str2[j+length]; length++ {
			}
			if length > len(ans) {
				ans = str1[i : i+length]
			}
		}
	}
	return
}
