package lc0144

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func preorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return
}
