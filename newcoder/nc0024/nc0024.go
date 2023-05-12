package nc0024

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func inorderTraversal(root *TreeNode) (vals []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root, stack = stack[len(stack)-1], stack[:len(stack)-1]
		vals = append(vals, root.Val)
		root = root.Right
	}
	return
}
