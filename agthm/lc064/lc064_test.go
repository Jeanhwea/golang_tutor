package lc064

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LC064_01(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	ans := minPathSum(grid)
	assert.Equal(t, ans, 7)
}
