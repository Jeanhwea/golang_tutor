package nc0008

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 链表中倒数最后k个结点
func FindKthToTail(head *ListNode, k int) (ans *ListNode) {
	fast, slow := head, head
	for i := 0; i < k; i++ {
		if fast == nil {
			return
		}
		fast = fast.Next
	}

	for fast != nil {
		fast, slow = fast.Next, slow.Next
	}

	ans = slow
	return
}
