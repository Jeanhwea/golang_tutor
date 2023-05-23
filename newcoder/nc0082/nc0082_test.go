package nc0082

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0082_01(t *testing.T) {
	ans := maxProfit([]int{8, 9, 3, 5, 1, 3})
	assert.Equal(t, 4, ans)
}

func Test_NC0082_02(t *testing.T) {
	ans := maxProfit([]int{1, 83, 74, 26, 63, 37, 25, 63, 28})
	assert.Equal(t, 120, ans)
}
