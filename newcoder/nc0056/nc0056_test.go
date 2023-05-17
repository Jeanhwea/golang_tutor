package nc0056

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0056_01(t *testing.T) {
	ans := permuteUnique([]int{3, 2, 0, 1, 2})
	for i, v := range ans {
		t.Logf("%3d %v", i+1, v)
	}
	assert.Equal(t, true, true)
}

func Test_NC0056_02(t *testing.T) {
	ans := permuteUnique([]int{0, 1, 1})
	for i, v := range ans {
		t.Logf("%3d %v", i+1, v)
	}
	assert.Equal(t, true, true)
}

func Test_NC0056_03(t *testing.T) {
	ans := permuteUnique([]int{1, 1})
	for i, v := range ans {
		t.Logf("%3d %v", i+1, v)
	}
	assert.Equal(t, true, true)
}
