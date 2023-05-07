package nc0009

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 删除链表的倒数第 n 个节点
func removeNthFromEnd(head *ListNode, n int) (ans *ListNode) {
	// 先用快慢指针找到倒数节点
	head = &ListNode{Next: head}
	fast, slow := head, head
	for i := 0; i <= n; i++ {
		if fast == nil {
			ans = head
			return
		}
		fast = fast.Next
	}
	for fast != nil {
		fast, slow = fast.Next, slow.Next
	}

	// 然后直接删除倒数 n 个节点
	if slow != nil && slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	ans = head.Next
	return
}
