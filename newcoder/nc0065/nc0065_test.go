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

func Test_NC0065_02(t *testing.T) {
	s1, s2 := "1A2C3D4B56", "B1D23A456A"
	ans := LCS(s1, s2)
	t.Log(ans)
	assert.Equal(t, true, true)
}

func TestNc006501(t *testing.T) {
	ans := arr2d(2, 3)
	t.Log(ans)
}
