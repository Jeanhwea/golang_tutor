package nc0096

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0096_01(t *testing.T) {
	startEnd := [][]int{
		{1, 2},
		{2, 3},
	}
	ans := minmumNumberOfHost(len(startEnd), startEnd)
	assert.Equal(t, 1, ans)
}

func Test_NC0096_02(t *testing.T) {
	startEnd := [][]int{
		{1, 3},
		{2, 4},
	}
	ans := minmumNumberOfHost(len(startEnd), startEnd)
	assert.Equal(t, 2, ans)
}

func Test_NC0096_03(t *testing.T) {
	startEnd := [][]int{
		{1, 5},
		{2, 6},
		{3, 4},
	}
	ans := minmumNumberOfHost(len(startEnd), startEnd)
	assert.Equal(t, 3, ans)
}
