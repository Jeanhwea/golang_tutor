package main

import (
	"testing"
)

func f(a []int) {
	a = append(a, 2)
}

func f1(a []int) {
	a[0] = 2
}

func f2(a []int) {
	a = append(a, 2)
	a[0] = 2
}

func TestTutor10Slice01(t *testing.T) {
	// a := []int{1}
	a := make([]int, 0, 4)
	t.Logf("a=%v, len=%v, cap=%v\n", a, len(a), cap(a))
	f2(a)
	t.Logf("a=%v, len=%v, cap=%v\n", a, len(a), cap(a))
	a = append(a, 2)
	t.Logf("a=%v, len=%v, cap=%v\n", a, len(a), cap(a))

	t.Logf("a=%p, len=%v, cap=%v\n", &a, len(a), cap(a))
}
