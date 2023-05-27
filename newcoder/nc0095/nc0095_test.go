package nc0095

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0095_01(t *testing.T) {
	scores := []int{1, 1, 2}
	ans := candy(scores)
	assert.Equal(t, 4, ans)
}

func Test_NC0095_02(t *testing.T) {
	scores := []int{2}
	ans := candy(scores)
	assert.Equal(t, 1, ans)
}

func Test_NC0095_03(t *testing.T) {
	scores := []int{10, 4, 10, 10, 4}
	ans := candy(scores)
	assert.Equal(t, 8, ans)
}
