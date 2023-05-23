package nc0081

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0081_01(t *testing.T) {
	ans := maxProfit([]int{8, 9, 2, 5, 4, 7, 1})
	assert.Equal(t, 7, ans)
}
