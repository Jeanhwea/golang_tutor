package nc0071

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0071_01(t *testing.T) {
	nums := []int{6, 3, 1, 5, 2, 3, 7}
	ans := LIS(nums)
	assert.Equal(t, 4, ans)
}

func Test_NC0071_02(t *testing.T) {
	nums := []int{3, 5, 7, 1, 2, 4, 6, 3, 8, 9, 5, 6}
	ans := LIS(nums)
	assert.Equal(t, 6, ans)
}
