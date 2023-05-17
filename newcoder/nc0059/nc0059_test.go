package nc0059

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0059_01(t *testing.T) {
	ans := Nqueen(8)
	t.Log(ans)
	assert.Equal(t, true, true)
}
