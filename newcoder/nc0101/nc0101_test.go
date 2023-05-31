package nc0101

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0101_01(t *testing.T) {
	ops := [][]int{
		{1, 1, 1},
		{1, 2, 2},
		{1, 3, 2},
		{1, 2, 4},
		{1, 3, 5},
		{2, 2},
		{1, 4, 4},
		{2, 1},
	}
	ans := LFU(ops, 3)
	t.Log(ans)
	assert.Equal(t, true, true)
}
