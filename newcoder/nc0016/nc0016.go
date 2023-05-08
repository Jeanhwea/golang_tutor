package nc0016

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 删除有序链表中重复的元素-II
func deleteDuplicates(head *ListNode) (ans *ListNode) {
	dummy := &ListNode{}
	tail := dummy
	for head != nil {
		// head 是最后一个元素, head 和后续的元素不相等, 这两种情况都需要添加到结果链表中
		if head.Next == nil || head.Val != head.Next.Val {
			tail.Next, tail, head = head, head, head.Next
		} else { // 如果发现连续 2 个重复的则跳过这一批数字
			val := head.Val
			for ; head != nil; head = head.Next {
				if head.Val != val {
					break
				}
			}
		}
	}
	tail.Next = nil

	ans = dummy.Next
	return
}
