package nc0060

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0060_01(t *testing.T) {
	ans := generateParenthesis(4)
	t.Log(ans)
	assert.Equal(t, true, true)
}
