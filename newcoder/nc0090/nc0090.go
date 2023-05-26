package nc0090

import "fmt"

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
	for {
		// 窗口扩张
		for j < n && !covered() {
			if _, ok := needCount[s[j]]; ok {
				needCount[s[j]]--
			}
			j++
		}

		if len(ans) == 0 || len(ans) > j-i+1 {
			ans = string(s[i : j+1])
			fmt.Printf("ans: %v\n", ans)
		}

		if j >= n {
			break
		}

		// 窗口收缩
		if _, ok := needCount[s[i]]; ok {
			needCount[s[i]]++
		}
		i++
	}

	return
}
