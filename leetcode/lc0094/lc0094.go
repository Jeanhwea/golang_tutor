package lc0094

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func inorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	curr := root
	stack := []*TreeNode{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)
		curr = curr.Right
	}

	return
}
