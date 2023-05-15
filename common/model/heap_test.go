package model

import (
	"container/heap"
	"testing"
)

func TestHeap01(t *testing.T) {
	h := &MinHeap{}
	heap.Init(h)
	nums := []int{5, 2, 3, 4, 1, 6, 7, 0, 8}
	for _, v := range nums {
		heap.Push(h, v)
		t.Log(h)
	}

	for h.Len() > 0 {
		t.Log(heap.Pop(h))
	}
}
