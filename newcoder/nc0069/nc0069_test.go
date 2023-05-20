package nc0069

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0069_01(t *testing.T) {
	ans := solve("12")
	assert.Equal(t, 2, ans)
}

func Test_NC0069_02(t *testing.T) {
	ans := solve("31717126241541717")
	assert.Equal(t, 192, ans)
}

func Test_NC0069_03(t *testing.T) {
	ans := solve("0")
	assert.Equal(t, 0, ans)
}
