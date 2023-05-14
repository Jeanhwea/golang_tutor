package nc0025

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

// 后序遍历
func postorderTraversal(root *TreeNode) (vals []int) {
	var stack []*TreeNode
	var prev *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		peek := stack[len(stack)-1]
		if peek.Right == nil || peek.Right == prev { // 这里 peek 的左子树都已经遍历过
			vals = append(vals, peek.Val)
			prev, stack = peek, stack[:len(stack)-1]
		} else {
			root = peek.Right
		}
	}
	return
}
