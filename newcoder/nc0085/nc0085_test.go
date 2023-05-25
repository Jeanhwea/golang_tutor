package nc0085

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0085_01(t *testing.T) {
	ans := solve("172.16.254.1")
	assert.Equal(t, "IPv4", ans)
}

func Test_NC0085_02(t *testing.T) {
	ans := solve("2001:0db8:85a3:0:0:8A2E:0370:7334")
	assert.Equal(t, "IPv6", ans)
}

func Test_NC0085_03(t *testing.T) {
	ans := solve("256.256.256.256")
	assert.Equal(t, "Neither", ans)
}
