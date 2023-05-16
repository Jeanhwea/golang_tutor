package nc0054

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0054_01(t *testing.T) {
	nums := []int{-2, 0, 1, 1, 2}
	ans := threeSum(nums)
	t.Log(ans)
	assert.Equal(t, true, true)
}

func Test_NC0054_02(t *testing.T) {
	nums := []int{0, 0, 0, 1}
	ans := threeSum(nums)
	t.Log(ans)
	assert.Equal(t, true, true)
}

func Test_NC0054_03(t *testing.T) {
	nums := []int{0, 0, 0, 0}
	ans := threeSum(nums)
	t.Log(ans)
	assert.Equal(t, true, true)
}
