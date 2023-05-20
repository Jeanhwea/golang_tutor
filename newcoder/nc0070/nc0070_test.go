package nc0070

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0070_01(t *testing.T) {
	ans := minMoney([]int{5, 2, 3}, 22)
	assert.Equal(t, 5, ans)
}
