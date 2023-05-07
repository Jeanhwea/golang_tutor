package nc0011

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0011_01(t *testing.T) {
	a := NewListNode([]int{2, 3})
	b := NewListNode([]int{8, 8})
	ans := addInList(a, b)
	assert.Equal(t, []int{1, 1, 1}, ans.ToSlice())
}
