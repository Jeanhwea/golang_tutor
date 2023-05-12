package nc0015

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 删除有序链表中重复的元素-I
func deleteDuplicates(head *ListNode) (ans *ListNode) {
	dummy := &ListNode{Next: head}
	var prev *ListNode
	tail := dummy
	for tail != nil {
		if prev != nil && prev.Val == tail.Val {
			prev.Next, tail = tail.Next, tail.Next
		} else {
			prev, tail = tail, tail.Next
		}
	}
	ans = dummy.Next
	return
}
