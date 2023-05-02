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

func Test_LC021_12(t *testing.T) {
	ans := makeList(nil)
	t.Log(sliceList(ans))
}

func Test_LC021_01(t *testing.T) {
	l1 := makeList([]int{1, 2, 4})
	l2 := makeList([]int{1, 3, 4})
	ans := mergeTwoLists(l1, l2)
	t.Log(sliceList(ans))
}
