package nc0023

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func preorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		if root == nil {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		vals = append(vals, root.Val)
		if root.Right != nil {
			stack = append(stack, root.Right)
		}
		root = root.Left
	}
	return
}
