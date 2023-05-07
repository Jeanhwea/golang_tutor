package model

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(values []int) (head *ListNode) {
	dummy := &ListNode{}
	tail := dummy
	for _, v := range values {
		curr := &ListNode{Val: v}
		tail, tail.Next = curr, curr
	}
	head = dummy.Next
	return
}

func (l *ListNode) String() (ans string) {
	var values []int
	for ; l != nil; l = l.Next {
		values = append(values, l.Val)
	}
	ans = fmt.Sprintf("%v", values)
	return
}

func (l *ListNode) ToSlice() (values []int) {
	for ; l != nil; l = l.Next {
		values = append(values, l.Val)
	}
	return
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
