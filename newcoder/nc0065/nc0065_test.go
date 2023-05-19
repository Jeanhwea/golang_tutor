package nc0065

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0065_01(t *testing.T) {
	s1, s2 := "1234", "123"
	ans := LCS(s1, s2)
	t.Log(ans)
	assert.Equal(t, true, true)
}
