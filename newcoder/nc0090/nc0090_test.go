package nc0090

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0090_01(t *testing.T) {
	ans := minWindow("XDOYEZODEYXNZ", "XYZ")
	assert.Equal(t, "YXNZ", ans)
}

func TestNc009001(t *testing.T) {
	m := make(map[byte]int)
	m['a']++
	t.Log(m)
}

func TestNc009002(t *testing.T) {
	s := "abc"
	t.Log(s[0:1])
}
