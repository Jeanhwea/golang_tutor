package nc0044

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0044_01(t *testing.T) {
	assert.Equal(t, true, isValid(""))
	assert.Equal(t, true, isValid("()"))
	assert.Equal(t, false, isValid("(]"))
	assert.Equal(t, false, isValid("]"))
}
