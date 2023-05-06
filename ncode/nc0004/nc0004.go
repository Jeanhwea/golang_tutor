package nc0004

func Merge(pHead1 *ListNode, pHead2 *ListNode) (ans *ListNode) {
	dummy := &ListNode{}
	tail := dummy
	for pHead1 != nil && pHead2 != nil {
		curr := pHead2
		if pHead1.Val < pHead2.Val {
			curr, pHead1 = pHead1, pHead1.Next
		} else {
			pHead2 = pHead2.Next
		}
		tail.Next, tail, curr = curr, curr, curr.Next
	}

	if pHead1 != nil {
		tail.Next = pHead1
	} else {
		tail.Next = pHead2
	}

	return dummy.Next
}
