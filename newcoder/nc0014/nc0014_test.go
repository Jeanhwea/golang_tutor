package nc0014

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0014_01(t *testing.T) {
	head := NewListNode([]int{1, 2, 3, 4, 5, 6})
	ans := oddEvenList(head)
	t.Log(ans)
	assert.Equal(t, []int{1, 3, 5, 2, 4, 6}, ans.ToSlice())
}
