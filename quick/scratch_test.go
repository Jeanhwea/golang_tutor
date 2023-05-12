package main

import (
	"encoding/json"
	"sort"
	"testing"
)

func TestScratch01(t *testing.T) {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			t.Log(a)
		}
		a[k] = 100 + v
	}
	t.Log(a)
}

func TestScratch02(t *testing.T) {
	a := []int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			t.Log(a)
		}
		a[k] = 100 + v
	}
	t.Log(a)
}

type Interval struct {
	Start int
	End   int
}

func TestScratch03(t *testing.T) {
	intervals := []*Interval{
		{Start: 3, End: 4},
		{Start: 13, End: 4},
		{Start: 3, End: 21},
		{Start: 3, End: 11},
		{Start: 1, End: 2},
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].Start < intervals[j].Start {
			return true
		} else if intervals[i].Start == intervals[j].Start {
			return intervals[i].End < intervals[j].End
		} else {
			return false
		}
	})
	out, _ := json.Marshal(intervals)
	t.Log(string(out))
}
