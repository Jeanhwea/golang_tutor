package nc0051

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0051_01(t *testing.T) {
	nums := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	assert.Equal(t, 2, MoreThanHalfNum_Solution(nums))
}
