package nc0016

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0016_01(t *testing.T) {
	head := NewListNode([]int{1, 1, 2, 3, 4})
	ans := deleteDuplicates(head)
	assert.Equal(t, []int{2, 3, 4}, ans.ToSlice())
}
