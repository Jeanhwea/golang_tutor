package nc0076

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0076_01(t *testing.T) {
	ans := match("aaab", "a*b")
	assert.Equal(t, true, ans)
}

func Test_NC0076_02(t *testing.T) {
	ans := match("ab", ".b")
	assert.Equal(t, true, ans)
}

func Test_NC0076_03(t *testing.T) {
	ans := match("ab", "ab")
	assert.Equal(t, true, ans)
}

func TestNc007601(t *testing.T) {
	ans := match("aaa", "a*a")
	assert.Equal(t, true, ans)
}

func TestNc007602(t *testing.T) {
	ans := match("aad", "c*a*d")
	assert.Equal(t, true, ans)
}

func TestNc007603(t *testing.T) {
	ans := match("a", ".*")
	assert.Equal(t, true, ans)
}

func TestNc007604(t *testing.T) {
	ans := match("ab", ".*..")
	assert.Equal(t, true, ans)
}
