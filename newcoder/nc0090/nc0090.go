package nc0090

// 最小覆盖字串
func minWindow(s string, t string) (ans string) {
	n, m := len(s), len(t)
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

	i, j := 0, 0
	for j < n {
		// 窗口扩张
		for j < n && !covered() {
			if _, ok := needCount[s[j]]; ok {
				needCount[s[j]]--
			}
			j++
		}

		// 窗口收缩
		for i < j && covered() {
			if len(ans) == 0 || len(ans) > j-i {
				ans = string(s[i:j])
			}

			if _, ok := needCount[s[i]]; ok {
				needCount[s[i]]++
			}
			i++
		}
	}

	return
}
