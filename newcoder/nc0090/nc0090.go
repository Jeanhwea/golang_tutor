package nc0090

// 最小覆盖字串
func minWindow(s string, t string) (ans string) {
	n, m, minLen := len(s), len(t), len(s)+1
	needCount := make(map[byte]int) // 维护需求的 map
	for i := 0; i < m; i++ {
		needCount[t[i]]++
	}

	// 判断 t 是否已经被全覆盖
	covered := func() bool {
		for _, count := range needCount {
			if count > 0 {
				return false
			}
		}
		return true
	}

	// 使用双指针来遍历 s, 期间维护一个最小覆盖串
	for i, j := 0, 0; j < n; {
		if _, ok := needCount[s[j]]; ok {
			needCount[s[j]]--
		}

		for i <= j && covered() {
			if minLen > j-i+1 {
				ans, minLen = string(s[i:j+1]), j-i+1
			}
			if _, ok := needCount[s[i]]; ok {
				needCount[s[i]]++
			}
			i++
		}

		j++
	}

	if minLen > n {
		ans = ""
	}
	return
}
