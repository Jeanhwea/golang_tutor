package lc021

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) (ans *ListNode) {
	dummy := &ListNode{}
	tail := dummy

	p, q := list1, list2
	for p != nil && q != nil {
		if p.Val < q.Val {
			tail.Next = p
			p = p.Next
		} else {
			tail.Next = q
			q = q.Next
		}
		tail = tail.Next
	}

	if p == nil {
		tail.Next = q
	} else {
		tail.Next = p
	}

	ans = dummy.Next
	return
}
