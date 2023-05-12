package nc0025

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func postorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		peek := stack[len(stack)-1]
		if peek.Right == nil || peek.Right == prev {
			vals = append(vals, peek.Val)
			prev, stack = peek, stack[:len(stack)-1]
		} else {
			root = peek.Right
		}
	}
	return
}
