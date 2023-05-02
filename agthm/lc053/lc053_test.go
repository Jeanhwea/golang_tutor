package lc053

import (
	"testing"
)

func Test_LC053_01(t *testing.T) {
	nums := []int{1}
	ans := maxSubArray(nums)
	t.Log(ans)
}

func Test_LC053_02(t *testing.T) {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	ans := maxSubArray(nums)
	t.Log(ans)
}

func Test_LC053_03(t *testing.T) {
	nums := []int{-1}
	ans := maxSubArray(nums)
	t.Log(ans)
}
