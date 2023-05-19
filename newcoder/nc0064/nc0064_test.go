package nc0064

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0064_01(t *testing.T) {
	cost := []int{2, 5, 20}
	assert.Equal(t, 5, minCostClimbingStairs(cost))
}
