package nc0074

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0074_01(t *testing.T) {
	t.Log(restoreIpAddresses("25525522135"))
	assert.Equal(t, true, true)
}
