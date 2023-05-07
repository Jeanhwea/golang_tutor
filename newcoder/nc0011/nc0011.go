package nc0011

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 链表相加(二)
func addInList(head1 *ListNode, head2 *ListNode) (ans *ListNode) {
	p, q := reverse(head1), reverse(head2)
	carry, sum := 0, 0
	for p != nil || q != nil {
		sum = carry
		if p != nil {
			sum += p.Val
			p = p.Next
		}
		if q != nil {
			sum += q.Val
			q = q.Next
		}
		carry = sum / 10
		curr := &ListNode{Val: sum % 10}
		curr.Next, ans = ans, curr
	}

	if carry > 0 {
		curr := &ListNode{Val: carry}
		curr.Next, ans = ans, curr
	}

	return
}

func reverse(head *ListNode) (ans *ListNode) {
	for head != nil {
		head.Next, ans, head = ans, head, head.Next
	}
	return
}
