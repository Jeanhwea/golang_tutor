package nc0093

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0093_01(t *testing.T) {
	height := []int{1, 7, 3, 2, 4, 5, 8, 2, 7}
	ans := maxArea(height)
	assert.Equal(t, 49, ans)
}
