package nc0084

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0084_01(t *testing.T) {
	strs := []string{"abca", "abc", "abca", "abc", "abcc"}
	ans := longestCommonPrefix(strs)
	assert.Equal(t, "abc", ans)
}
