package code0512a

import (
	"strconv"
	"strings"
)

func restoreIpAddress(s string) []string {
	ans = []string{}
	bfs([]string{}, s, 0)
	return ans
}

var (
	ans []string
)

func bfs(curr []string, left string, level int) {
	if len(left) <= 0 && level == 4 {
		ans = append(ans, strings.Join(curr, "."))
		return
	}

	for i := 1; i <= 3 && i <= len(left); i++ {
		num := left[:i]
		if isIpPart(num) {
			newCurr := append(curr, num)
			bfs(newCurr, left[i:], level+1)
		}
	}
}

func isIpPart(num string) bool {
	if len(num) > 1 && strings.HasPrefix(num, "0") {
		return false
	}
	n, err := strconv.Atoi(num)
	if err != nil {
		return false
	}
	if n >= 0 && n <= 255 {
		return true
	}
	return false
}
