package nc0013

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 是否是回文链表
func isPail(head *ListNode) (ans bool) {
	var reverse *ListNode
	for p := head; p != nil; p = p.Next {
		curr := &ListNode{Val: p.Val}
		reverse, curr.Next = curr, p
	}

	for p, q := head, reverse; p != nil && q != nil; p, q = p.Next, q.Next {
		if p.Val != q.Val {
			ans = false
			return
		}
	}
	ans = true
	return
}
