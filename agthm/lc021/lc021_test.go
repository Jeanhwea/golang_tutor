package lc021

import (
	"testing"
)

func makeList(vals []int) (head *ListNode) {
	dummy := &ListNode{}
	tail := dummy
	for _, val := range vals {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}
	head = dummy.Next
	return
}

func sliceList(head *ListNode) (vals []int) {
	curr := head
	for curr != nil {
		vals = append(vals, curr.Val)
		curr = curr.Next
	}
	return
}

func Test_LC021_11(t *testing.T) {
	ans := makeList([]int{1, 2, 3})
	t.Log(sliceList(ans))
}
