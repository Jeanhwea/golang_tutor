package nc0007

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 链表中环的入口结点
func EntryNodeOfLoop(head *ListNode) (ans *ListNode) {
	fast, slow := head, head

	// 先找到快慢指针相遇的点
	var meet *ListNode
	for fast != nil && fast.Next != nil && slow != nil {
		fast, slow = fast.Next.Next, slow.Next
		if fast == slow {
			meet = fast
			break
		}
	}

	// 通过计算距离可以得知: head 出发 和 相遇点 出发汇合的距离就是环的入口
	for head != nil && meet != nil {
		if head == meet {
			ans = meet
			return
		}
		head, meet = head.Next, meet.Next
	}

	return
}
