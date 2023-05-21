package nc0073

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0073_01(t *testing.T) {
	s := "baabccc"
	ans := getLongestPalindrome(s)
	assert.Equal(t, 4, ans)
}
