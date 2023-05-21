package nc0072

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0072_01(t *testing.T) {
	nums := []int{1, -2, 3, 10, -4, 7, 2, -5}
	ans := FindGreatestSumOfSubArray(nums)
	assert.Equal(t, 18, ans)
}

func Test_NC0072_02(t *testing.T) {
	nums := []int{-5}
	ans := FindGreatestSumOfSubArray(nums)
	assert.Equal(t, -5, ans)
}
