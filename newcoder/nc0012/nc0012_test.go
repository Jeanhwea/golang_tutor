package nc0012

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0012_01(t *testing.T) {
	head := NewListNode([]int{9, 3, 4, 5})
	ans := sortInList(head)
	assert.Equal(t, []int{3, 4, 5, 9}, ans.ToSlice())
}

func TestNc001202(t *testing.T) {
	a := NewListNode([]int{9, 3, 4, 5})
	b := NewListNode([]int{9, 3, 4, 5})
	test := mergeTwo(a, b)
	t.Log(test)
}
