package nc0086

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0086_01(t *testing.T) {
	a, b := "10", "21"
	assert.Equal(t, "31", solve(a,b))
}
