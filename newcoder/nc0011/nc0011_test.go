package nc0011

import (
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Test_NC0011_01(t *testing.T) {
	a := NewListNode([]int{2, 3})
	b := NewListNode([]int{8, 8})
	t.Log(a)
	t.Log(b)
	ans := addInList(a, b)
	t.Log(ans)
}
