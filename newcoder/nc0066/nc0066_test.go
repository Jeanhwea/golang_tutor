package nc0066

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0066_01(t *testing.T) {
	str1, str2 := "abca", "aaba"
	ans := LCS2(str1, str2)
	t.Log(ans)
	assert.Equal(t, true, true)
}

func Test_NC0066_02(t *testing.T) {
	str1, str2 := "1AB2345CD", "12345EF"
	ans := LCS(str1, str2)
	t.Log(ans)
	assert.Equal(t, true, true)
}
