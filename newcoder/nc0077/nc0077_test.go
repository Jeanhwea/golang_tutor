package nc0077

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0077_01(t *testing.T) {
	ans := longestValidParentheses2("(()")
	assert.Equal(t, 2, ans)
}

func Test_NC0077_02(t *testing.T) {
	ans := longestValidParentheses2(")")
	assert.Equal(t, 0, ans)
}

func Test_NC0077_03(t *testing.T) {
	ans := longestValidParentheses("()")
	assert.Equal(t, 2, ans)
}

func Test_NC0077_04(t *testing.T) {
	ans := longestValidParentheses("())")
	assert.Equal(t, 2, ans)
}

func Test_NC0077_05(t *testing.T) {
	ans := longestValidParentheses("()(())")
	assert.Equal(t, 6, ans)
}
