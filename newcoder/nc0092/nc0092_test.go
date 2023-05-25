package nc0092

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0092_01(t *testing.T) {
	ans := maxLength([]int{2, 2, 3, 4, 3})
	assert.Equal(t, 3, ans)
}

func Test_NC0092_02(t *testing.T) {
	ans := maxLength([]int{3, 3, 3, 3, 3, 3})
	assert.Equal(t, 1, ans)
}
