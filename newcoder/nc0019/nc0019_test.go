package nc0019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0019_01(t *testing.T) {
	nums := []int{1, 2, 3, 4, 3}
	ans := findPeakElement(nums)
	assert.Equal(t, 3, ans)
}
