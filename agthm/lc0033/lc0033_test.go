package lc0033

import (
	"testing"
)

func Test_LC0033_01(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	ans := search(nums, 4)
	t.Log(ans)
}

func Test_LC0033_02(t *testing.T) {
	nums := []int{3, 1}
	ans := search(nums, 1)
	t.Log(ans)
}
