package lc0145

import (
	. "github.com/jeanhwea/golang_tutor/common/model"
)

func postorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	var prev *TreeNode
	curr := root
	stack := []*TreeNode{}
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		top := stack[len(stack)-1]
		if top.Right == nil || top.Right == prev {
			ans = append(ans, top.Val)
			prev = top
			stack = stack[:len(stack)-1]
		} else {
			curr = top.Right
		}
	}
	return
}
