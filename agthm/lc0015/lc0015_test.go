package lc0015

import (
	"testing"
)

func Test_LC0015_01(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ans := threeSum(nums)
	t.Log(ans)
}

func Test_LC0015_02(t *testing.T) {
	nums := []int{0, 0, 0}
	ans := threeSum(nums)
	t.Log(ans)
}

func Test_LC0015_03(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	ans := threeSum(nums)
	t.Log(ans)
}
