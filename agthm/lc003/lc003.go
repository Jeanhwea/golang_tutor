package lc003

func lengthOfLongestSubstring(s string) (ans int) {
	n := len(s)
	m := make(map[byte]bool, 128) // 保存 s[i:j-1] 的字符

	for i, j := 0, 0; i < n; i++ {
		for ; j < n; j++ {
			// fmt.Printf("i=%v, j=%v, substr=%v\n", i, j, s[i:j])
			if ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = true
		}

		if j-i > ans {
			ans = j - i
		}
		if j == n {
			return
		}

		// i 将会增加, 这个先置空
		m[s[i]] = false
	}

	return
}
