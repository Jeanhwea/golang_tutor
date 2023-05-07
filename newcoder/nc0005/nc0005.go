package nc0005

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 合并 K 哥有序链表
func mergeKLists(lists []*ListNode) (ans *ListNode) {
	return doMerge(lists, 0, len(lists)-1)
}

// 归并排序, 对数组下标进行归并
func doMerge(lists []*ListNode, left, right int) (out *ListNode) {
	if left > right {
		return
	} else if left == right {
		return lists[left]
	}

	// left < right
	mid := left + (right-left)/2
	out = merge(doMerge(lists, left, mid), doMerge(lists, mid+1, right))
	return
}

// 合并两链表
func merge(p, q *ListNode) (head *ListNode) {
	dummy := &ListNode{}
	tail := dummy
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next, p.Next, tail, p = p, nil, p, p.Next
		} else {
			tail.Next, q.Next, tail, q = q, nil, q, q.Next
		}
	}

	if p == nil {
		tail.Next = q
	} else {
		tail.Next = p
	}

	head = dummy.Next
	return
}
