package nc0052

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0052_01(t *testing.T) {
	nums := []int{1, 1, 32, 16}
	ans := FindNumsAppearOnce(nums)
	t.Log(ans)
	assert.Equal(t, true, true)
}
