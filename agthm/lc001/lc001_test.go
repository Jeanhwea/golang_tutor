package lc001

import (
	"testing"
)

func Test_LC001_01(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	ans := twoSum(nums, target)
	t.Log(ans)
}
