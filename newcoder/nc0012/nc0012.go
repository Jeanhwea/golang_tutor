package nc0012

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 单链表的排序
func sortInList(head *ListNode) (ans *ListNode) {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	ans = merge(head, n)
	return
}

func merge(head *ListNode, n int) (out *ListNode) {
	if n <= 1 {
		return head
	}

	left, right := head, head
	for i := 1; i < n/2; i++ {
		right = right.Next
	}
	right.Next, right = nil, right.Next

	out = mergeTwo(merge(left, n/2), merge(right, n-n/2))
	return
}

func mergeTwo(head1, head2 *ListNode) (ans *ListNode) {
	dummy := &ListNode{}
	tail, p, q := dummy, head1, head2
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next, tail, p = p, p, p.Next
		} else {
			tail.Next, tail, q = q, q, q.Next
		}
	}
	if p == nil {
		tail.Next = q
	} else {
		tail.Next = p
	}
	ans = dummy.Next
	return
}
