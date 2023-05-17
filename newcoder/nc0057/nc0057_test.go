package nc0057

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0057_01(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'0', '1', '0', '1', '1'},
		{'0', '0', '0', '1', '1'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '1', '1', '1'},
	}
	ans := solve(grid)
	assert.Equal(t, 3, ans)

}
