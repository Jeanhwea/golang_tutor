package nc0020

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0020_01(t *testing.T) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 0}
	ans := InversePairs(data)
	assert.Equal(t, 7, ans)
}

func TestNc002002(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	temp := data[1:4]
	temp[0] = 1
	t.Log(data)
}
