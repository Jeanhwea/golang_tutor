package nc0049

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0049_01(t *testing.T) {
	assert.Equal(t, 3, solve("1+2"))
}

func Test_NC0049_02(t *testing.T) {
	assert.Equal(t, 13, solve("11+2"))
}

func Test_NC0049_03(t *testing.T) {
	assert.Equal(t, 9, solve("11-2"))
}

func Test_NC0049_04(t *testing.T) {
	assert.Equal(t, 7, solve("11-2*2"))
}

func Test_NC0049_05(t *testing.T) {
	assert.Equal(t, 26, solve("3+2*3*4-1"))
}

func Test_NC0049_06(t *testing.T) {
	assert.Equal(t, -10, solve("(2*(3-4))*5"))
}

func Test_NC0049_07(t *testing.T) {
	assert.Equal(t, -10, solve("-10"))
}
