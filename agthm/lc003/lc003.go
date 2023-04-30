package lc003

// 无重复字符的最长子串
//
//   1. 使用双指针解法
//   2. 用一个 map 保持当前字符串的所有字符, 用于提高查找效率
//
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
