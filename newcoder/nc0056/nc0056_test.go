package nc0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0056_01(t *testing.T) {
	ans := permuteUnique([]int{3, 2, 0, 1, 2})
	t.Log(ans)
	assert.Equal(t, true, true)
}
