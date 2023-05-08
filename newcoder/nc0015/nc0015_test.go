package nc0015

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0015_01(t *testing.T) {
	head := NewListNode([]int{1, 1, 2, 3})
	ans := deleteDuplicates(head)
	assert.Equal(t, []int{1, 2, 3}, ans.ToSlice())
}

func Test_NC0015_02(t *testing.T) {
	head := NewListNode([]int{1, 1, 1, 1})
	ans := deleteDuplicates(head)
	assert.Equal(t, []int{1}, ans.ToSlice())
}
