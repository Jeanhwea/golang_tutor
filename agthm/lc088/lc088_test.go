package lc088

import (
	"testing"
)

func Test_LC088_01(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m, n := 3, 3
	nums2 := []int{2, 5, 6}
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}
