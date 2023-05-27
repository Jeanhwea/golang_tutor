package nc0094

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0094_01(t *testing.T) {
	arr := []int{3, 1, 2, 5, 2, 4}
	ans := maxWater2(arr)
	assert.Equal(t, int64(5), ans)
}
