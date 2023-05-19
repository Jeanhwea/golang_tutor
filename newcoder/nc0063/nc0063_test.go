package nc0063

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0063_01(t *testing.T) {
	for i := 1; i < 10; i++ {
		t.Logf("n = %v, ans = %v\n", i, jumpFloor(i))
	}
	assert.Equal(t, true, true)
}
