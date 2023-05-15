package nc0047

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0047_01(t *testing.T) {
	nums := []int{10, 10, 9, 9, 8, 7, 5, 6, 4, 3, 4, 2}
	assert.Equal(t, 9, findKth(nums, len(nums), 3))
}
