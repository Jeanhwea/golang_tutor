package nc0045

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0045_01(t *testing.T) {
	num := []int{2, 3, 4, 2, 6, 2, 5, 1}
	assert.Equal(t, []int{4, 4, 6, 6, 6, 5}, maxInWindows(num, 3))
}
