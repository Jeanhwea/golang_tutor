package lc007

import (
	"math"
	"testing"
)

func Test_LC007_01(t *testing.T) {
	for _, x := range []int{123, -34, -2147483648, 2147483647} {
		ans := reverse(x)
		t.Logf("%v -> %v\n", x, ans)
	}
}

func Test_LC007_11(t *testing.T) {
	t.Log(math.MinInt32)
	t.Log(math.MaxInt32)
	t.Log(-8%10)
}
