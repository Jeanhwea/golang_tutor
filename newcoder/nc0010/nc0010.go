package nc0010

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 两个链表的第一个公共结点
func FindFirstCommonNode(pHead1 *ListNode, pHead2 *ListNode) (ans *ListNode) {
	p, q := pHead1, pHead2
	for p != q {
		if p != nil {
			p = p.Next
		} else {
			p = pHead2
		}
		if q != nil {
			q = q.Next
		} else {
			q = pHead1
		}
	}

	// 路程 a: pHead1 -- meet
	// 路程 b: pHead2 -- meet
	// 路程 c: meet   -- end
	// 相遇时 p 走过的路程 = a + c + b
	// 相遇时 q 走过的路程 = b + c + a
	// 所以相遇的节点就是交点

	ans = p
	return
}
