package nc0079

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0079_01(t *testing.T) {
	ans := rob([]int{1, 2, 3, 4})
	assert.Equal(t, 6, ans)
}

func Test_NC0079_02(t *testing.T) {
	ans := rob([]int{1, 3, 6})
	assert.Equal(t, 6, ans)
}

func Test_NC0079_03(t *testing.T) {
	ans := rob([]int{69, 27, 25, 44, 1, 16, 76, 98, 22, 52})
	assert.Equal(t, 237, ans)
}
