package lc0088

import (
	"testing"
)

func Test_LC0088_01(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m, n := 3, 3
	nums2 := []int{2, 5, 6}
	merge(nums1, m, nums2, n)
	t.Log(nums1)
}

func Benchmark_LC088_01(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums1 := []int{1, 2, 3, 0, 0, 0}
		m, n := 3, 3
		nums2 := []int{2, 5, 6}
		merge(nums1, m, nums2, n)
	}
}
