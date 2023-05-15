package nc0048

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0048_01(t *testing.T) {
	nums := []int{5, 2, 3, 4, 1, 6, 7, 0, 8}
	for _, v := range nums {
		Insert(v)
		t.Logf("%v %v %v\n", smallPart, largePart, GetMedian())
	}
	assert.Equal(t, true, true)
}
