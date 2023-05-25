package nc0089

import (
	"sort"
)

type Interval struct {
	Start int
	End   int
}

// 合并区间
func merge(intervals []*Interval) (ans []*Interval) {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		} else if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		}
		return false
	})

	n := len(intervals)
	var last *Interval
	for i := 0; i < n; i++ {
		curr := intervals[i]
		if last == nil {
			last = curr
			continue
		}
		if curr.Start > last.End {
			ans = append(ans, last)
			last = curr
		} else {
			if curr.End > last.End {
				last.End = curr.End
			}
		}
	}
	if last != nil {
		ans = append(ans, last)
	}
	return
}
