package lc003

func lengthOfLongestSubstring(s string) (ans int) {
	n := len(s)
	m := make(map[byte]bool, 128) // 保存 s[i:j-1] 的字符

	for i, j := 0, 0; i < n; i++ {
		for ; j < n; j++ {
			// fmt.Printf("i=%v, j=%v, substr=%v\n", i, j, s[i:j])
			if ok := m[s[j]]; ok {
				ans = max(ans, j-i)
				break
			} else {
				m[s[j]] = true
			}
		}

		if j == n {
			ans = max(ans, j-i)
			return
		}

		// i 将会增加, 这个先置空
		m[s[i]] = false
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
