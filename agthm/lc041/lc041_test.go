package lc041

import (
	"testing"
)

func Test_LC041_01(t *testing.T) {
	nums := []int{1, 2, 0}
	t.Log(firstMissingPositive(nums))
}

func Test_LC041_02(t *testing.T) {
	nums := []int{1}
	t.Log(firstMissingPositive(nums))
}
