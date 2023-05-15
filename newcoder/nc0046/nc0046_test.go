package nc0046

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NC0046_01(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestNc004601(t *testing.T) {
	nums := []int{3, 4, 2, 1}
	var heap IntHeap
	for _, num := range nums {
		heap.push(num)
		t.Log(heap)
	}
	for i := 0; i < 3; i++ {
		t.Log(heap.pop())
		t.Log(heap)
	}
}
