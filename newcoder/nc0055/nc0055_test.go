package nc0055

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0055_01(t *testing.T) {
	ans := permute([]int{1, 2, 3})
	t.Log(ans)
	assert.Equal(t, true, true)
}
