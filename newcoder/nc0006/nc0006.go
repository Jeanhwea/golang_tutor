package nc0006

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 判断链表中是否有环
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil && slow != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}
