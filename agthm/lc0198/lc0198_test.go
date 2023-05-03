package lc0198

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LC0198_01(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	ans := rob(nums)
	assert.Equal(t, 4, ans)
}

func Test_LC0198_02(t *testing.T) {
	nums := []int{2, 1, 1, 2}
	ans := rob(nums)
	assert.Equal(t, 4, ans)
}
