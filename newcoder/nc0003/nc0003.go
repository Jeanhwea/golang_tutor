package nc0003

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// K 个链表一组翻转
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}

	tail := dummy
	for curr := head; curr != nil; {
		var subHead, subTail *ListNode
		i := 0
		for ; i < k && curr != nil; i++ {
			if i == 0 {
				subHead, subTail, curr.Next, curr = curr, curr, nil, curr.Next
			} else {
				curr.Next, subHead, curr = subHead, curr, curr.Next
			}
		}

		if i == k {
			tail.Next, tail = subHead, subTail
		} else {
			var newHead *ListNode
			for curr := subHead; curr != nil; {
				if newHead == nil {
					newHead, curr.Next, curr = curr, nil, curr.Next
				} else {
					curr.Next, newHead, curr = newHead, curr, curr.Next
				}
			}
			tail.Next = newHead
		}
	}

	return dummy.Next
}
