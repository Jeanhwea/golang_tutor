package nc0014

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 链表的奇偶重排
func oddEvenList(head *ListNode) (ans *ListNode) {
	odd, even := &ListNode{}, &ListNode{}
	oddTail, evenTail := odd, even
	for i := 1; head != nil; i++ {
		if i%2 == 1 {
			oddTail.Next, oddTail, head.Next, head = head, head, nil, head.Next
		} else {
			evenTail.Next, evenTail, head.Next, head = head, head, nil, head.Next
		}
	}
	oddTail.Next = even.Next
	ans = odd.Next
	return
}
