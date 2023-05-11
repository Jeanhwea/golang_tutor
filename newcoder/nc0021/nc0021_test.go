package nc0021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0021_01(t *testing.T) {
	nums := []int{4, 3}
	ans := minNumberInRotateArray(nums)
	assert.Equal(t, 3, ans)
}
