package nc0061

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0061_01(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	ans := solve(matrix)
	assert.Equal(t, 5, ans)
}
