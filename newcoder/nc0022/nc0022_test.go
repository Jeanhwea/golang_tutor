package nc0022

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0022_01(t *testing.T) {
	ans := compare("1.2", "1.2.1")
	assert.Equal(t, -1, ans)
}

func Test_NC0022_02(t *testing.T) {
	ans := compare("1.1", "1.01")
	assert.Equal(t, 0, ans)
}
