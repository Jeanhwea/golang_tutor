package nc0045

import (
	"container/heap"
	"testing"

	. "github.com/jeanhwea/golang_tutor/common/model"
	"github.com/stretchr/testify/assert"
)

func Test_NC0045_01(t *testing.T) {
	num := []int{2, 3, 4, 2, 6, 2, 5, 1}
	assert.Equal(t, []int{4, 4, 6, 6, 6, 5}, maxInWindows(num, 3))
}

func TestNc004501(t *testing.T) {
	h := new(MinHeap)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 4)
	t.Logf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		t.Logf("%d ", heap.Pop(h))
	}
}
