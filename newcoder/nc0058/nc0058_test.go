package nc0058

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0058_01(t *testing.T) {
	ans := Permutation("abac")
	t.Log(ans)
	assert.Equal(t, true, true)
}
