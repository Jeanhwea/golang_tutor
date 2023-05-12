package nc0001

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// BM1 翻转链表
func ReverseList(head *ListNode) (ans *ListNode) {
	dummy := &ListNode{}
	tail := dummy.Next
	for head != nil {
		dummy.Next = head
		head = head.Next
		dummy.Next.Next = tail
		tail = dummy.Next
	}
	return dummy.Next
}
