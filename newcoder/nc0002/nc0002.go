package nc0002

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// BM2 链表内指定区间反转
func reverseBetween(head *ListNode, m int, n int) (ans *ListNode) {
	dummy := &ListNode{
		Next: head,
	}
	var curr, tail1, head2, tail2 *ListNode
	curr = dummy

	// 跳过前 m-1 个节点
	for i := 0; i < m; i++ {
		tail1, curr = curr, curr.Next
	}

	// 翻转 m 到 n 节点, 头插法
	for i := m; i <= n; i++ {
		if tail2 == nil {
			curr, head2, tail2 = curr.Next, curr, curr
		} else {
			curr, head2, curr.Next = curr.Next, curr, head2
		}
	}

	// 合并 3 个子链表
	tail1.Next, tail2.Next = head2, curr

	ans = dummy.Next
	return
}
