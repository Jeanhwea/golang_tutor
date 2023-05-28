package nc0096

import (
	"sort"
)

// 主持人调度（二）
func minmumNumberOfHost(n int, startEnd [][]int) (ans int) {
	beg, end := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		beg[i], end[i] = startEnd[i][0], startEnd[i][1]
	}
	sort.Ints(beg)
	sort.Ints(end)

	j := 0
	for i := 0; i < n; i++ {
		if beg[i] >= end[j] {
			j++
		} else {
			ans++
		}
	}

	return
}
