package nc0083

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0083_01(t *testing.T) {
	s := "This is a sample"
	ans := trans(s, len(s))
	assert.Equal(t, "SAMPLE A IS tHIS", ans)
}

func Test_NC0083_02(t *testing.T) {
	s := "h i "
	ans := trans(s, len(s))
	assert.Equal(t, " I H", ans)
}
