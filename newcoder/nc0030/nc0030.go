package nc0030

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func Convert(root *TreeNode) (head *TreeNode) {
	dummy := &TreeNode{}
	tail := dummy

	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		tail.Right, root.Left, tail = root, tail, root
		root = root.Right
	}

	head = dummy.Right
	if head != nil {
		head.Left = nil
	}
	return
}
